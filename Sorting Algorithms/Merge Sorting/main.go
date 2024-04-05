package main

import "fmt"

func main() {
	arr := []int{1, 2, 43, 22, 11, 66, 77, 3, 88, 99, 12, 6, 9, 73, 4, 23, 64, 72, 97} // array we gunna to sort
	sortedArr := mergeSort(arr)
	fmt.Println(sortedArr)

}

func mergeSort(arr []int) []int {
	var sortedArr []int
	// if the length of the array is 1 or 0 we don't need to sort it, so we will return it
	if len(arr) <= 1 {
		return arr
	}

	// finding the middle index
	middle := len(arr) / 2

	// we divide the main array into left and right of the main array , then we use recursive on them to sort them
	leftSortedArray := mergeSort(arr[:middle])
	rightSortedArray := mergeSort(arr[middle:])

	rightI, leftI := 0, 0

	// we sort the array in here by checking the left and right side of the array
	for leftI < len(leftSortedArray) && rightI < len(rightSortedArray) {
		if leftSortedArray[leftI] < rightSortedArray[rightI] {
			sortedArr = append(sortedArr, leftSortedArray[leftI])
			leftI++
		} else if rightSortedArray[rightI] < leftSortedArray[leftI] {
			sortedArr = append(sortedArr, rightSortedArray[rightI])
			rightI++
		}
	}

	sortedArr = append(sortedArr, leftSortedArray[leftI:]...)
	sortedArr = append(sortedArr, rightSortedArray[rightI:]...)

	return sortedArr

}
