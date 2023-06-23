# Go Data Structures and Algorithms

Flowing along with The Primagen's
['Last Algorithms Course You'll need'](https://frontendmasters.com/courses/algorithms/) 
and converting the learning from Typescript examples to Go.

## <a id="algorithms"></a>Algorithms

### <a id="linear-search"></a>Linear search

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

Linear search has a time complexity of $\mathcal{O}N$ 

### <a id="binary-search"></a>Binary search

This implements the same interface as [Linear search](#linear-search), however,
requiers the input to be an ordered list. 

With this extra constraint binary search allows to skip out checks and reduces
the time complexity of this search to $\mathcal{OlogN}$.

### <a id="two-crystal-balls"></a>Two Crystal balls

### <a id="bubble-sort"></a>Bubble sort

### <a id="linked-list"></a>Linked list

### <a id="queue"></a>Queue

### <a id="stack"></a>Stack

### <a id="array-list"></a>Array list

### <a id="array-buffer"></a>Array buffer

### <a id="path-finding"></a>Path finding (recursion)

### <a id="quick-sort"></a>Quick sort

### <a id="doubly-linked-list"></a>Doubly linked list

### <a id="tree-traversal"></a>Tree traversal

### <a id="breadth-first-search"></a>Breadth-first search (BFS)

### <a id="binary-tree-comaprison"></a>Binary tree comaprison

### <a id="depth-first-search (dfs)"></a>Depth first search (DFS)

### <a id="heap"></a>Heap

### <a id="bfd-on-adjacency-matrix"></a>BFD on adjacency matrix

### <a id="dfs-on-adjacency-list"></a>DFS on adjacency list

### <a id="dijkstras-shortest-path"></a>Dijkstra's shortest path

### <a id="lru-cache"></a>LRU cache

## <a id="how-to-run-tests"></a>How to run tests

All tests can be run using `go test` with a verbose output usign `go test -v`.

A specific test can be run using `go test -run <function-name>`. For Binary 
Search this would be `go test -run BinarySearch` again with the option of the
`-v` flag for verbose outputs.
