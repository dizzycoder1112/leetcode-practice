package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		// 一直往左走，把節點推入 stack
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// 彈出節點，處理它
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		// 轉向右子樹
		current = current.Right
	}

	return result
}
