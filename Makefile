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
	./sudoku-solver.exe 060070004008090760710406005007500326000903000346007500600308079073010200200050010
.PHONY: run

clean:
	go clean
.PHONY: clean
