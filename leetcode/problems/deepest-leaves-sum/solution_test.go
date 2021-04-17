package main

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) build(in []int) *TreeNode {
	if len(in) == 0 {
		return root
	}

	var cur, next []**TreeNode
	cur = []**TreeNode{&root}

	for len(in) != 0 {
		next = make([]**TreeNode, 2*len(cur))
		for i := range cur {
			if in[i] == -1 {
				continue
			}
			if *cur[i] == nil {
				*cur[i] = new(TreeNode)
			}
			node := *cur[i]
			node.Val = in[i]
			next[2*i] = &node.Left
			next[2*i+1] = &node.Right
		}
		in = in[len(cur):]
		cur = next
	}

	return root
}

// Solution for https://leetcode.com/problems/deepest-leaves-sum/.
func TestDeepestLeavesSum(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected int
	}{
		{
			root: new(TreeNode).build([]int{
				1, 2, 3, 4, 5, -1, 6, 7, -1, -1, -1, -1, -1, -1, 8,
			}),
			expected: 15,
		},
		{
			root: new(TreeNode).build([]int{
				6, 7, 8, 2, 7, 1, 3, 9, -1, 1, 4, -1, -1, -1, 5,
			}),
			expected: 19,
		},
	}

	for i, test := range tests {
		obtained := deepestLeavesSum(test.root)
		if obtained != test.expected {
			t.Errorf("expected: %d, obtained: %d (case %d)",
				test.expected, obtained, i)
		}
	}
}

// log(n)
func deepestLeavesSum(root *TreeNode) int {
	var dive func(node *TreeNode, depth int) (int, int)
	dive = func(node *TreeNode, depth int) (int, int) {
		if node.Left == nil && node.Right == nil {
			return node.Val, depth
		}
		var (
			lVal, lDepth int
			rVal, rDepth int
		)
		if node.Left != nil {
			lVal, lDepth = dive(node.Left, depth+1)
		}
		if node.Right != nil {
			rVal, rDepth = dive(node.Right, depth+1)
		}
		if lDepth > rDepth {
			return lVal, lDepth
		}
		if rDepth > lDepth {
			return rVal, rDepth
		}
		return lVal + rVal, lDepth
	}
	sum, _ := dive(root, 1)
	return sum
}
