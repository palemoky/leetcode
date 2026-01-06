package kth_largest_element_in_an_array

import (
	"container/heap"
	"math/rand/v2"
	"sort"
)

// 解法一：使用sort包排序后获取倒数第k个元素
func findKthLargestSortedArray(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 使其满足 heap.Interface 接口
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // < 表示最小堆
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 解法二：使用最小堆求解，因为大顶堆的结果不在堆顶
func findKthLargestHeap(nums []int, k int) int {
	// 创建一个最小堆，并初始化
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range nums {
		// 将元素推入堆中
		heap.Push(h, num)
		// 如果堆的大小超过了 k，就弹出堆顶（最小的那个）
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 遍历结束后，堆顶就是第 k 大的元素
	// 注意：Pop 返回的是 any 类型，需要类型断言
	return (*h)[0]
}

// 解法三：快速选择（关注第k个元素，仅处理部分数组），最优解
func findKthLargestQuickSelect(nums []int, k int) int {
	targetIndex := len(nums) - k
	left, right := 0, len(nums)-1

	for {
		// partition 函数会返回 pivot 的最终位置
		pivotIndex := partition(nums, left, right)
		if pivotIndex == targetIndex {
			return nums[pivotIndex]
		} else if pivotIndex < targetIndex {
			// 在右半部分继续寻找
			left = pivotIndex + 1
		} else {
			// 在左半部分继续寻找
			right = pivotIndex - 1
		}
	}
}

// partition 函数（经典的三路快排 partition）
func partition(nums []int, left, right int) int {
	// 随机选择 pivot 以避免最坏情况
	randIndex := left + rand.IntN(right-left+1)
	nums[randIndex], nums[right] = nums[right], nums[randIndex]
	pivot := nums[right]

	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]

	return i
}
