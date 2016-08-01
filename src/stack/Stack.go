package stack

type node struct{
    value interface{}
    next *node
}

type Stack struct {
    top *node
    size int
}

func (stack *Stack) Size() int{
    return stack.size
}

func (stack *Stack) Push(value interface{}){
    stack.top=&node{
        value: value,
        next: stack.top,
    }
    stack.size++
}

func (stack *Stack) Pop() interface{}{
    if stack.size==0 {
        panic("Stack is empty")
    }
    topNode:=stack.top
    stack.top=topNode.next
    stack.size--
    return topNode.value
}

func (stack *Stack) Peek() interface{}{
    if stack.size==0 {
        panic("Stack is empty")
    }

    return stack.top.value
}
