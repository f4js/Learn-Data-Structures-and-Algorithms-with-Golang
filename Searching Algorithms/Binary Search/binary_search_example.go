package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binarySearch(array []int, choose int) int {
	low := array[0]        // lowest number in the array
	high := len(array) - 1 // largest number in the array
	for low <= high {      // until lowest number is smaller or equal to the largest number
		middle := (low + high) / 2 // we find the middle number here
		guess := array[middle]     // our guess

		if guess == choose { // if we find the choose then we will return
			return guess
		}
		if guess < choose { // if choose is lower than guess
			low = middle + 1
		} else { // if choose is higher than guess
			high = middle - 1
		}
	}
	return -1 // if we didn't find the number choose
}

func main() {
	listRange := 1000000000
	// making array
	ourArray := make([]int, listRange, listRange)
	for i := 0; i < len(ourArray); i++ {
		ourArray[i] = i + 1
	}

	now := time.Now()                                      // start the timer
	number := binarySearch(ourArray, rand.Intn(listRange)) // start the function with random number to find
	if number == -1 {
		fmt.Println("there was no number like you choose in array")
	} else {
		fmt.Printf("we have find it %v \n", number)
	}
	fmt.Println(time.Since(now)) // print how much time it took to find the number
}
