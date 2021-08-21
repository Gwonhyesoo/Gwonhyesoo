//합

package main

import "fmt"

func main() {
	var a int
	var sum int

	fmt.Scan(&a)

	for i := 1; i <= a; i++ {
		sum += i
	}
	fmt.Println(sum)
}
