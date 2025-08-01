package add_to_array_form_of_integer

const BASE = 10

// Time: O(max(N, log k)), Space: O(n)
func addToArrayFormBasic(num []int, k int) []int {
	// 0. 把k转为数组形式，注意转换后是逆序形式
	kArr := []int{}
	for k > 0 {
		kArr = append(kArr, k%BASE) // 取一位数
		k /= BASE                   // 去掉这位数
	}
	kArrayReversed := reverse(kArr)

	carry := 0
	ansReversed := []int{}
	// 1. 同时对两个数组逆序遍历
	i, j := len(num)-1, len(kArrayReversed)-1
	// 3. 遍历条件为任何一个元素>=0或carry>0
	for i >= 0 || j >= 0 || carry > 0 {
		// 2. 根据数学规律 sum%base 和 sum/base 对留位与进位处理
		sum := carry
		if i >= 0 {
			sum += num[i]
			i--
		}
		if j >= 0 {
			sum += kArrayReversed[j]
			j--
		}
		carry = sum / BASE
		ansReversed = append(ansReversed, sum%BASE)
	}

	// 4. 对结果反转
	return reverse(ansReversed)
}

func reverse(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

// 解法优化：直接把k作为进位，节省一次遍历和反转
// Time: O(max(N, M)), Space: O(1)
func addToArrayForm(num []int, k int) []int {
	carry := k
	result := []int{}

	for i := len(num) - 1; i >= 0 || carry > 0; i-- {
		sum := carry
		if i >= 0 {
			sum += num[i]
		}

		carry = sum / BASE
		result = append(result, sum%BASE)
	}

	return reverse(result)
}
