package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewLinkedList tests the creation of a new list from a slice.
func TestNewLinkedList(t *testing.T) {
	assert := assert.New(t)

	t.Run("Create from empty slice", func(t *testing.T) {
		head := NewLinkedList([]int{})
		assert.Nil(head, "NewLinkedList with empty slice should return nil")
	})

	t.Run("Create from non-empty slice", func(t *testing.T) {
		head := NewLinkedList([]int{1, 2, 3})
		assert.NotNil(head)
		assert.Equal([]int{1, 2, 3}, ToSlice(head), "The list should match the initial slice")
	})
}

// TestLenAndToSlice covers two fundamental helper functions.
func TestLenAndToSlice(t *testing.T) {
	assert := assert.New(t)

	t.Run("Empty list", func(t *testing.T) {
		var head *ListNode
		assert.Equal(0, Len(head), "Length of a nil list should be 0")
		assert.Equal([]int{}, ToSlice(head), "ToSlice on a nil list should be an empty slice")
	})

	t.Run("Single node list", func(t *testing.T) {
		head := &ListNode{Value: 10}
		assert.Equal(1, Len(head))
		assert.Equal([]int{10}, ToSlice(head))
	})

	t.Run("Multi-node list", func(t *testing.T) {
		head := NewLinkedList([]int{10, 20, 30})
		assert.Equal(3, Len(head))
		assert.Equal([]int{10, 20, 30}, ToSlice(head))
	})
}

// TestAppend tests adding a node to the end of the list.
func TestAppend(t *testing.T) {
	assert := assert.New(t)

	t.Run("Append to nil list", func(t *testing.T) {
		head := Append(nil, 10)
		assert.Equal([]int{10}, ToSlice(head))
	})

	t.Run("Append to non-empty list", func(t *testing.T) {
		head := NewLinkedList([]int{10, 20})
		head = Append(head, 30)
		assert.Equal([]int{10, 20, 30}, ToSlice(head))
	})
}

// TestPrepend tests adding a node to the beginning of the list.
func TestPrepend(t *testing.T) {
	assert := assert.New(t)

	t.Run("Prepend to nil list", func(t *testing.T) {
		head := Prepend(nil, 10)
		assert.Equal([]int{10}, ToSlice(head))
	})

	t.Run("Prepend to non-empty list", func(t *testing.T) {
		head := NewLinkedList([]int{20, 30})
		head = Prepend(head, 10)
		assert.Equal([]int{10, 20, 30}, ToSlice(head))
	})
}

// TestInsert tests inserting a node at a specific index.
func TestInsert(t *testing.T) {
	testCases := []struct {
		name        string
		array       []int
		index       int
		value       int
		expectErr   bool
		expectedVal []int
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
			head, err := Insert(NewLinkedList(tc.array), tc.index, tc.value)

			if tc.expectErr {
				assert.Error(err)
			} else {
				// We can add more assertions for the non-nil case
				assert.NoError(err)
				if head != nil { // Check for nil to prevent panic on the next line
					assert.Equal(tc.expectedVal, ToSlice(head))
				}
			}
		})
	}
}

// TestDelete tests deleting the first occurrence of a value.
func TestDelete(t *testing.T) {
	testCases := []struct {
		name        string
		array       []int
		value       int
		expectNil   bool
		expectedVal []int
	}{
		// Cases where we expect a node
		{"Delete head", []int{10, 20, 30}, 10, false, []int{20, 30}},
		{"Delete middle node", []int{10, 20, 30}, 20, false, []int{10, 30}},
		{"Delete tail node", []int{10, 20, 30}, 30, false, []int{10, 20}},
		{"Delete value not present", []int{10, 20, 30}, 40, false, []int{10, 20, 30}},
		{"Delete first of duplicates", []int{10, 20, 10, 30}, 10, false, []int{20, 10, 30}},

		// Cases where we expect nil
		{"Delete from nil list", []int{}, 10, true, nil},
		{"Delete from single node list", []int{10}, 10, true, nil},
	}

	assert := assert.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := Delete(NewLinkedList(tc.array), tc.value)

			if tc.expectNil {
				assert.Nil(head)
			} else {
				// We can add more assertions for the non-nil case
				assert.NotNil(head)
				if head != nil { // Check for nil to prevent panic on the next line
					assert.Equal(tc.expectedVal, ToSlice(head))
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
		expectErr   bool
		expectedVal []int
	}{
		// Cases where we expect a node
		{"Delete at index 0", []int{10, 20, 30}, 0, false, []int{20, 30}},
		{"Delete at middle index", []int{10, 20, 30}, 1, false, []int{10, 30}},
		{"Delete at last index", []int{10, 20, 30}, 2, false, []int{10, 20}},
		// Cases where we expect nil
		{"Index out of range (negative)", []int{10, 20}, -1, true, nil},
		{"Index out of range (equal to length)", []int{10, 20}, 2, true, nil},
	}

	assert := assert.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head, err := DeleteAt(NewLinkedList(tc.array), tc.index)

			if tc.expectErr {
				assert.Error(err)
			} else {
				// We can add more assertions for the non-nil case
				assert.NoError(err)
				if head != nil { // Check for nil to prevent panic on the next line
					assert.Equal(tc.expectedVal, ToSlice(head))
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
			node := Find(NewLinkedList(tc.array), tc.value)

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
			node := Get(NewLinkedList(tc.array), tc.index)

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
