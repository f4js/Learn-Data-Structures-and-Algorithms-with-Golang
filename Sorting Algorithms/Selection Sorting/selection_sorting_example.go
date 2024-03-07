package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var firstList []int
	var sort []int

	firstList = makeList(firstList, 100)
	fmt.Println(firstList)

	timer := time.Now() // this line will get the current time later we gunna calculate the total execute time

	for range firstList {
		biggestIndex, biggest := findBiggest(firstList)
		sort = append(sort, biggest)
		firstList = deleteIndex(firstList, biggestIndex)
	}
	fmt.Println("\n took: ", time.Since(timer))
	fmt.Println("\n", sort)
}

// makeList is the function that make a list and then shuffle the new array
func makeList(array []int, index int) []int {
	// here we make the array
	for i := 0; i < index; i++ {
		array = append(array, i)
	}
	// here we shuffle the array
	arrayShuffle(array)
	return array
}

// arrayShuffle is function that shuffle the array
func arrayShuffle(array []int) {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	random.Shuffle(len(array), func(i, j int) {
		array[i], array[j] = array[j], array[i]
	})
}

// findBiggest is the function that find the biggest number in the array
func findBiggest(arr []int) (int, int) {
	var biggest int
	var biggestIndex int
	for index, value := range arr {
		if value > biggest {
			biggest = value
			biggestIndex = index
		}
	}
	return biggestIndex, biggest
}

// deleteIndex this function will delete the given index
func deleteIndex(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)

}
