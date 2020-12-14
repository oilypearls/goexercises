package main

import "fmt"

func Sum(arr []int) int {
	var arrLen int = len(arr)
	fmt.Println(arrLen)
	result := 0

	for _, val := range arr {
		result += val
	}
	return result
}
