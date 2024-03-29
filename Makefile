main: main.go
	go build main.go

.PHONY: run
run: main
	./main
