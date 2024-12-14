package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func build(in string) string {
	var out string
	out = "";
	if in[0] == '[' {
		s, __ := strings.CutPrefix(in, "[\"")
		if ( __ ) {
			str := strings.Split(s, "\",")
			out = str[0]
		}
	}
	if in[0] == '"' {
		s,__  := strings.CutPrefix(in, "\"")
		if ( __ ) {
			str := strings.Split(s, "\":")
			out = str[0]
		}
	}
	return out;
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Syntax: extract <input file> <output file>\n")
		os.Exit(1)
	}

	in, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open %s\n", os.Args[1])
		os.Exit(1)
	}
	defer in.Close()

	out, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open %s\n", os.Args[2])
		os.Exit(1)
	}
	defer out.Close()

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		s := scanner.Text()
		if ( s != "" ) {
			switch s[0] {
			case '"', '[':
				t := build(s)
				fmt.Fprintln(out, t)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}
}

