package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &t)

	var a, b int

	for i := 1; i <= t; i++ {

		fmt.Fscan(reader, &a, &b)
		fmt.Fprintf(writer, "Case #%d: %d\n", i, a+b)
	}
	writer.Flush()

}
