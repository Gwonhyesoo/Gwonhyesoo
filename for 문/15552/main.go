package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int

	reader := bufio.NewReader(os.Stdin) //표준입출력방식. stdin입력받은후 stdout으로 출력
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &t)

	var a, b int

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &a, &b) //Fscan 대량의 입력받을때
		fmt.Fprintln(writer, a+b)   // 빠르게 출력 그러나 /n 이 많을땐 그냥print가 더 빠름
	}
	writer.Flush()
}
