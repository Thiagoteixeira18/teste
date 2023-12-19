package services

import "fmt"

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func mais() {
	fmt.Println(Max(5, 7))
}