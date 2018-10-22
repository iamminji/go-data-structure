package stack

type ArrayStack struct {
    array []int
    pos   int
}

func (this *ArrayStack) Push(value int) {

    if this.pos < len(this.array) {
        this.array[this.pos] = value
        this.pos += 1
    } else {
        this.array = append(this.array, value)
        this.pos += 1
    }

}

func (this *ArrayStack) Pop() int {

    if this.pos == 0 && this.pos < len(this.array) {
        return -1
    }
    value := this.array[this.pos-1]
    this.pos -= 1
    return value
}

func (this *ArrayStack) Top() int {
    return this.array[this.pos-1]
}

func (this *ArrayStack) Size() int {
    return this.pos
}
