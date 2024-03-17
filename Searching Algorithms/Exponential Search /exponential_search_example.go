package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 7, 8, 9, 12, 14, 17, 19} // Array that we gunna search in it
	key := 9                                       // Key we gunna search for it in array

	keyIndex := exponentialSearch(arr, key)

	// Here we check if we find it
	if keyIndex != -1 {
		fmt.Println("key in the array after exponentialSearch is in index of ", keyIndex)
	} else {
		fmt.Println("unable to find the key")
	}
}

// binarySearch is function that get array from exponentialSearch function and do a binary search on it to find key
func binarySearch(arr []int, key int) int {
	if len(arr) <= 2 {
		// If it has 0 length
		if len(arr) == 0 {
			return 0
		}
		// If array has length of 1
		if len(arr) == 1 {
			if arr[0] == key {
				return 0
			} else {
				return -1
			}
		}
		// If it has length of 2
		if len(arr) == 2 {
			if arr[0] == key {
				return 0
			} else if arr[1] == key {
				return 1
			} else {
				return -1
			}
		}
		return -1
	}
	// Main part of binary search
	start := 0
	end := len(arr) - 1
	for start <= end {
		middle := (start + end) / 2
		if arr[middle] == key {
			return middle
		} else if arr[middle] > key {
			end = middle - 1
		} else if arr[middle] < key {
			start = middle + 1
		}
	}
	return -1
}

// exponentialSearch will find the range of the key we want to find
func exponentialSearch(arr []int, key int) int {
	// Check if the key is in 0 and 1 index
	if arr[0] == key {
		return 0
	} else if arr[1] == key {
		return 1
	}
	// Main part of the "Exponential Search"
	i := 1
	for i <= len(arr) {
		if arr[i] <= key {
			i *= 2
		} else {
			break
		}
	}
	// If the "i" is bigger than length of array, we put "i" length of array
	if i > len(arr) {
		i = len(arr)
	}
	// Do the binary search
	return binarySearch(arr[i/2:i], key)
}
