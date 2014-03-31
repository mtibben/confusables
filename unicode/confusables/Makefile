
maketables: maketables.go
	go build $^

tables.go: maketables
	./maketables > tables.go
	gofmt -w tables.go
	go test

