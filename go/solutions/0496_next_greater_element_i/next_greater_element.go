package next_greater_element_i

// Time: O(n), Space: O(n)
func nextGreaterElementStackLeftToRight(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)

	// 根据 nums1 的长度初始化 ans
	ans := make([]int, n1)
	for i := range n1 {
		ans[i] = -1
	}

	// 根据 nums1 中的元素获取索引，以便在 ans 的对应位置更新
	idx := make(map[int]int, n1)
	for i, v := range nums1 {
		idx[v] = i
	}

	st := []int{}
	for _, v2 := range nums2 {
		for len(st) > 0 && v2 > st[len(st)-1] {
			top := len(st) - 1
			ans[idx[st[top]]] = v2
			st = st[:top]
		}

		if _, ok := idx[v2]; ok {
			st = append(st, v2)
		}
	}

	return ans
}

// Time: O(n), Space: O(n)
func nextGreaterElementStackRightToLeft(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)

	// 根据 nums1 的长度初始化 ans
	ans := make([]int, n1)
	for i := range n1 {
		ans[i] = -1
	}

	// 根据 nums1 中的元素获取索引，以便在 ans 的对应位置更新
	idx := make(map[int]int, n1)
	for i, v := range nums1 {
		idx[v] = i
	}

	st := []int{}
	for i := n2 - 1; i >= 0; i-- {
		v2 := nums2[i]
		for len(st) > 0 && v2 >= st[len(st)-1] {
			st = st[:len(st)-1]
		}

		if len(st) > 0 {
			if j, ok := idx[v2]; ok {
				ans[j] = st[len(st)-1]
			}
		}

		st = append(st, v2)
	}

	return ans
}
