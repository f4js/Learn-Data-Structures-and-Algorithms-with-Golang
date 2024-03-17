package main

import (
	"fmt"
	"slices"
)

func main() {
	graph := graphMaker()
	costsGraph, parentGraph := initialCostGraph(graph)
	var processed []string

	node := findLowestCostNode(costsGraph, processed) // initial node to start with
	for len(processed) <= 4 {
		nodeCost := costsGraph[node] // gets the cost map of the lowest node
		nodeNeighbors := graph[node] // gets all the neighbors of the lowest node
		for nKey, _ := range nodeNeighbors {
			// calculate the path to the current node to the neighbor
			nodeNewCost := nodeCost + nodeNeighbors[nKey]
			// if new Cost is lower than the cost of that neighbor that means we find a new path to that neighbor
			// with lower cost, so we will change the cos of that neighbor and change its parent node to the current node
			if costsGraph[nKey] > nodeNewCost {
				costsGraph[nKey] = nodeNewCost
				parentGraph[nKey] = node
			}
		}
		processed = append(processed, node)
		node = findLowestCostNode(costsGraph, processed)
	}
	pathFinder(parentGraph)
}

// graphMaker function will create the graph
func graphMaker() map[string]map[string]int {
	graph := map[string]map[string]int{
		"Start": {
			"A": 6,
			"B": 2,
		},
		"A": {
			"Finish": 1,
		},
		"B": {
			"A":      3,
			"Finish": 5,
		},
		"Finish": {},
	}
	return graph
}

// findLowestCostNode function will find the node with the lowest cost
func findLowestCostNode(cg map[string]int, pr []string) string {
	lowest := "" // initial value
	for key, value := range cg {
		// here we check if the node we are checking is already processed or no
		if slices.Contains(pr, key) {
			continue
		}
		// if lowest is empty we will put the lowest first graph we get
		if len(lowest) <= 0 {
			lowest = key
		}
		// find the lowest cost in the graph every time
		if cg[lowest] > value {
			lowest = key
		}
	}
	return lowest
}

// pathFinder will get the parent graph and find the shortest way to Fin
func pathFinder(parentGraph map[string]string) {
	finalPath := []string{"Start"}
	in := 0
	for { // infinite loop until we reach the Fin
		parent := finalPath[in] // in every loop we find the next element child
		child := ""
		for cKey, cValue := range parentGraph {
			if cValue == parent {
				child = cKey
				break
			}
		}
		// we append the child to left of the slice
		finalPath = append(finalPath, child)
		in += 1
		// if the child is the Fin we will break the loop
		if child == "Finish" {
			break
		}
	}
	fmt.Println(finalPath)
}

// initialCostGraph function that return cost of the neighbors of Start node and parentGraph of neighbors of Start
func initialCostGraph(graph map[string]map[string]int) (map[string]int, map[string]string) {
	startGraph := graph["Start"]

	costsGraph := make(map[string]int)
	parentGraph := make(map[string]string)

	costsGraph["Finish"] = 1000

	for key, value := range startGraph {
		costsGraph[key] = value
		parentGraph[key] = "Start"
	}
	return costsGraph, parentGraph
}
