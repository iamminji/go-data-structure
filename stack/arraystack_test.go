package stack

import (
    "testing"
)

func TestArrayStack_Push(t *testing.T) {
    as := ArrayStack{}
    as.Push(3)
    as.Push(2)

    if as.Size() != 2 {
        t.Error("error", as.Size())
    }

    as.Push(5)
    if as.Size() != 3 {
        t.Error("error", as.Size())
    }
}

func TestArrayStack_Pop(t *testing.T) {
    as := ArrayStack{}
    as.Push(1)
    as.Push(2)
    as.Push(3)
    as.Push(4)
    as.Push(5)

    for i := 5; i > 0; i-- {
        temp := as.Pop()
        if i != temp {
            t.Error("error", temp)
        }
    }

    if as.Pop() != -1 {
        t.Error(as)
    }
}

func TestArrayStack_Top(t *testing.T) {
    as := ArrayStack{}
    as.Push(1)
    as.Push(2)
    as.Push(3)
    as.Push(4)

    if as.Top() != 4 {
        t.Error("error")
    }
}

func TestAllArrayStack(t *testing.T) {
    as := ArrayStack{}

    as.Push(1)
    as.Push(2)
    as.Push(10)
    as.Push(21)
    as.Push(32)

    if as.Pop() != 32 {
        t.Error("error")
    }

    if as.Pop() != 21 {
        t.Error("error")
    }

    if as.Size() != 3 {
        t.Error("error")
    }

    as.Push(11)

    if as.Size() != 4 {
        t.Error("error")
    }

    if as.Pop() != 11 {
        t.Error("error")
    }

    as.Push(100)

    if as.Top() != 100 {
        t.Error("error")
    }

    as.Push(101)
    as.Push(301)
    if as.Size() != 6 {
        t.Error("error", as.Size())
    }

}
