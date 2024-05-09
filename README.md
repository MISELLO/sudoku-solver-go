
# Sudoku Solver Go

This is my solution (written in [Go](https://golang.org/)) to be used in order to solve sudokus via command line.

## Badges

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Go Report Card](https://goreportcard.com/badge/github.com/misello/sudoku-solver-go)](https://goreportcard.com/report/github.com/misello/sudoku-solver-go)

## What is a sudoku and why it needs to be solved?

A sudoku is a logic-based, combinatorial number-placement puzzle. In classic Sudoku, the objective is to fill a 9 × 9 grid with digits so that each column, each row, and each of the nine 3 × 3 subgrids that compose the grid (also called "boxes", "blocks", or "regions") contains all of the digits from 1 to 9. The puzzle setter provides a partially completed grid, which for a well-posed puzzle has a single solution ([Wikipedia](https://en.wikipedia.org/wiki/Sudoku)).

The reasons someone would like to solve automatically a sudoku instead of taking it as a personal challenge may vary depending on individual preferences and circumstances that are not necessarily cheating. This could be to check a solution, educative or give you a hand when you are stuck on a puzzle.

## Compatibility

Everything has been tested under **Ubuntu 22.04.4 LTS** and **go version 1.21.3**.

I have not tested it under windows, but I assume the visual input would fail (better use the string or file format) and the colours will not work as expected, I recommend to use the -nc flag.

## Compiling

The code comes with a Makefile that can be executed with:

    make

Alternativelly, you can compile it yourself with:

    go build -o sudoku-solver.exe main.go input.go output.go

It is recomended but not mandatory to have installed `gocyclo`, `goimports`, `misspell` and `golint`.

## Usage

By default, the generated executable file is called `sudoku-solver.exe`. You can call it with arguments.

    $ ./sudoku-solver.exe [FLAGS] [INPUT]

### Flags
- -a &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;All solutions, displays all solutions instead of just the first one.  
- -b &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Brute-force, it uses brute-force when necessary.  
- -bt &nbsp;&nbsp;&nbsp;Brute-force time, time in seconds the brute-force algorithm is allowed to run. If this time is reached a "(+)" will appear next to the Solutions amount output (automatically activates brute-force [-b]). (default 60)  
- -ms &nbsp;&nbsp;Max solutions, defines the maximum number of solutions the algorithm will calculate. If this number is reached a "(+)" will appear next to the Solutions amount output.(automatically activates brute-force [-b]). (default 10)  
- -nc &nbsp;&nbsp;&nbsp;No color, if set disables the colors of the results.  
- -s &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;String format output, displays only the result as a string of 81 digits.  

### Input  
Data should be introduced in one of this three ways:  
- String: A 81 digit string that represents the sudoku in writing order (top to bottom, left to right) using zeroes for the unknown cells.
- File:   A text file that contains the sudoku we want to solve (spaces are allowed).
- Visual: If no arguments for the sudoku are provided, an empty board will be displayed, fill it manually and press enter.

## Pending

There are some more features pending to be implemented
- Generate solutions changing lines and rows from a base solved sudoku
- Add execution time
- Add an estimated difficulty
- Search for solutions with concurrency/parallelism
- A goroutine that tells how many solutions have been found so far
- The ability to abort an execution
