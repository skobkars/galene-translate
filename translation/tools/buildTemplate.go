package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func build(input string, output *os.File) {
	list := strings.HasPrefix(input, "\"")

	if list {
		index := strings.Index(input, "\":")
		if index == -1 {
			fmt.Fprintf(os.Stderr, "Invalid input: %s\n", input)
			os.Exit(1)
		}
		formatted := fmt.Sprintf("%-120s\"\",\n", input[:index+2])
		output.WriteString(formatted)
	} else {
		index := strings.Index(input, "\",")
		if index == -1 {
			fmt.Fprintf(os.Stderr, "Invalid input: %s\n", input)
			os.Exit(1)
		}
		formatted := fmt.Sprintf("%-120s\"\"],\n", input[:index+2])
		output.WriteString(formatted)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Syntax: buildTemplate <input file> <output file>")
		os.Exit(1)
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open %s: %v\n", os.Args[2], err)
		os.Exit(1)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		build(line, outputFile)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		os.Exit(1)
	}
}
