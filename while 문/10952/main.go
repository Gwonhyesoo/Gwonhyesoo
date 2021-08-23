package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b int

	defer writer.Flush()
	for {
		fmt.Fscan(reader, &a, &b)
		if a == 0 && b == 0 {
			break
		}
		fmt.Fprintln(writer, a+b)
	}
}
