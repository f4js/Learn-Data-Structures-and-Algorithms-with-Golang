package main

import (
	"fmt"
	"math/rand"
	"time"
)

// listMaker is a function that make an array with random numbers
func listMaker(start int, end int, length int) []int {
	random := rand.New(rand.NewSource(time.Now().Unix())) // make a new seed
	randomSlice := make([]int, length)
	for num := 0; num < length; num++ {
		randomSlice[num] = random.Intn(end-start+1) + start
	}
	return randomSlice
}

// listSorter function that is responsible for sorting the array with recursive
func listSorter(inputSlice []int) (sortedSlice []int) {

	// if slice have 0 or 1 element we will return its self
	if len(inputSlice) == 0 || len(inputSlice) == 1 {
		return inputSlice

		// if slice have just two element we sort just the two and return it
	} else if len(inputSlice) == 2 {
		firstNumber := inputSlice[0]
		secondNumber := inputSlice[1]
		if firstNumber > secondNumber {
			sortedSlice = append(sortedSlice, secondNumber, firstNumber)
			return sortedSlice
		} else {
			sortedSlice = append(sortedSlice, firstNumber, secondNumber)
			return sortedSlice
		}
	}

	// we choose a pivot
	pivot := inputSlice[len(inputSlice)-(len(inputSlice)/2)]

	var smaller []int
	var bigger []int
	var middles []int

	// separate the elements in slice to bigger and smaller
	for index, _ := range inputSlice {
		if inputSlice[index] > pivot {
			bigger = append(bigger, inputSlice[index])

		} else if inputSlice[index] == pivot {
			middles = append(middles, inputSlice[index])

		} else {
			smaller = append(smaller, inputSlice[index])
		}
	}
	// recursive function to sort
	smaller = listSorter(smaller)
	bigger = listSorter(bigger)

	// append the slices together
	sortedSlice = append(smaller, middles...)
	sortedSlice = append(sortedSlice, bigger...)
	return sortedSlice
}

func main() {
	// makes a random list with given args
	ourSlice := listMaker(500, 900000, 100000)

	// Starts timer
	start := time.Now()
	// sorts the list
	list := listSorter(ourSlice)
	// calculate the time and print it
	elapsed := time.Since(start).Seconds()
	fmt.Println(elapsed)
	fmt.Println(list)

}
