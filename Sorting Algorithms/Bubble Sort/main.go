package main

import "fmt"

func main() {
	arr := []int{1, 2, 50, 12, 22, 11, 90}
	sorted := 0
	for i := 0; ; i++ {
		if i == len(arr)-1 {
			i = 0
			sorted = 0
		}
		index1 := i
		index2 := i + 1
		if arr[index1] > arr[index2] {
			value1 := arr[index1]
			value2 := arr[index2]

			arr[index1] = value2
			arr[index2] = value1
		} else {
			sorted++
			if sorted == len(arr)-1 {
				fmt.Println(arr)
				break
			}
		}
	}
}
