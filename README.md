# go_balanced_brackets

Checks if input file with lines of brackets are balanced, and match an output file give.

## Usage
```
Usage of ./go_balanced_brackets:
  -i string
        Input file (default "input.txt")
  -o string
        Output file (default "output.txt")
```

## File Format
File format is based from input and output files from HackerRank.

Input files first line is the count of brackets, and all other lines are brackets to be check.
Ouput files are "YES" and "NO" base of if the matching brackets is balanced or not.

## Unit tests and Benchmarks
`go test` will run the unit test
`go test -bench=.` will run the benchmarks and output something like,

```
goos: linux
goarch: amd64
pkg: github.com/amscotti/go_balanced_brackets
cpu: 11th Gen Intel(R) Core(TM) i5-11400 @ 2.60GHz
BenchmarkCalculate-12             115782             10245 ns/op
PASS
ok      github.com/amscotti/go_balanced_brackets        1.297s
```

## Build and run
`go build` and `./go_balanced_brackets -i input17.txt -o output17.txt`
