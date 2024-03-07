## Quick Sorting 

### Introduction
1. Quick Sorting is a good algorithm to sort a unordered array.
2. It's based on the divide and 'Divide and Conquer algorithm'.
3. This algorithm is heavily relies on recursion for its efficient operation.

### Steps to create a Quick sorting algorithm 
1. First choose a pivot (a random number, most of the time its middle number of the array)
2. then we can separate the elements in the array 2 other arrays , elements that are bigger or smaller than the pivot
3. (OPTIONAL) some time we separate elements in 3 arrays , 1.bigger than pivot, 2.smaller than pivot and 3.equal to pivot
this must be done for arrays that is possible to have repeated random numbers in array
4. then we give both smaller and bigger arrays to two quick sorting function 
5. the two function do these steps too, This is called recursive


### The big O
1. Quick sorting on average time complexity of O(n log n).
2. On worst case that we choose wrong pivot it has O(n<sup>2</sup>) .
3. On best case that the array to sort it has just one, two or three elements it has time complexity of O(1)

### Pros
1. Fast on large datasets
2. Use 'Divide and Conquer algorithm' that makes it easier

### Cons 
1. It has worst time complexity of O(n<sup>2</sup>)

