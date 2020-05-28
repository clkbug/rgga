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
			l := strings.Join(buf[j][i:i+*n], *sep)
			fmt.Println(l)
		}
	}
}
