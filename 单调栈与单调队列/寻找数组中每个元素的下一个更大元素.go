package mono

// NextGreaterElement 寻找数组中每个元素的下一个更大元素
func NextGreaterElement(nums []int) []int {
	nextGreater := make([]int, len(nums))

	// 从右向左遍历数组， 构造单调递减栈，单调栈记录当前元素右侧可能成为答案的元素
	monostack := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(monostack) > 0 && nums[i] >= monostack[len(monostack)-1] {
			monostack = monostack[:len(monostack)-1]
		}

		if len(monostack) == 0 {
			nextGreater[i] = -1
		} else {
			nextGreater[i] = monostack[len(monostack)-1]
		}

		monostack = append(monostack, nums[i])
	}

	return nextGreater
}
