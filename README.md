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
// This will return -1 as 1000 is not a member of the array
```

This algorithm iterates over each value in the array and checks if it is the 
needle.

Linear search has a time complexity of $\mathcal{O}N$ 

### <a id="binary-search"></a>Binary search

This implements the same interface as [Linear search](#linear-search), however,
requires the input to be an ordered list. 

With this extra constraint binary search allows to skip out checks and reduces
the time complexity of this search to $\mathcal{OlogN}$.

### <a id="two-crystal-balls"></a>Two Crystal balls

Two crystal balls solves problem using different forms of 
[Linear search](#linear-search).

The problem is: Given two crystal balls and a tall building with many floors,
find the maximum height that a ball can be dropped from without breaking. Both
balls are allowed to be broken to find this answer.

```go
testBuilding := []bool{false, false, false, false, true, true, true}
// The test building represents each floor of the building with each index of 
// array. The ground floor being testBuilding[0], first floor being 
// testBuilding[1], and so on.

TwoCrystalBalls(testBuilding)
// In this instance will return 3 as that is the highest floor that the ball
// can be dropped from without smashing.
```

This could be done solely using [Liner search](#linear-search) until a ball 
breaks. However, we have been given two balls so can afford to break one to 
improve speed.

This can be done by performed by stepping through the array in large steps 
until one of the balls breaks. When we have found that index go back to the 
last known point of safety +1 and then linearly search until the second ball 
breaks.

To maximise speed the size of the steps need to be changed. With steps that are
too large result in the linear search being too long. In the extreme case 
where the steps are the size of the array, one step would be taken and found to 
break the ball. Then linearly searching through the whole array anyway to find 
where the second ball breaks. Conversly if the steps are too small this will 
again approach linear search time. The balance is found by taking steps that
are $\sqrt{ N }$.


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
