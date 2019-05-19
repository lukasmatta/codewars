package main

import (
	"fmt"
)

func main() {
	FindUniq([]float32{6, 1, 1, 1, 1, 1, 1})
}

func FindUniq(arr []float32) float32 {
	var good_value float32
	var new_good_value float32
	var wrong_value float32
	good_value = arr[0]
	fmt.Printf("Good value %v.\n", good_value)
	new_good_value = arr[1]
	fmt.Printf("Another Good value %v.\n", new_good_value)
	if new_good_value != good_value {
		fmt.Printf("Chose two different values %v and %v.\n", good_value, new_good_value)
		fmt.Printf("Getting third value to decide which is %v.\n", arr[2])
		decide_value := arr[2]

		if decide_value == good_value {
			fmt.Printf("Wrong value was %v.\n", new_good_value)
			wrong_value = new_good_value
		} else {
			fmt.Printf("Wrong up value was %v.\n", good_value)
			wrong_value = good_value
		}

		return wrong_value
	}

	for _, element := range arr[2:] {
		fmt.Printf("Is %v == %v ??\n", element, good_value)
		if element != good_value {
			fmt.Printf("Got it! Wrong value is %v.\n", element)
			wrong_value = element
			break
		}
	}

	return wrong_value
}
