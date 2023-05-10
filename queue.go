package datastructuresandalgorithms

type Node struct{
    Val any 
    Next *Node
}

type Queue struct{
    length int
    head *Node
    tail *Node
}

func (q *Queue) Enqueue(item any){
    newNode := &Node{
        Val: item,
        Next: nil,
    }

    q.length++

    if q.tail == nil{
        q.tail = newNode
        q.head = newNode
        return
    }

    q.tail.Next = newNode
    q.tail = newNode
}

func (q *Queue) Deque() any{
    if q.head == nil{
        return nil
    }

    item := q.head.Val
    q.head = q.head.Next

    if q.head == nil{
        q.tail = nil
    }

    q.length--

    return item
}

func (q *Queue) Peek() any{
    return q.head.Val
}
