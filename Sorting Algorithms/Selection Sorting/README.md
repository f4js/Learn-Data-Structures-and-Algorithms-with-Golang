## Selection Sorting

### Introduction
- the Selection Sorting is a simple sorting algorithm for sorting unordered arrays.
- It can order arrays from biggest to smallest or opposite.
- in this algorithm we iterate throw the array and every time find the biggest or smallest number
in the array and appending it to the another array that is going to be ordered.

### The big 'O'
Selection sorting have time complexity of O(n<sup>2</sup>)  so it's a slow sorting algorithm.

### Steps to make a selection sorting

1. Iterate throw the array and find the biggest or smallest number in the array
2. Appending this number to another array
3. Remove the number from the unsorted array
4. Loop again ,Do this until the unsorted array is empty

### Pros
1. It's simple
2. Good for small datasets

### Cons
1. Slow for large datasets
2. Not stable