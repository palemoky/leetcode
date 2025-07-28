package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ===== 辅助函数 =====
// 为了方便测试，我们创建一些辅助函数

// createLinkedList 从一个整数切片创建链表
func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}

	return head
}

// toSlice 将链表转换回整数切片，方便断言比较
func toSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var res []int
	current := head
	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}
	return res
}

// ===== 测试函数 =====

// TestReverse 测试反转链表函数
func TestReverse(t *testing.T) {
	// 定义测试用例的结构体
	type testCase struct {
		name     string    // 测试用例名称
		input    *ListNode // 输入链表
		expected *ListNode // 期望的输出链表
	}

	// 定义测试用例表
	testCases := []testCase{
		{
			name:     "nil list",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single node list",
			input:    createLinkedList([]int{1}),
			expected: createLinkedList([]int{1}),
		},
		{
			name:     "two nodes list",
			input:    createLinkedList([]int{1, 2}),
			expected: createLinkedList([]int{2, 1}),
		},
		{
			name:     "multiple nodes list (odd)",
			input:    createLinkedList([]int{1, 2, 3, 4, 5}),
			expected: createLinkedList([]int{5, 4, 3, 2, 1}),
		},
		{
			name:     "multiple nodes list (even)",
			input:    createLinkedList([]int{1, 2, 3, 4}),
			expected: createLinkedList([]int{4, 3, 2, 1}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := reverse(tc.input)
			assert.Equal(t, toSlice(tc.expected), toSlice(actual))
		})
	}
}

// TestFindMiddle 统一测试两个查找中间节点的函数
// 因为它们的行为和预期结果应该完全一样
func TestFindMiddle(t *testing.T) {
	// 为测试用例构建链表，并提前获取中间节点的指针
	list1 := createLinkedList([]int{1, 2, 3, 4, 5}) // 中间节点是 3 (l.Next.Next)
	list2 := createLinkedList([]int{1, 2, 3, 4})    // 中间节点是 3 (l.Next.Next) - 按实现逻辑，偶数取第二个
	list3 := createLinkedList([]int{1})             // 中间节点是 1 (l)
	list4 := createLinkedList([]int{1, 2})          // 中间节点是 2 (l.Next)

	// 定义测试用例的结构体
	type testCase struct {
		name     string    // 测试用例名称
		input    *ListNode // 输入链表
		expected *ListNode // 期望的中间节点指针
	}

	testCases := []testCase{
		{
			name:     "nil list",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single node list",
			input:    list3,
			expected: list3, // 期望返回节点1本身
		},
		{
			name:     "odd length list",
			input:    list1,
			expected: list1.Next.Next, // 期望返回节点3
		},
		{
			name:     "even length list",
			input:    list2,
			expected: list2.Next.Next, // 期望返回节点3
		},
		{
			name:     "two nodes list",
			input:    list4,
			expected: list4.Next, // 期望返回节点2
		},
	}

	// 将待测试的函数放入 map 中，方便统一调用
	functionsToTest := map[string]func(*ListNode) *ListNode{
		"ArrayMethod":       findMiddleArray,
		"TwoPointersMethod": findMiddleTwoPointers,
	}

	for funcName, findMiddleFunc := range functionsToTest {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					actual := findMiddleFunc(tc.input)

					// 对于查找中间节点，我们期望返回的是原链表中的某个节点
					// 因此，我们应该比较它们的指针地址是否相同
					// assert.Same() 用于检查两个指针是否指向同一个对象
					assert.Same(t, tc.expected, actual)
				})
			}
		})
	}
}

func TestMergeSortedTwoLists(t *testing.T) {
	// 定义测试用例表，覆盖各种情况
	testCases := []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			name:     "both lists are nil",
			l1:       nil,
			l2:       nil,
			expected: nil,
		},
		{
			name:     "l1 is nil, l2 is not",
			l1:       nil,
			l2:       createLinkedList([]int{1, 2, 3}),
			expected: createLinkedList([]int{1, 2, 3}),
		},
		{
			name:     "l2 is nil, l1 is not",
			l1:       createLinkedList([]int{4, 5, 6}),
			l2:       nil,
			expected: createLinkedList([]int{4, 5, 6}),
		},
		{
			name:     "standard merge with interleaved values",
			l1:       createLinkedList([]int{1, 3, 5}),
			l2:       createLinkedList([]int{2, 4, 6}),
			expected: createLinkedList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:     "merge with different lengths (l1 shorter)",
			l1:       createLinkedList([]int{1, 5}),
			l2:       createLinkedList([]int{2, 3, 4}),
			expected: createLinkedList([]int{1, 2, 3, 4, 5}),
		},
		{
			name:     "merge with different lengths (l2 shorter)",
			l1:       createLinkedList([]int{1, 2, 10}),
			l2:       createLinkedList([]int{5, 6}),
			expected: createLinkedList([]int{1, 2, 5, 6, 10}),
		},
		{
			name:     "l1 values are all smaller than l2",
			l1:       createLinkedList([]int{1, 2, 3}),
			l2:       createLinkedList([]int{4, 5, 6}),
			expected: createLinkedList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:     "merge with duplicate values",
			l1:       createLinkedList([]int{1, 2, 2, 5}),
			l2:       createLinkedList([]int{1, 3, 6}),
			expected: createLinkedList([]int{1, 1, 2, 2, 3, 5, 6}),
		},
		{
			name:     "merge with single node lists",
			l1:       createLinkedList([]int{10}),
			l2:       createLinkedList([]int{4}),
			expected: createLinkedList([]int{4, 10}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := mergeSortedTwoLists(tc.l1, tc.l2)
			assert.Equal(t, toSlice(tc.expected), toSlice(actual))
		})
	}
}

// createLinkedListWithCycle 从一个整数切片创建链表，并能指定环的入口
// vals: 链表节点的值
// cyclePos: 尾节点指向的节点的索引。如果为 -1，则不创建环。
func createLinkedListWithCycle(vals []int, cyclePos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	// 1. 创建所有节点，并用一个切片存起来，方便后续通过索引访问
	nodes := make([]*ListNode, len(vals))
	for i, val := range vals {
		nodes[i] = &ListNode{Val: val}
	}

	// 2. 将节点线性连接起来
	for i := 0; i < len(vals)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// 3. 如果需要，创建环
	if cyclePos != -1 && cyclePos < len(vals) {
		tail := nodes[len(vals)-1]
		tail.Next = nodes[cyclePos] // 尾节点指向环的入口
	}

	return nodes[0] // 返回头节点
}

// ===== 测试是否有环 =====

func TestHasCycle(t *testing.T) {
	// 定义测试用例表
	testCases := []struct {
		name     string
		input    *ListNode
		expected bool
	}{
		{
			name:     "nil list",
			input:    nil,
			expected: false,
		},
		{
			name:     "single node no cycle",
			input:    createLinkedListWithCycle([]int{1}, -1),
			expected: false,
		},
		{
			name:     "multiple nodes no cycle",
			input:    createLinkedListWithCycle([]int{1, 2, 3, 4}, -1),
			expected: false,
		},
		{
			name:     "cycle to head",
			input:    createLinkedListWithCycle([]int{1, 2, 3}, 0),
			expected: true,
		},
		{
			name:     "cycle to middle",
			input:    createLinkedListWithCycle([]int{1, 2, 3, 4, 5}, 2),
			expected: true,
		},
		{
			name:     "cycle to tail (self loop)",
			input:    createLinkedListWithCycle([]int{1, 2}, 1),
			expected: true,
		},
	}

	// 将待测试的函数放入 map 中
	functionsToTest := map[string]func(*ListNode) bool{
		"HashMapMethod":     hasCycleHashMap,
		"TwoPointersMethod": hasCycleTwoPointers,
	}

	for funcName, hasCycleFunc := range functionsToTest {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					actual := hasCycleFunc(tc.input)
					assert.Equal(t, tc.expected, actual)
				})
			}
		})
	}
}

// ===== 测试查找环的入口 =====

func TestFindCycleEntry(t *testing.T) {
	// --- 准备测试数据 ---
	// 对于查找入口的测试，我们需要获得确切的节点指针作为期望结果

	// Case 1: 无环
	noCycleList := createLinkedListWithCycle([]int{1, 2, 3}, -1)

	// Case 2: 环指向头部
	cycleToHeadList := createLinkedListWithCycle([]int{1, 2, 3}, 0)
	expectedHeadEntry := cycleToHeadList // 入口就是头节点

	// Case 3: 环指向中间
	// 我们需要手动构建以获取中间节点的指针
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3} // 期望的入口节点
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n3 // 环指向 n3
	cycleToMiddleList := n1
	expectedMiddleEntry := n3

	// 定义测试用例表
	testCases := []struct {
		name     string
		input    *ListNode
		expected *ListNode
	}{
		{
			name:     "nil list",
			input:    nil,
			expected: nil,
		},
		{
			name:     "no cycle list",
			input:    noCycleList,
			expected: nil,
		},
		{
			name:     "cycle to head",
			input:    cycleToHeadList,
			expected: expectedHeadEntry,
		},
		{
			name:     "cycle to middle",
			input:    cycleToMiddleList,
			expected: expectedMiddleEntry,
		},
	}

	// 将待测试的函数放入 map 中
	functionsToTest := map[string]func(*ListNode) *ListNode{
		"HashMapMethod":         findCycleEnteryHashMap,
		"MathTwoPointersMethod": findCycleEnteryMathTwoPointers,
	}

	for funcName, findEntryFunc := range functionsToTest {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					actual := findEntryFunc(tc.input)
					// 检查返回的指针是否与期望的指针指向同一个内存地址
					assert.Same(t, tc.expected, actual)
				})
			}
		})
	}
}

func TestRemoveNthFromEndTwoPointers(t *testing.T) {
	testCases := []struct {
		name         string // 测试用例的描述
		inputList    []int  // 输入的链表（用切片表示）
		n            int    // 要移除的倒数第 n 个节点
		expectedList []int  // 期望的结果链表（用切片表示）
	}{
		{
			name:         "Remove from the middle",
			inputList:    []int{1, 2, 3, 4, 5},
			n:            2, // 移除倒数第2个节点 (4)
			expectedList: []int{1, 2, 3, 5},
		},
		{
			name:         "Remove the head of the list",
			inputList:    []int{1, 2, 3, 4, 5},
			n:            5, // 移除倒数第5个节点 (1)
			expectedList: []int{2, 3, 4, 5},
		},
		{
			name:         "Remove the tail of the list",
			inputList:    []int{1, 2, 3, 4, 5},
			n:            1, // 移除倒数第1个节点 (5)
			expectedList: []int{1, 2, 3, 4},
		},
		{
			name:         "List with only one node",
			inputList:    []int{1},
			n:            1,       // 移除倒数第1个节点 (1)
			expectedList: []int{}, // 结果为空链表
		},
		{
			name:         "List with two nodes, remove head",
			inputList:    []int{1, 2},
			n:            2, // 移除倒数第2个节点 (1)
			expectedList: []int{2},
		},
		{
			name:         "List with two nodes, remove tail",
			inputList:    []int{1, 2},
			n:            1, // 移除倒数第1个节点 (2)
			expectedList: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := createLinkedList(tc.inputList)
			resultHead := removeNthFromEndTwoPointers(head, tc.n)
			resultSlice := toSlice(resultHead)
			assert.Equal(t, tc.expectedList, resultSlice, "The resulting linked list should match the expected one.")
		})
	}
}

func TestSwapPairs(t *testing.T) {
	testCases := []struct {
		name         string
		inputList    []int
		expectedList []int
	}{
		{
			name:         "Even number of nodes",
			inputList:    []int{1, 2, 3, 4},
			expectedList: []int{2, 1, 4, 3},
		},
		{
			name:         "Odd number of nodes",
			inputList:    []int{1, 2, 3, 4, 5},
			expectedList: []int{2, 1, 4, 3, 5},
		},
		{
			name:         "Empty list",
			inputList:    []int{},
			expectedList: []int{},
		},
		{
			name:         "Single node list",
			inputList:    []int{1},
			expectedList: []int{1},
		},
		{
			name:         "Two nodes list",
			inputList:    []int{1, 2},
			expectedList: []int{2, 1},
		},
	}

	funcsToTest := map[string]func(*ListNode) *ListNode{
		"Iterative": swapPairsIterative,
		"Recursive": swapPairsRecursive,
	}

	for funcName, swapFunc := range funcsToTest {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					head := createLinkedList(tc.inputList)
					resultHead := swapFunc(head)
					resultSlice := toSlice(resultHead)
					assert.Equal(t, tc.expectedList, resultSlice)
				})
			}
		})
	}
}
