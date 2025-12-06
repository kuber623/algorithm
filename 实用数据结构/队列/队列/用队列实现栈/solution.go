package leetcode_225

// https://leetcode.cn/problems/implement-stack-using-queues/
// 难度：简单
// 题解：两队列解法

type MyStack struct {
	queue1 *Queue
	queue2 *Queue
}

func Constructor() MyStack {
	return MyStack{
		queue1: NewQueue(),
		queue2: NewQueue(),
	}
}

func (stack *MyStack) Push(x int) {
	stack.queue1.Push(x)
}

func (stack *MyStack) Pop() int {
	for stack.queue1.Size() > 1 {
		stack.queue2.Push(stack.queue1.Pop())
	}
	x := stack.queue1.Pop()
	stack.queue1, stack.queue2 = stack.queue2, stack.queue1
	return x
}

func (stack *MyStack) Top() int {
	x := 0
	for !stack.Empty() {
		x = stack.queue1.Pop()
		stack.queue2.Push(x)
	}
	stack.queue1, stack.queue2 = stack.queue2, stack.queue1
	return x
}

func (stack *MyStack) Empty() bool {
	return stack.queue1.Empty()
}

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}

func (queue *Queue) Push(x int) {
	queue.data = append(queue.data, x)
}

func (queue *Queue) Pop() int {
	x := queue.Peek()
	queue.data = queue.data[1:]
	return x
}

func (queue *Queue) Peek() int {
	return queue.data[0]
}

func (queue *Queue) Empty() bool {
	return queue.Size() == 0
}

func (queue *Queue) Size() int {
	return len(queue.data)
}
