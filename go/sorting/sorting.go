package sorting

import (
	"container/heap"
	"math"
)

// 所有排序算法均为升序实现

// 冒泡排序：不断两两比较交换，直到所有元素有序
// 优化点：已排序区可跳过；全部有序可提前结束遍历
// Time: O(n^2)
func bubble(nums []int) []int {
	n := len(nums)
	// 空或 1 个元素都无需排序
	if n < 2 {
		if n == 0 {
			return []int{}
		}

		return nums
	}

	// 外层循环控制轮数，最后一个数不需要比较
	for i := 0; i < n-1; i++ {
		swapped := false

		// 内层循环进行比较交换
		// 由于每次外层循环结束后最大的元素会被移动到数组的末尾，因此内层循环的范围可以逐渐缩小
		for j := 1; j < n-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
				swapped = true
			}
		}

		// 未发生交换则已排好序，提前结束比较
		if !swapped {
			break
		}
	}

	return nums
}

// 选择排序：在待排序区不停选择最小值，然后与待排序区首个元素交换（已排序区末尾或待排序区头部）。不断循环，即可对所有元素排序
// Time: O(n^2)
func selection(nums []int) []int {
	n := len(nums)
	if n < 2 {
		if n == 0 {
			return []int{}
		}

		return nums
	}

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}

// 插入排序：从第二个元素开始，不断将之后的元素放在之前的已排序位置
// Time: O(n^2)
func insertion(nums []int) []int {
	n := len(nums)
	if n < 2 {
		if n == 0 {
			return []int{}
		}

		return nums
	}

	// [0,i)为已排序区，[i,n)为待排序区
	for i := 1; i < n; i++ {
		key := nums[i] // key 为从桌上拿起的新牌
		j := i - 1
		// j 负责逆序遍历已排序区，为 key 找到插入位置
		for j >= 0 && nums[j] > key {
			// 向右移动元素，腾出插入空间
			nums[j+1] = nums[j]
			j--
		}
		// 将 key 插入正确位置
		nums[j+1] = key
	}

	return nums
}

// 快速排序
func quick(nums []int) []int {
	n := len(nums)
	if n < 2 {
		if n == 0 {
			return []int{}
		}

		return nums
	}

	// 调用私有的递归辅助函数
	quickSortRecursive(nums, 0, n-1)

	return nums
}

// quickSortRecursive 是递归的核心
func quickSortRecursive(nums []int, left, right int) {
	if left >= right {
		return
	}

	// 分区，并拿到 pivot 的最终位置
	pivotIndex := quickSortPartition(nums, left, right)

	// 递归地对左右两部分进行排序
	quickSortRecursive(nums, left, pivotIndex-1)
	quickSortRecursive(nums, pivotIndex+1, right)
}

// quickSortPartition 负责分区操作，并返回 pivot 的索引
func quickSortPartition(nums []int, left, right int) int {
	pivotValue := nums[left] // 选择第一个元素作为 pivot
	l, r := left, right

	for l < r {
		// 从右向左找第一个小于 pivot 的数
		for l < r && nums[r] >= pivotValue {
			r--
		}

		nums[l] = nums[r]
		// 从左向右找第一个大于 pivot 的数
		for l < r && nums[l] <= pivotValue {
			l++
		}

		nums[r] = nums[l]
	}

	// 将 pivot 放回正确的位置
	nums[l] = pivotValue

	return l
}

// 归并排序
func merge(nums []int) []int {
	if len(nums) < 2 {
		if nums == nil {
			return []int{}
		}

		return nums
	}

	mergeSortRecursive(nums, 0, len(nums)-1)

	return nums
}

// mergeSortRecursive 是递归的核心
func mergeSortRecursive(nums []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	mergeSortRecursive(nums, left, mid)
	mergeSortRecursive(nums, mid+1, right)
	mergeHalves(nums, left, mid, right)
}

// mergeHalves 负责合并两个已排序的子数组
func mergeHalves(nums []int, left, mid, right int) {
	// 创建一个临时切片来存储合并后的结果
	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	// 比较左右两部分，将较小的元素放入 tmp
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}

	// 将剩余的元素（如果 L 或 R 部分有剩下）拷贝到 tmp
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}

	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}

	// 将排好序的 tmp 内容拷贝回原始的 nums 切片的对应位置
	copy(nums[left:right+1], tmp)
}

// 希尔排序
func shell(nums []int) []int {
	n := len(nums)
	if n < 2 {
		if n == 0 {
			return []int{}
		}

		return nums
	}

	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := nums[i]
			j := i
			for j >= gap && nums[j-gap] > temp {
				nums[j] = nums[j-gap]
				j -= gap
			}
			nums[j] = temp
		}
		gap /= 2
	}

	return nums
}

// 堆排序
func heapSorting(nums []int) []int {
	n := len(nums)
	if n < 2 {
		if n == 0 {
			return []int{}
		}

		return nums
	}

	h := &intHeap{}
	for _, v := range nums {
		heap.Push(h, v)
	}

	for i := range nums {
		nums[i] = heap.Pop(h).(int)
	}

	return nums
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 计数排序（假设 nums 非负且最大值不大）
func counting(nums []int) []int {
	n := len(nums)
	if n < 2 {
		if n == 0 {
			return []int{}
		}

		return nums
	}

	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	count := make([]int, maxVal+1)
	for _, v := range nums {
		count[v]++
	}

	idx := 0
	for i, c := range count {
		for c > 0 {
			nums[idx] = i
			idx++
			c--
		}
	}

	return nums
}

// 桶排序（假设 nums 非负且分布均匀）
func bucket(nums []int) []int {
	n := len(nums)
	if n < 2 {
		if n == 0 {
			return []int{}
		}

		return nums
	}

	maxVal, minVal := nums[0], nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}

		if v < minVal {
			minVal = v
		}
	}

	bucketNum := len(nums)
	buckets := make([][]int, bucketNum)
	interval := int(math.Ceil(float64(maxVal-minVal+1) / float64(bucketNum)))
	for _, v := range nums {
		idx := (v - minVal) / interval
		buckets[idx] = append(buckets[idx], v)
	}

	idx := 0
	for _, b := range buckets {
		insertion(b)
		for _, v := range b {
			nums[idx] = v
			idx++
		}
	}

	return nums
}

// 基数排序（假设 nums 非负整数）
func radix(nums []int) []int {
	n := len(nums)
	if n < 2 {
		if n == 0 {
			return []int{}
		}

		return nums
	}

	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	exp := 1
	buf := make([]int, n)
	for maxVal/exp > 0 {
		count := make([]int, 10)
		for i := range n {
			count[(nums[i]/exp)%10]++
		}

		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}

		for i := n - 1; i >= 0; i-- {
			digit := (nums[i] / exp) % 10
			buf[count[digit]-1] = nums[i]
			count[digit]--
		}
		copy(nums, buf)
		exp *= 10
	}

	return nums
}
