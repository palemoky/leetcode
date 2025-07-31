package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDoublyList(t *testing.T) {
	t.Parallel()

	t.Run("Create from empty array", func(t *testing.T) {
		list := NewDoublyList([]int{})
		assert.Equal(t, 0, list.Len)
		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
	})

	t.Run("Create from non-empty array", func(t *testing.T) {
		t.Parallel()

		array := []int{10, 20, 30}
		list := NewDoublyList(array)

		assert.Equal(t, len(array), list.Len)
		assert.Equal(t, 10, list.Head.Value)
		assert.Equal(t, 30, list.Tail.Value)
		assert.Equal(t, array, toSlice(list))
		assert.Equal(t, reverseSlice(array), toSliceReverse(list))
	})
}

func TestDoublyAppend(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		initial  []int
		value    int
		expected []int
	}{
		{"Append to empty list", []int{}, 10, []int{10}},
		{"Append to non-empty list", []int{10, 20}, 30, []int{10, 20, 30}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			list := NewDoublyList(tc.initial)
			list.DoublyAppend(tc.value)

			assert.Equal(t, len(tc.expected), list.Len)
			assert.Equal(t, tc.expected[0], list.Head.Value)
			assert.Equal(t, tc.expected[len(tc.expected)-1], list.Tail.Value)
			assert.Equal(t, tc.expected, toSlice(list))
			assert.Equal(t, reverseSlice(tc.expected), toSliceReverse(list))
		})
	}
}

func TestDoublyPrepend(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		initial  []int
		value    int
		expected []int
	}{
		{"Prepend to empty list", []int{}, 10, []int{10}},
		{"Prepend to non-empty list", []int{20, 30}, 10, []int{10, 20, 30}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			list := NewDoublyList(tc.initial)
			list.DoublyPrepend(tc.value)

			assert.Equal(t, len(tc.expected), list.Len)
			assert.Equal(t, tc.value, list.Head.Value)
			assert.Equal(t, tc.expected[len(tc.expected)-1], list.Tail.Value)
			assert.Equal(t, tc.expected, toSlice(list))
			assert.Equal(t, reverseSlice(tc.expected), toSliceReverse(list))
		})
	}
}

func TestDoublyInsert(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		initial     []int
		index       int
		value       int
		shouldPanic bool
		expected    []int
	}{
		{"Insert into empty list", []int{}, 0, 10, false, []int{10}},
		{"Insert at head", []int{20, 30}, 0, 10, false, []int{10, 20, 30}},
		{"Insert in middle", []int{10, 30}, 1, 20, false, []int{10, 20, 30}},
		{"Insert at tail", []int{10, 20}, 2, 30, false, []int{10, 20, 30}},
		{"Panic on negative index", []int{10, 20}, -1, 5, true, nil},
		{"Panic on index > length", []int{10, 20}, 3, 30, true, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			list := NewDoublyList(tc.initial)
			insertFunc := func() {
				list.DoublyInsert(tc.value, tc.index)
			}

			if tc.shouldPanic {
				assert.Panics(t, insertFunc)
			} else {
				assert.NotPanics(t, insertFunc)
				assert.Equal(t, len(tc.expected), list.Len)
				assert.Equal(t, tc.expected, toSlice(list))
				assert.Equal(t, reverseSlice(tc.expected), toSliceReverse(list))
			}
		})
	}
}

func TestDoublyDeleteAt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		initial  []int
		index    int
		expected []int
	}{
		{"Delete from single node list", []int{10}, 0, []int{}},
		{"Delete head", []int{10, 20, 30}, 0, []int{20, 30}},
		{"Delete tail", []int{10, 20, 30}, 2, []int{10, 20}},
		{"Delete middle", []int{10, 20, 30, 40}, 1, []int{10, 30, 40}},
		{"Invalid negative index does nothing", []int{10, 20}, -1, []int{10, 20}},
		{"Invalid large index does nothing", []int{10, 20}, 2, []int{10, 20}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			list := NewDoublyList(tc.initial)
			list.DoublyDeleteAt(tc.index)

			assert.Equal(t, len(tc.expected), list.Len)
			assert.Equal(t, tc.expected, toSlice(list))
			if len(tc.expected) > 0 {
				assert.Equal(t, reverseSlice(tc.expected), toSliceReverse(list))
			}
		})
	}
}

func TestDoublyDelete(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		initial  []int
		value    int
		expected []int
	}{
		{"Delete non-existent value", []int{10, 20, 30}, 40, []int{10, 20, 30}},
		{"Delete head", []int{10, 20, 30}, 10, []int{20, 30}},
		{"Delete tail", []int{10, 20, 30}, 30, []int{10, 20}},
		{"Delete middle", []int{10, 20, 30}, 20, []int{10, 30}},
		{"Delete only node", []int{10}, 10, []int{}},
		{"Delete multiple occurrences", []int{10, 20, 10, 30, 10}, 10, []int{20, 30}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			list := NewDoublyList(tc.initial)
			list.DoublyDelete(tc.value)

			assert.Equal(t, len(tc.expected), list.Len)
			assert.Equal(t, tc.expected, toSlice(list))
			if len(tc.expected) > 0 {
				assert.Equal(t, reverseSlice(tc.expected), toSliceReverse(list))
			}
		})
	}
}

func TestDoublyGet(t *testing.T) {
	t.Parallel()

	list := NewDoublyList([]int{10, 20, 30, 40, 50})

	testCases := []struct {
		name      string
		index     int
		expectNil bool
		expectVal int
	}{
		{"Get from first half", 1, false, 20},
		{"Get from second half", 3, false, 40},
		{"Get head", 0, false, 10},
		{"Get tail", 4, false, 50},
		{"Get invalid negative index", -1, true, 0},
		{"Get invalid large index", 5, true, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			node := list.DoublyGet(tc.index)
			if tc.expectNil {
				assert.Nil(t, node)
			} else {
				assert.NotNil(t, node)
				assert.Equal(t, tc.expectVal, node.Value)
			}
		})
	}
}

func TestDoublyFind(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		initial   []int
		target    int
		expectNil bool
		expectVal int
	}{
		{name: "Find existing value", initial: []int{10, 20, 30}, target: 20, expectNil: false, expectVal: 20},
		{name: "Find non-existent value", initial: []int{10, 20, 30}, target: 40, expectNil: true},
		{name: "Find in empty list", initial: []int{}, target: 10, expectNil: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			list := NewDoublyList(tc.initial)
			node := list.DoublyFind(tc.target)

			if tc.expectNil {
				assert.Nil(t, node)
			} else {
				assert.NotNil(t, node)
				assert.Equal(t, tc.expectVal, node.Value)
			}
		})
	}
}
