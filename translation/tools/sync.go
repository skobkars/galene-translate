package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func compare(input string, line string, out *os.File) bool {
	line = strings.TrimSpace(line)
	if input == line {
		out.WriteString(input + "\n")
		return true
	}
	var str []string;
	if idx := strings.Index(line, "\":"); idx != -1 {
		str = strings.SplitAfterN(line, "\":", 2);
		fmt.Fprintf(out,"%-120s\"\",\n",str[0])
	} else if idx := strings.Index(line, "\","); idx != -1 {
		str = strings.SplitAfterN(line,"\",", 2);
		fmt.Fprintf(out,"%-120s\"\"],\n",str[0])
	} else {
		str[0] = line
		out.WriteString(str[0] + "\n")
	}
	return true
}

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "Syntax: sync <input file> <check file> <output file>")
		os.Exit(1)
	}
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open input file: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()
	checkFile, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open check file: %v\n", err)
		os.Exit(1)
	}
	defer checkFile.Close()
	outputFile, err := os.Create(os.Args[3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	reader := bufio.NewScanner(checkFile)
	reader.Scan()
	var checkLine string
	checkLine = reader.Text()
	for scanner.Scan() {
		line := scanner.Text()
		if compare(line, checkLine, outputFile) {
			reader.Scan()
			checkLine = reader.Text()
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		os.Exit(1)
	}
}
