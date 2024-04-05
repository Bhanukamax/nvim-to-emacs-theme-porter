main: *.go **/*.go
	go build .

.PHONY: run
run: main
	./themer
