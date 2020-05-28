package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	n := flag.Int("n", 1, "the number of columns")
	d := flag.String("d", "\t ", "separator of input")
	sep := flag.String("s", "\t", "separator of output")
	w := flag.Bool("warning", false, "ignore warnings")
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)

	buf := make([][]string, 0, 32)
	colnum := 0
	for s.Scan() {
		f := strings.FieldsFunc(s.Text(), func(r rune) bool { return strings.ContainsRune(*d, r) })
		buf = append(buf, f)

		if colnum < len(f) {
			colnum = len(f)
		}
	}

	for i := 0; i < colnum; i += *n {
		for j := 0; j < len(buf); j++ {
			start := i
			end := i + *n
			if end >= len(buf[j]) {
				end = len(buf[j])
			}
			if start >= end {
				if !*w {
					fmt.Fprintf(os.Stderr, "WARNING: row %d, column %d\n", j, i)
				}
				continue
			}
			l := strings.Join(buf[j][start:end], *sep)
			fmt.Println(l)
		}
	}
}
