package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscan(reader, &n)
	defer writer.Flush()
	for i := 0; i < n; i++ {
		for j := n - 1 - i; j > 0; j-- {
			fmt.Fprint(writer, " ")
		}
		for k := 0; k < i+1; k++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprint(writer, "\n")
	}
}
