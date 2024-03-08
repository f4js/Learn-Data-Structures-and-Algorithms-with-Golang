package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	toFind := 50
	myArray := makeList(300) // first make a list with length of 300
	// we iterate throw it until we find the number we want
	for index, number := range myArray {
		if number == toFind {
			fmt.Println("Find the number ", number, " in index", index)
			return

		}
	}
	fmt.Println("didnt find the number ", toFind)

}

// arrayShuffle is function that shuffle the array
func arrayShuffle(array []int) {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	random.Shuffle(len(array), func(i, j int) {
		array[i], array[j] = array[j], array[i]
	})
}

// makeList is the function that make a list and then shuffle the new array
func makeList(length int) (array []int) {
	// here we make the array
	for i := 0; i < length; i++ {
		array = append(array, i)
	}
	// here we shuffle the array
	arrayShuffle(array)
	return array
}
