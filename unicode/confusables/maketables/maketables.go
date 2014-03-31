// Confusables table generator.
// See http://www.unicode.org/reports/tr39/

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	loadUnicodeData()
	makeTables()
}

var url = flag.String("url",
	"http://www.unicode.org/Public/security/latest/",
	"URL of Unicode database directory")

var localFiles = flag.Bool("local",
	false,
	"data files have been copied to the current directory; for debugging only")

// confusables.txt has form:
//	309C ;	030A ;	SL	#* ( ゜ → ̊ ) KATAKANA-HIRAGANA SEMI-VOICED SOUND MARK → COMBINING RING ABOVE	# →ﾟ→→゚→
// See http://www.unicode.org/reports/tr39/ for full explanation
// The fields:
const (
	CSourceCodePoint = iota
	CTargetCodePoint
	CType
	NumField

	MaxChar = 0x10FFFF // anything above this shouldn't exist
)

func openReader(file string) (input io.ReadCloser) {
	if *localFiles {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		input = f
	} else {
		path := *url + file
		// log.Println("Downloading " + path)
		resp, err := http.Get(path)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != 200 {
			log.Fatal("bad GET status for "+path, resp.Status)
		}
		input = resp.Body
	}
	return
}

func parsePoint(pointString string, line string) rune {
	x, err := strconv.ParseUint(pointString, 16, 64)
	point := rune(x)
	if err != nil {
		log.Fatalf("%.5s...: %s", line, err)
	}
	if point == 0 {
		log.Fatalf("%5s: Unknown rune %X", line, point)
	}
	if point > MaxChar {
		log.Fatalf("%5s: Rune %X > MaxChar (%X)", line, point, MaxChar)
	}

	return point
}

var confusablesMap = make(map[rune][]rune)

func parseCharacter(line string) {
	if len(line) == 0 || line[0] == '#' {
		return
	}
	// strip BOM
	if len(line) > 3 && bytes.Compare(([]byte(line[0:3])), []byte{0xEF, 0xBB, 0xBF}) == 0 {
		return
	}
	field := strings.Split(line, " ;\t")
	if len(field) != NumField {
		log.Fatalf("%5s: %d fields (expected %d)\n", line, len(field), NumField)
	}

	if !strings.HasPrefix(field[2], "MA") {
		// The MA table is a superset anyway
		return
	}

	sourceRune := parsePoint(field[CSourceCodePoint], line)
	var targetRune []rune
	targetCodePoints := strings.Split(field[CTargetCodePoint], " ")
	for _, targetCP := range targetCodePoints {
		targetRune = append(targetRune, parsePoint(targetCP, line))
	}

	confusablesMap[sourceRune] = targetRune
}

func loadUnicodeData() {
	f := openReader("confusables.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parseCharacter(scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

const fileHeader = `// Generated by running maketables
// DO NOT EDIT

package confusables
`

func makeTables() {
	fmt.Println(fileHeader)
	fmt.Println("// confusablesMap")
	fmt.Print("var confusablesMap = map[rune][]rune{\n\n")
	for k, v := range confusablesMap {
		fmt.Printf("0x%.8X: []rune{\n", k)
		for _, r := range v {
			fmt.Printf("0x%.8X,\n", r)
		}
		fmt.Println("},")
	}
	fmt.Print("}\n\n")
}
