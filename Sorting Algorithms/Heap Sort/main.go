package main

import (
	"fmt"
	"math"
)

type Heap struct {
	array    []int
	heapSize int
}

func main() {
	array := []int{3, 5, 13, 15, 16, 18, 9, 43, 19, 23, 42, 11} // our initial array
	heap := BinaryTreeMaker(array)
	HeapSort(&heap)
	fmt.Println(heap.array)

}

// BinaryTreeMaker is a function that will make a max heap binary tree from the given array
func BinaryTreeMaker(array []int) Heap {
	heap := Heap{array: []int{}, heapSize: 0}

	// a loop to add the elements in the heap structure
	for i := 1; i <= len(array); i++ {
		MaxHeapSorter(array[i-1], &heap)
	}
	return heap
}

// MaxHeapSorter function will get pointer of heap and the newNode index after appending the new Node to the heap array
// it will sort the heap array by max heap
func MaxHeapSorter(newNode int, heap *Heap) {

	heap.array = append(heap.array, newNode)
	heap.heapSize++

	// we get the index of the new node we just added
	newNodeIndex := len(heap.array)

	// we calculate the index of the parent of the newNode
	parentIndex := int(math.Floor(float64(newNodeIndex / 2)))

	// we loop and check the parents until we reach the root
	for parentIndex != 0 {
		// if the parent is smaller we change the places
		if heap.array[parentIndex-1] < heap.array[newNodeIndex-1] {
			// here we change the places of the parent and child places
			heap.array[parentIndex-1], heap.array[newNodeIndex-1] = heap.array[newNodeIndex-1], heap.array[parentIndex-1]

			// we calculate the next node and parent node
			newNodeIndex = parentIndex
			parentIndex = int(math.Floor(float64(newNodeIndex / 2)))

		} else { // if the parent is bigger than the new node we calculate the next parent and node
			newNodeIndex = parentIndex
			parentIndex = int(math.Floor(float64(newNodeIndex / 2)))
		}
	}
}

// HeapSort is a function that will change tha place of root and place of the last element in the heap array.
// then we will use MaxHeapSorter to sort the slice again
func HeapSort(heap *Heap) *Heap {
	// a loop to get from last to root of the heap array
	for i := len(heap.array); i >= 1; i-- {
		// we replace the root with last element in the heap array
		heap.array[0], heap.array[i-1] = heap.array[i-1], heap.array[0]
		heap.heapSize--
		//fmt.Println("before root to child:", heap.array)
		RootToChildMaxHeapSorter(heap)
		//fmt.Println("after:", heap.array)

	}

	return heap
}

// RootToChildMaxHeapSorter is a function that will sort the heap array from the root
func RootToChildMaxHeapSorter(heap *Heap) {
	index := 1
	var largestChild int
	for {
		leftChild := 2 * index
		rightChild := (2 * index) + 1
		// if node have right child that is mean it has left child too
		if rightChild <= heap.heapSize { // we check if it does right child
			if heap.array[leftChild-1] > heap.array[rightChild-1] { // we here check which child is bigger
				largestChild = leftChild
			} else {
				largestChild = rightChild
			}

		} else if leftChild <= heap.heapSize { // we check if it does left child
			largestChild = leftChild

		} else { // if it doesn't have any child then we break
			break
		}

		// we replace the largest child with the index node if it is bigger
		if heap.array[largestChild-1] > heap.array[index-1] {
			heap.array[index-1], heap.array[largestChild-1] = heap.array[largestChild-1], heap.array[index-1]
			index = largestChild
		} else { // we break if largest child is not bigger than the index node
			break
		}
	}
}
