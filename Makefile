run: build
	./mycorrhiza metarrhiza

build:
	go build .

test:
	go test .

help:
	echo "Read the Makefile to see what it can do. It is simple."

