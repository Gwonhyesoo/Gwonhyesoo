//알람 시계

package main

import "fmt"

func main() {
	var h int
	var m int

	fmt.Scan(&h)
	fmt.Scan(&m)
	if h == 0 {
		if m >= 45 {
			fmt.Println(h, (m - 45))
		} else if m < 45 {
			h = 23
			m = (m + 60) - 45
			fmt.Println(h, m)
		}
	} else if h != 0 {
		if m >= 45 {
			fmt.Println(h, (m - 45))
		} else if m < 45 {
			fmt.Println((h - 1), (m + 60 - 45))
		}
	}
}
