package main

import "fmt"

// NOTE:This example is done with recursive functions, Please make sure you learn it first
// NOTE:Array must be ordered to make Interpolation Search work
func main() {
	arr := []int{1, 5, 7, 9, 10, 14, 17, 19, 22, 31, 34} // or example array
	x := 10                                              // number we want to find in the array
	low := 0
	high := len(arr) - 1
	xIndex := interpolationSearch(arr, x, low, high)
	if xIndex != -1 {
		fmt.Println(xIndex)
	} else {
		fmt.Println("not found")
	}
}

func interpolationSearch(arr []int, x int, low int, high int) int {
	// calculate the pos with the formula its like middle value in binary search
	pos := low + ((high - low) * (x - arr[low]) / (arr[high] - arr[low]))

	// we return the index of the x in array
	if arr[pos] == x {
		return pos
	}
	// if number we want to find is bigger than the pos
	if arr[pos] < x {
		low = pos + 1
		return interpolationSearch(arr, x, low, high)
	}
	// if number we want to find is smaller than pos in the array
	if arr[pos] > x {
		high = pos - 1
		return interpolationSearch(arr, x, low, high)
	}
	return -1
}
