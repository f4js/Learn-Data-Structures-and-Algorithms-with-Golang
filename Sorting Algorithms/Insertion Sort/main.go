package main

import "fmt"

func main() {
	arr := []int{76, 54, 43, 99, 555, 222, 111, 999}
	sortedArr := insertionSort(arr)
	fmt.Println(sortedArr)

}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		keyIndex := i
		for keyIndex > 0 && keyIndex <= len(arr) {
			if arr[keyIndex-1] > arr[keyIndex] {
				iValue := arr[keyIndex]
				i2Value := arr[keyIndex-1]

				arr[keyIndex] = i2Value
				arr[keyIndex-1] = iValue

				keyIndex--
			} else {
				break
			}
		}
	}

	return arr

}
