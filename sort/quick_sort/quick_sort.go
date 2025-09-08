package quick_sort

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// 指定第一个元素作为基准值
	pivot := nums[0]

	// 遍历数组并将小于基准值的元素移动到左侧
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[0] = nums[0], nums[i]

	QuickSort(nums[:i])
	QuickSort(nums[i+1:])
}
