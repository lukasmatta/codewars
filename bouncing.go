package main

import (
	"fmt"
)

func main() {
	fmt.Printf("3 == %v?\n", BouncingBall(3, 0.66, 1.5))
}

func BouncingBall(h, bounce, window float64) int {
	if h <= 0 || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}

	times := 1
	new_height := h

	for new_height > window {
		times += 1
		new_height *= bounce
		fmt.Printf("new height is %v\n", new_height)
	}

	return times
}
