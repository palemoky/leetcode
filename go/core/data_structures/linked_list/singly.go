package linkedlist

import (
	"fmt"
)

type SinglyNode struct {
	Value int
	Next  *SinglyNode
}

type SinglyList struct {
	Head *SinglyNode
	Len  int
}

func NewSinglyList(arr []int) *SinglyList {
	list := &SinglyList{}
	if len(arr) == 0 {
		return list
	}

	headNode := &SinglyNode{Value: arr[0]}
	list.Head = headNode
	list.Len = 1

	current := headNode
	for i := 1; i < len(arr); i++ {
		newNode := &SinglyNode{Value: arr[i]}
		current.Next = newNode
		current = newNode
		list.Len++
	}

	return list
}

func (list *SinglyList) SinglyAppend(value int) {
	if list.Len == 0 {
		list.Head = &SinglyNode{Value: value}
		list.Len++
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &SinglyNode{Value: value}
	list.Len++
}

func (list *SinglyList) SinglyPrepend(value int) {
	list.Head = &SinglyNode{Value: value, Next: list.Head}
	list.Len++
}

func (list *SinglyList) SinglyInsert(index, value int) {
	if index < 0 || index > list.Len {
		panic(fmt.Errorf("index out of range"))
	}

	dummy := &SinglyNode{Next: list.Head}
	prev := dummy
	for range index {
		prev = prev.Next
	}
	newNode := &SinglyNode{Value: value, Next: prev.Next}
	prev.Next = newNode
	list.Len++

	list.Head = dummy.Next
}

func (list *SinglyList) SinglyDelete(value int) {
	dummy := &SinglyNode{Next: list.Head}

	prev := dummy
	for prev.Next != nil {
		if prev.Next.Value == value {
			prev.Next = prev.Next.Next
			list.Len--
		} else {
			prev = prev.Next
		}
	}
	list.Head = dummy.Next
}

func (list *SinglyList) SinglyDeleteAt(index int) {
	if index < 0 || index >= list.Len {
		panic(fmt.Errorf("index out of range: index=%d, len=%d", index, list.Len))
	}

	dummy := &SinglyNode{Next: list.Head}
	prev := dummy
	for range index {
		prev = prev.Next
	}

	if prev.Next != nil {
		prev.Next = prev.Next.Next
	}

	list.Len--
	list.Head = dummy.Next
}

func (list *SinglyList) SinglyFind(value int) *SinglyNode {
	current := list.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}

	return nil
}

func (list *SinglyList) SinglyGet(index int) *SinglyNode {
	if index < 0 || index >= list.Len {
		return nil
	}

	current := list.Head
	for range index {
		current = current.Next
	}

	return current
}

func ToSlice(head *SinglyNode) []int {
	s := []int{}
	if head == nil {
		return s
	}

	for head != nil {
		s = append(s, head.Value)
		head = head.Next
	}

	return s
}

// 反转链表
func (list *SinglyList) SinglyReverse() {
	var prev, next *SinglyNode // 注意不能写 &SinglyNode{}，这样的初始值是0和nil，只有*SinglyNode的值才是nil

	current := list.Head
	for current != nil {
		next = current.Next // 1. 保存下一个节点
		current.Next = prev // 2. 反转指针
		prev = current      // 3. prev 前进
		current = next      // 4. current 前进
	}

	list.Head = prev // 新的头节点是 prev
}

// 查找中间节点
func (list *SinglyList) SinglyMiddleNodeArray() *SinglyNode {
	if list.Head == nil {
		return nil
	}

	nodes := []*SinglyNode{}
	current := list.Head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	return nodes[len(nodes)/2]
}

// 快慢指针法
// 注意fast指针的非空判断条件
// 注意链表奇偶长度对结果的影响
func (list *SinglyList) SinglyMiddleNodeTwoPointers() *SinglyNode {
	if list.Head == nil {
		return nil
	}

	fast, slow := list.Head, list.Head
	for fast != nil && fast.Next != nil { // 因为fast要跨两步，所以要确定两步都非空
		fast, slow = fast.Next.Next, slow.Next
	}

	return slow
}
