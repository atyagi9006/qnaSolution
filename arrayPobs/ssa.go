package main

import "fmt"

func main() {
	sortedInput := []int{3, 4, 5, 6, 7}
	fmt.Println("Sorted Array Input: ", sortedInput)
	var result []int
	for _, n := range sortedInput {
		result = append(result, (n * n))
	}
	fmt.Println("Result Sqr Of Sroted Arry Elements :", result)
}
