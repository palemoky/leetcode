package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewSinglyList tests the creation of a new list from a slice.
func TestNewSinglyList(t *testing.T) {
	assert := assert.New(t)

	t.Run("Create from empty slice", func(t *testing.T) {
		list := NewSinglyList([]int{})
		assert.Nil(list.Head, "NewLinkedList with empty slice should return nil")
	})

	t.Run("Create from non-empty slice", func(t *testing.T) {
		list := NewSinglyList([]int{1, 2, 3})
		assert.NotNil(list)
		assert.Equal([]int{1, 2, 3}, toSlice(list), "The list should match the initial slice")
	})
}

// TestAppend tests adding a node to the end of the list.
func TestAppend(t *testing.T) {
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

			list := NewSinglyList(tc.initial)
			list.SinglyAppend(tc.value)

			assert.Equal(t, len(tc.expected), list.Len)
			assert.Equal(t, tc.expected[0], list.Head.Value)
			assert.Equal(t, tc.expected, toSlice(list))
		})
	}
}

// TestPrepend tests adding a node to the beginning of the list.
func TestPrepend(t *testing.T) {
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

			list := NewSinglyList(tc.initial)
			list.SinglyPrepend(tc.value)

			assert.Equal(t, len(tc.expected), list.Len)
			assert.Equal(t, tc.value, list.Head.Value)
			assert.Equal(t, tc.expected, toSlice(list))
		})
	}
}

// TestInsert tests inserting a node at a specific index.
func TestInsert(t *testing.T) {
	testCases := []struct {
		name        string
		array       []int
		index       int
		value       int
		expectPanic bool
		expectVal   []int
	}{
		// Cases where we expect a node
		{"Insert into empty list at index 0", []int{}, 0, 10, false, []int{10}},
		{"Insert at head", []int{20, 30}, 0, 10, false, []int{10, 20, 30}},
		{"Insert in middle", []int{10, 30}, 1, 20, false, []int{10, 20, 30}},
		{"Insert at end", []int{10, 20}, 2, 30, false, []int{10, 20, 30}},
		// Cases where we expect nil
		{"Index out of range (negative)", []int{10, 20}, -1, 5, true, nil},
		{"Index out of range (equal to length)", []int{10, 20}, 3, 30, true, nil},
	}

	assert := assert.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			list := NewSinglyList(tc.array)

			if tc.expectPanic {
				assert.Panics(func() {
					list.SinglyInsert(tc.index, tc.value)
				}, "The code panic as expected")
			} else {
				list.SinglyInsert(tc.index, tc.value)

				assert.Equal(len(tc.expectVal), list.Len, "List length is incorrect")

				if tc.expectVal == nil {
					assert.Nil(list.Head, "List head should be nil")
				} else {
					assert.NotNil(list.Head, "List head should not be nil")
					assert.Equal(tc.expectVal, ToSlice(list.Head), "List content is incorrect")
				}
			}
		})
	}
}

// TestDelete tests deleting the first occurrence of a value.
func TestDelete(t *testing.T) {
	testCases := []struct {
		name      string
		array     []int
		value     int
		expectNil bool
		expectVal []int
	}{
		// Cases where we expect a node
		{"Delete head", []int{10, 20, 30}, 10, false, []int{20, 30}},
		{"Delete middle node", []int{10, 20, 30}, 20, false, []int{10, 30}},
		{"Delete tail node", []int{10, 20, 30}, 30, false, []int{10, 20}},
		{"Delete value not present", []int{10, 20, 30}, 40, false, []int{10, 20, 30}},
		{"Delete duplicates value", []int{10, 20, 10, 30}, 10, false, []int{20, 30}},

		// Cases where we expect nil
		{"Delete from nil list", []int{}, 10, true, nil},
		{"Delete from single node list", []int{10}, 10, true, nil},
	}

	assert := assert.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			list := NewSinglyList(tc.array)
			list.SinglyDelete(tc.value)

			if tc.expectNil {
				assert.Nil(list.Head)
			} else {
				// We can add more assertions for the non-nil case
				assert.NotNil(list.Head)
				if list.Head != nil { // Check for nil to prevent panic on the next line
					assert.Equal(tc.expectVal, ToSlice(list.Head))
				}
			}
		})
	}
}

// TestDeleteAt tests deleting a node at a specific index.
func TestDeleteAt(t *testing.T) {
	testCases := []struct {
		name        string
		array       []int
		index       int
		expectPanic bool
		expectedVal []int
	}{
		// Cases where we expect a node
		{"Delete at index 0", []int{10, 20, 30}, 0, false, []int{20, 30}},
		{"Delete at middle index", []int{10, 20, 30}, 1, false, []int{10, 30}},
		{"Delete at last index", []int{10, 20, 30}, 2, false, []int{10, 20}},
		{"Delete the only node", []int{10}, 0, false, nil},
		// Cases where we expect panic
		{"Index out of range (negative)", []int{10, 20}, -1, true, nil},
		{"Index out of range (equal to length)", []int{10, 20}, 2, true, nil},
		{"Index out of range (on empty list)", []int{}, 0, true, nil},
	}

	assert := assert.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			list := NewSinglyList(tc.array)

			if tc.expectPanic {
				assert.Panics(func() {
					list.SinglyDeleteAt(tc.index)
				}, "The code panic as expected")
			} else {
				list.SinglyDeleteAt(tc.index)

				assert.Equal(len(tc.expectedVal), list.Len, "List length is incorrect")

				if tc.expectedVal == nil {
					assert.Nil(list.Head, "List head should be nil")
				} else {
					assert.NotNil(list.Head, "List head should not be nil")
					assert.Equal(tc.expectedVal, ToSlice(list.Head), "List content is incorrect")
				}
			}
		})
	}
}

// TestFind tests finding a node with a specific value.
func TestFind(t *testing.T) {
	testCases := []struct {
		name        string
		array       []int
		value       int
		expectNil   bool
		expectedVal int
	}{
		// Cases where we expect a node
		{"Find existing element", []int{10, 20, 30}, 20, false, 20},
		// Cases where we expect nil
		{"Find non-existent element", []int{10, 20, 30}, 40, true, 0},
		{"Find in nil list", nil, 10, true, 0},
	}

	assert := assert.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			list := NewSinglyList(tc.array)
			node := list.SinglyFind(tc.value)

			if tc.expectNil {
				assert.Nil(node, "Expected a nil node but got one")
			} else {
				// We can add more assertions for the non-nil case
				assert.NotNil(node, "Expected a node but got nil")
				if node != nil { // Check for nil to prevent panic on the next line
					assert.Equal(tc.expectedVal, node.Value, "Node value does not match expected value")
				}
			}
		})
	}
}

// TestGet tests getting a node at a specific index.
func TestGet(t *testing.T) {
	testCases := []struct {
		name        string
		array       []int
		index       int
		expectNil   bool // A flag to indicate if we expect a nil result
		expectedVal int  // Only relevant if expectNil is false
	}{
		// Cases where we expect a node
		{"Get at index 0", []int{10, 20, 30}, 0, false, 10},
		{"Get at valid middle index", []int{10, 20, 30}, 1, false, 20},
		{"Get at last index", []int{10, 20, 30}, 2, false, 30},
		// Cases where we expect nil
		{"Get out of bounds index", []int{10, 20, 30}, 5, true, 0}, // expectedVal can be zero value
		{"Get from nil list", nil, 0, true, 0},
		{"Get with negative index", []int{10, 20, 30}, -1, true, 0}, // Add this case for completeness
	}

	assert := assert.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			list := NewSinglyList(tc.array)
			node := list.SinglyGet(tc.index)

			if tc.expectNil {
				assert.Nil(node, "Expected a nil node but got one")
			} else {
				// We can add more assertions for the non-nil case
				assert.NotNil(node, "Expected a node but got nil")
				if node != nil { // Check for nil to prevent panic on the next line
					assert.Equal(tc.expectedVal, node.Value, "Node value does not match expected value")
				}
			}
		})
	}
}
