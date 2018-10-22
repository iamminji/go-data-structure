package stack

import (
    "testing"
)

func TestProcedure(t *testing.T) {
    mstack := MinStack{}

    /*
        100 -> 99 -> 21 -> 32 -> 3 -> 2 -> 10 -> 1
    */
    mstack.Push(100)
    mstack.Push(99)
    mstack.Push(21)
    mstack.Push(32)
    mstack.Push(3)
    mstack.Push(2)
    mstack.Push(10)
    mstack.Push(1)

    if mstack.Size() != 8 {
        t.Error("size error", mstack.Size())
    }

    p1 := mstack.Pop()
    if p1 != 1 {
        t.Error("error", p1)
    }

    if mstack.Size() != 7 {
        t.Error("size error", mstack.Size())
    }

    m1 := mstack.GetMin()
    if m1 != 2 {
        t.Error("error", m1)
    }

    p2 := mstack.Pop()
    if p2 != 10 {
        t.Error("error", p2)
    }

    m2 := mstack.GetMin()
    if m2 != 2 {
        t.Error("error", m2)
    }

    for mstack.Size() > 0 {
        t.Log("current MinValue :", mstack.GetMin(), "| current Value :", mstack.Pop())
    }

    if mstack.GetMin() != -1 {
        t.Error("error")
    }

}
