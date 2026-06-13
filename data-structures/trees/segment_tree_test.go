package trees

import (
	"testing"
)

func SumOperation(first, second int) int {
	return first + second
}

func TestSegmentTree_Query(t *testing.T) {
	tests := []struct {
		name    string
		nodes   []int
		updates []struct {
			index int
			value int
		}
		left     int
		right    int
		expected int
	}{
		{
			name:     "full range",
			nodes:    []int{1, 2, 3, 4, 5, 6},
			left:     0,
			right:    4,
			expected: 15,
		},
		{
			name:     "single element",
			nodes:    []int{1, 2, 3, 4, 5},
			left:     2,
			right:    2,
			expected: 3,
		},
		{
			name:     "middle range",
			nodes:    []int{1, 2, 3, 4, 5},
			left:     1,
			right:    3,
			expected: 9,
		},
		{
			name:     "after update",
			nodes:    []int{1, 2, 3, 4, 5},
			updates:  []struct{ index, value int }{{2, 10}},
			left:     0,
			right:    4,
			expected: 22,
		},
		{
			name:  "multiple updates",
			nodes: []int{1, 2, 3, 4, 5},
			updates: []struct {
				index int
				value int
			}{
				{0, 10},
				{4, 20},
			},
			left:     0,
			right:    4,
			expected: 39,
		},
		{
			name:     "query updated element",
			nodes:    []int{1, 2, 3, 4, 5},
			updates:  []struct{ index, value int }{{3, 100}},
			left:     3,
			right:    3,
			expected: 100,
		},
		{
			name:     "update does not affect range",
			nodes:    []int{1, 2, 3, 4, 5},
			updates:  []struct{ index, value int }{{4, 100}},
			left:     0,
			right:    2,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewSegmentTree(tt.nodes, SumOperation)

			for _, upd := range tt.updates {
				tree.Update(upd.index, upd.value)
			}

			result := tree.Query(tt.left, tt.right)

			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
