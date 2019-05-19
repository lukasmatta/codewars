package main

import (
	"fmt"
	"strings"
)

func main() {
	dotest("abcdea", 1)
	dotest("abcdeaB11", 3)
	dotest("indivisibility", 1)
}

func dotest(prod string, exp int) {
	var ans = duplicate_count(prod)
	fmt.Printf("%v == %v\n", ans, exp)
}

func duplicate_count(s1 string) int {
	var letters []string
	var sum int

	for i := 0; i < len(s1); i++ {
		counter := 0
		for _, elem := range letters {
			if strings.ToUpper(elem) == strings.ToUpper(string(s1[i])) {
				counter++
			}
		}

		if counter == 1 {
			sum++
		}

		letters = append(letters, string(s1[i]))
	}

	return sum
}
