package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	current := root

	for current != nil {
		if current.Val > p.Val && current.Val > q.Val {
			current = current.Left
		} else if current.Val < p.Val && current.Val < q.Val {
			current = current.Right
		} else {
			return current
		}
	}

	return nil
}
