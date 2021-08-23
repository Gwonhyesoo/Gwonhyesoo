package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var x int
	var a int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &x)

	defer writer.Flush()

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a)
		if x > a {
			fmt.Fprintf(writer, "%d ", a)
		}

	}

}
