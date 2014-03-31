
maketables: maketables.go
	go build $^

tables:	maketables
	./maketables > tables.go
	gofmt -w tables.go
