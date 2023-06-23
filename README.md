# Go Data Structures and Algorithms

Flowing along with The Primagen's
['Last Algorithms Course You'll need'](https://frontendmasters.com/courses/algorithms/) 
and converting the learning from Typescript examples to Go.

## Algorithms

### Linear search

Given an array find a value's location with that array.

``` go
arr := []int{1, 2, 3, 5, 8, 29, 4, 12}

needle := 8
LinearSearch(arr, needle)
// This will return 4 as the needle 8 is at the fourth index of the array

needle = 1000
LinearSearch(arr, needle)
// This will return -1 as 1000 is not a memeber of the array

// This will return 
```

This algorithm itterates over each value in the array and checks if it is the 
needle.

Linear search has a time complexity of $\mathcal{0}$ 

### Binary search

### Two Crystal balls

### Bubble sort

### Linked list

### Queue

### Stack

### Array list

### Array buffer

### Path finding (recursion)

### Quick sort

### Doubly linked list

### Tree traversal

### Breadth-first search (BFS)

### Binary tree comaprison

### Depth first search (DFS)

### Heap

### BFD on adjacency matrix

### DFS on adjacency list

### Dijkstra's shortest path

### LRU cache

## How to run tests

All tests can be run using `go test` with a verbose output usign `go test -v`.

A specific test can be run using `go test -run <function-name>`. For Binary 
Search this would be `go test -run BinarySearch` again with the option of the
`-v` flag for verbose outputs.
