package merge_sorted_array

// 本题乍看像合并链表，但完全没有关系

// 解法：双指针倒序填充
// 解题关键：倒序填充，以避免修改数组元素导致的移动
// Time: O(m+n), Space: O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, tail := m-1, n-1, m+n-1

	// 只需检查 p2 >= 0，因为 nums2 处理完后，nums1 剩余元素已在正确位置
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[tail] = nums1[p1]
			p1--
		} else {
			nums1[tail] = nums2[p2]
			p2--
		}
		tail--
	}
}
