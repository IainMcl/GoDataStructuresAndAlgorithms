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
where the second ball breaks. Conversely if the steps are too small this will
again approach linear search time. The balance is found by taking steps that
are $\sqrt{ N }$ in size.


### <a id="bubble-sort"></a>Bubble sort

Sorts an array of numbers by progressively 'bubbling' the largest number to the
top (end) of the array.

The array is iterated through and if the next index is smaller than the
current index being checked then the values are swapped in place. This moves
the larger value towards the top of the array.

At the end of each iteration through the array the highest value will be at
the end index. This allows for a time saving measure by on the next iteration
only looping over the length of the array ($N$) - 1, then $N$-2, ...

The time complexity of this algorithm can range depending on the current state
of the array from $\mathcal{O(N)}$ to $\mathcal{O(N^2)}$. If the array is
already sorted or if the array only has the largest value out of place then
this will be confirmed to be sorted in one pass of the array, hence
$\mathcal{ON}$. And the worst case being if the array was completely inverted
from its sorted order; here every value will need to be swapped and every
iteration will need to be completed to confirm that the array has been sorted.

The example shown in [bubble_sort.go](bubble_sort.go) could be improved by
counting the number of swaps in each pass of the array. If in an array pass
there has been no swaps then the array is sorted and can returned.

### <a id="linked-list"></a>Linked list

A linked list is a data structure that can show connections between certain
pieces of data. This typically consists of a series of nodes the each contain a
value and a link to another node. The links can also link to the next and
previous nodes to form a doubly linked list.

#### Node implementation
```go
type Node struct {
    Val any
    Next *Node
    // Optional
    Prev *Node
}
```

### <a id="queue"></a>Queue

A queue can be implemented using a [linked list](#linked-list). A queue adds
some additional functionality to the linked list.

These methods are typically:

* Enqueue - Add a new value to the tail of the linked list

* Deque - Receive the value from the head of the linked list and remove that
value from the linked list

* Peek - View the value stored at the head of the linked list while leaving the
node in place.

Queues can be used for a veriety of things with a common commercial
implementation being that used in [RabbieMQ](https://www.rabbitmq.com/).
Technologies like this can provide the links between different services in a
micro service infrastructure. A queue can also be used to allow for easier
scaling of applications. If there is more demand for a service than one
instance of that service can deal with itself then another instance can be spun
up and read inputs from the same queue.

### <a id="stack"></a>Stack

A stack is another implementation of a [linked list](#linked#list) and can be
thought of as a reversed [queue](#queue). In a stack implementation rather than
new values being added to the tail of the linked list they are instead pushed
on to the head of the linked list and similarly to dequeing values are popped
from the head of the list.

Stacks are used in many places. On of the most common places that a programmer
will have seen a stack implementation is when an error is thrown and a
'call stack' is shown with the error. This 'call stack' shows the functions and
methods in the order that they were called before the error was thrown.

When a new function or method is called that is pushed on to the 'call stack'
and popped off when the value returns successfully. This will preserve the
stack up to that point.

```go
func one(){
    two()
}
func two(){
    // Error
}

func main(){
    one()
}
```

Here the call stack will start in `main`. `main` will be pushed on to the empty
stack. `main` calls `one` so one is also pushed to the stack

`main -> one`

`one` in turn calls `two`

`main` -> `one` -> `two`

`two` raises an error so the above call stack is returned.

If `two` were to not give an error but instead the error was to be raised
after `two` had completed then `two` would be popped from the stack and the new
returned stack won't include `two`.

```go
func one(){
    two()
    three()
}
func two(){}

func three(){
    // Error
}

func main(){
    one()
}
```

Here the call stack will grow as

`main` push main

`main` -> `one` push one

`main` -> `one` -> `two` push two

`main` -> `one` pop two

`main` -> `one` -> `three` push three and return


### <a id="array-list"></a>Array list

### <a id="array-buffer"></a>Array buffer

### <a id="path-finding"></a>Path finding (recursion)

### <a id="quick-sort"></a>Quick sort

### <a id="doubly-linked-list"></a>Doubly linked list

### <a id="tree-traversal"></a>Tree traversal

### <a id="breadth-first-search"></a>Breadth-first search (BFS)

### <a id="binary-tree-comparison"></a>Binary tree comparison

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
