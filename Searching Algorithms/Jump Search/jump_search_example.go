package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 5, 6, 8, 10, 12, 15, 18, 22, 26, 28, 29, 45}
	x := 26
	xIndex := jumpSearch(arr, len(arr), x)
	fmt.Println(xIndex)
}

func jumpSearch(arr []int, arrLen int, x int) int {
	i := 0
	// square root of the length of the array
	sqrt := int(math.Sqrt(float64(arrLen)))
	copySQRT := sqrt
	// continue loop until value in the array is smaller than x and sqrt is smaller than the length of array
	for arr[sqrt] <= x && sqrt < arrLen {
		i = sqrt
		sqrt += copySQRT
		// if sqrt is bigger than length of array that means reach the end or further of the array, and we didn't find x
		if sqrt > arrLen-1 {
			return -1
		}
	}
	// leaner search on that part of the array that we think the value is
	for j := i; j <= sqrt; j++ {
		index := j
		if arr[index] == x {
			return index
		}
	}
	return -1
}
