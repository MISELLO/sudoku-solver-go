.DEFAULT_GOAL := run

check:
	gocyclo -over 15 .
	goimports -d .
	misspell .
	golint .
.PHONY: check

build: check
	go build -o sudoku-solver.exe main.go input.go output.go
.PHONY: build

run: build
	./sudoku-solver.exe -ms 20 -b 123456789456789123789123456234567891567891234891234567345670000000000000000000000
.PHONY: run

clean:
	go clean
.PHONY: clean
