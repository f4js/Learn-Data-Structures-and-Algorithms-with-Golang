package main

import (
	"fmt"
)

func main() {
	friendToFind := "Seller" // the name we want to find
	graph := graphMaker()
	var searched []string                  // searched names
	var queue []string                     // this is our queue
	queue = append(queue, graph["You"]...) // the first queue we will add it manually

	// we iterate on the queue until queue length is empty
	for len(queue) > 0 {
		person := queue[0] // get the person

		// we check if the person is searched before or no , if it was searched we will remove it from queue
		isSearched := containInSlice(searched, person)
		if isSearched == true {
			queue = queue[1:]
			continue
		} else if person == friendToFind { // here we will check if this person is the person we are looking at
			fmt.Println("I have find it", person)
			break
		} else {
			// if this person wa not the name we are looking at we will remove it from queue and add its friends to
			// the queue
			queue = queue[1:]
			queue = append(queue, graph[person]...)
			searched = append(searched, person)
		}
	}

}

// graphMaker is a function that for example generate a graph
func graphMaker() map[string][]string {
	graph := map[string][]string{ // it's key type is string and value type is array of string
		"You": {"Alice", "Bob", "Claire"},
		// First Degree
		"Alice":  {"Tom"},
		"Bob":    {"Tom", "Jerry"},
		"Claire": {"Anton", "Tomas"},
		// Second Degree
		"Tom":   {},
		"Jerry": {},
		"Anton": {},
		"Tomas": {"Seller"},
		// Third Degree
		"Seller": {},
	}
	return graph
}

// containInSlice is a function to check if the given variable is existing in slice
func containInSlice(s []string, str string) bool {
	for _, value := range s {
		if value == str {
			return true
		}
	}
	return false
}
