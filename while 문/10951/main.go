package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		_, err := fmt.Fscanln(reader, &a, &b)
		if err != nil { //숫자가 아닌경우
			break
		}
		fmt.Fprintln(writer, a+b)
	}
}
