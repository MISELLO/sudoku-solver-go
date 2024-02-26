.DEFAULT_GOAL := run

check:
	gocyclo -over 15 .
	goimports -d .
	misspell .
	golint .
.PHONY: check

build: check
	go build -o sudoku-solver.exe main.go
.PHONY: build

run: build
	./sudoku-solver.exe
.PHONY: run

clean:
	go clean
.PHONY: clean
