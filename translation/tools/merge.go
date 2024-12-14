package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// correctTranslated removes misplaced backslashes or other unwanted characters from the translated text
func correctTranslated(trans string) string {
	result := strings.Builder{}
	i := 0
	for i < len(trans) {
		if trans[i] == '\\' {
			switch {
			case i+1 < len(trans) && (trans[i+1] == '\\' || trans[i+1] == ' ' || trans[i+1] == '(' || trans[i+1] == ')'):
				i += 2
				continue
			case i+1 < len(trans) && trans[i+1] == '"':
				result.WriteByte(trans[i])
				i++
			}
		}
		result.WriteByte(trans[i])
		i++
	}
	return result.String()
}

// build formats and writes the input and translated text to the output file
func build(in string, trans string, out *os.File) {
	var e int
	if strings.HasPrefix(in, "\"") {
		e = strings.Index(in, "\":") + 2
		fmt.Fprintf(out, "%-120s\"%s\",\n", in[:e], trans)
	} else {
		e = strings.Index(in, "\",") + 2
		fmt.Fprintf(out, "%-120s\"%s\"],\n", in[:e], trans)
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Syntax: merge <template file> <translation file> <output file>\n")
		os.Exit(1)
	}

	templateFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open %s\n", os.Args[1])
		os.Exit(1)
	}
	defer templateFile.Close()

	translationFile, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open %s\n", os.Args[2])
		os.Exit(1)
	}
	defer translationFile.Close()

	outputFile, err := os.Create(os.Args[3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open %s\n", os.Args[3])
		os.Exit(1)
	}
	defer outputFile.Close()

	templateScanner := bufio.NewScanner(templateFile)
	translationScanner := bufio.NewScanner(translationFile)

	for templateScanner.Scan() {
		s := templateScanner.Text()
		switch s[0] {
		case '"', '[':
			if !translationScanner.Scan() {
				fmt.Fprintf(os.Stderr, "Error: wrong translation text\n")
				os.Exit(1)
			}
			transLine := correctTranslated(translationScanner.Text())
			build(s, transLine, outputFile)
		default:
			fmt.Fprintln(outputFile, s)
		}
	}

	if err := templateScanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading template file: %v\n", err)
		os.Exit(1)
	}

	if err := translationScanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading translation file: %v\n", err)
		os.Exit(1)
	}
}
