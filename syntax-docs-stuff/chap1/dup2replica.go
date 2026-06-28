package main 

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	count := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin,count)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f,count)
			f.Close()
		}
	}
	for line, n := range(count) {
		if n > 1 {
			fmt.Printf("%d\t %s\n", n, line)
		}
	}
}
func countLines(f *os.File, count map[string]int) {
		input := bufio.NewScanner(f)
		for input.Scan() {
			count[input.Text()]++
		}
	}

