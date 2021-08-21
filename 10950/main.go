package main

import "fmt"

func main() {
	var a int
	var b int
	var t int

	fmt.Scan(&t)
	arr := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b)
		arr[i] = a + b
	}

	for i := 0; i < t; i++ {
		fmt.Println(arr[i])
	}
}
