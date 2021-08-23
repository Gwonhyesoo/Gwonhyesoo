//기찍 N

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &N)

	for i := N; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	writer.Flush()
}
