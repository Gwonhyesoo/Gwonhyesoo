package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a * ((b % 100) % 10))
	fmt.Println(a * ((b % 100) / 10))
	fmt.Println(a * (b / 100))
	fmt.Println(a * b)
}
