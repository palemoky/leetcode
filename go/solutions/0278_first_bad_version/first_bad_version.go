package first_bad_version

var bad int

func isBadVersion(version int) bool {
	return version >= bad
}

// 可用二分查找来加速搜索
func firstBadVersion(n int) int {
	// 注意此处的查找不是通过数组的索引，版本号从 1 开始
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else { // isBadVersion(mid) 为 false 时，说明第一个错误版本在右侧区间
			left = mid + 1
		}
	}

	return left
}
