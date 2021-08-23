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

	for i := 1; i < t+1; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintf(writer, "Case #%d: %d + %d = %d \n", i, a, b, a+b)
	}
	writer.Flush()
}
