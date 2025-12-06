package leetcode_232

// https://leetcode.cn/problems/implement-queue-using-stacks/
// 难度：简单

type MyQueue struct {
	in  Stack
	out Stack
}

func Constructor() MyQueue {
	return MyQueue{
		in:  NewStack(),
		out: NewStack(),
	}
}

func (queue *MyQueue) Push(x int) {
	queue.in.Push(x)
}

func (queue *MyQueue) Pop() int {
	if queue.out.Empty() {
		queue.in2out()
	}
	return queue.out.Pop()
}

func (queue *MyQueue) Peek() int {
	if queue.out.Empty() {
		queue.in2out()
	}
	return queue.out.Peek()
}

func (queue *MyQueue) Empty() bool {
	return queue.in.Empty() && queue.out.Empty()
}

func (queue *MyQueue) in2out() {
	for !queue.in.Empty() {
		x := queue.in.Pop()
		queue.out.Push(x)
	}
}

type Stack struct {
	data []int
}

func NewStack() Stack {
	return Stack{
		data: make([]int, 0),
	}
}

func (stack *Stack) Push(x int) {
	stack.data = append(stack.data, x)
}

func (stack *Stack) Pop() int {
	val := stack.Peek()
	stack.data = stack.data[:len(stack.data)-1]
	return val
}

func (stack *Stack) Peek() int {
	return stack.data[len(stack.data)-1]
}

func (stack *Stack) Empty() bool {
	return len(stack.data) == 0
}
