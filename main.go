package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/schollz/progressbar/v3"
)

const (
	YES = "YES"
	NO  = "NO"
)

func boolToWord(check bool) string {
	if check {
		return YES
	} else {
		return NO
	}
}

type Bracket struct {
	input    string
	expected string
	actual   string
	match    bool
}

func readFiles(inputFile string, outputFile string) ([]string, []string, int64, error) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, nil, 0, err
	}

	output, err := os.ReadFile(outputFile)
	if err != nil {
		return nil, nil, 0, err
	}

	inputList := strings.Split(string(input), "\n")
	outputList := strings.Split(string(output), "\n")

	count, err := strconv.ParseInt(inputList[0], 10, 64)
	if err != nil {
		return nil, nil, 0, err
	}

	return inputList[1:], outputList, count, nil
}

func worker(jobs <-chan Bracket, results chan<- Bracket) {
	for bracket := range jobs {
		actual := IsBalanced(bracket.input)
		bracket.match = bracket.expected == boolToWord(actual)

		results <- bracket
	}
}

func run() error {
	var inputFile string
	var outputFile string

	flag.StringVar(&inputFile, "i", "input.txt", "Input file")
	flag.StringVar(&outputFile, "o", "output.txt", "Output file")
	flag.Parse()

	var inputList, outputList, count, err = readFiles(inputFile, outputFile)
	if err != nil {
		return err
	}

	jobs := make(chan Bracket, 1000)
	results := make(chan Bracket, 1000)

	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(jobs, results)
	}

	go func() {
		for i, input := range inputList {
			jobs <- Bracket{input: input, expected: outputList[i]}
		}
	}()

	bar := progressbar.Default(count)

	var i int64
	for i = 0; i < count; i++ {
		bracket := <-results
		if !bracket.match {
			fmt.Printf("Mismatch: %s\n %s vs %s\n", bracket.input, bracket.expected, bracket.actual)
		}
		bar.Add(1)
	}

	close(jobs)
	close(results)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
