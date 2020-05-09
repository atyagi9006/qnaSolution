package main

import "fmt"

func main() {
	fmt.Println("start ...")
	input := []int{4, 5, 3, 1, 6, 3}
	fmt.Println("Find missing and repeated ... from ", input)
	inputRange := len(input)
	countMap := make(map[int]int)
	for _, n := range input {
		if val, ok := countMap[n]; ok {
			countMap[n] = val + 1
		} else {
			countMap[n] = 1
		}
	}

	for i := 1; i <= inputRange; i++ {
		if val, ok := countMap[i]; ok {
			if val > 1 {
				fmt.Println("Repeated num ", i)
			}
		} else {
			fmt.Println("Missing num ", i)
		}
	}
}
