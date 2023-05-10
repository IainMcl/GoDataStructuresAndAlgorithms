package datastructuresandalgorithms

import (
    "testing"
)

func TestQueue(t *testing.T){
    queue := Queue{}

    queue.Enqueue(10)
    queue.Enqueue(42)
    queue.Enqueue(38)

    if queue.length != 3{
        t.Error("Length not equal to 3")
    }

    if queue.Peek() != 38{
        t.Error()
    }


}
