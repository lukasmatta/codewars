package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{}
	a = nil
	b := []int{1, 5, 9, 25, 16}
	fmt.Printf("Result is: %v", Comp(a, b))
}

func Comp(array1 []int, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}

	b_c := array2
	for _, a := range array1 {
		fmt.Printf("testing %v\n", a)
		for index, b := range b_c {
			fmt.Printf("is %v == %v^2?\n", b, a)
			if b == int(math.Pow(float64(a), 2)) {
				fmt.Printf("yes!\n")
				b_c = append(b_c[:index], b_c[index+1:]...)
				fmt.Printf("new b_c is %v\n", b_c)
				break
			} else {
				continue
			}
		}
	}

	if len(b_c) == 0 {
		return true
	}
	return false
}
