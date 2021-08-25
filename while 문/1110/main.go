package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var count int

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n)

	n2 := n

	for {
		count++
		var sum int
		if n2 < 10 {
			sum = n2
		} else {
			sum = n2/10 + n2%10
		}
		n2 = n2%10*10 + sum%10
		if n2 == n {
			break
		}
	}
	fmt.Println(count)
}

// 26 -> 4 // 55-> 3
