package container_with_most_water

// Solution 1:
// Time: O(n), Space: O(1)
func maxArea(height []int) int {
	maxArea, left, right := 0, 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right]) // 以较矮的作为高计算面积
		maxArea = max(maxArea, area)
		// 只移动较矮的一边
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
