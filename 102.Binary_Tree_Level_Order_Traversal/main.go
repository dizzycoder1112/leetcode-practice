package main

import "fmt"

func main() {

	node1 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 9}},
			Right: &TreeNode{Val: 5,
				Left:  &TreeNode{Val: 10},
				Right: &TreeNode{Val: 11}}},
		Right: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 6,
				Left:  &TreeNode{Val: 12},
				Right: &TreeNode{Val: 13}},
			Right: &TreeNode{Val: 7,
				Left:  &TreeNode{Val: 14},
				Right: &TreeNode{Val: 15}}}}

	fmt.Println(levelOrder(node1))
}

/**
 * Definition for a binary tree node.

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	// 边界情况:空树
	if root == nil {
		return result
	}

	// 队列初始化:放入根节点
	queue := []*TreeNode{root}

	// 只要队列不为空,就继续处理
	for len(queue) > 0 {
		// 关键:记录当前层有多少个节点
		levelSize := len(queue)
		currentLevelValues := make([]int, 0)

		// 处理当前层的所有节点
		for range levelSize {
			// 取出队首节点
			node := queue[0]
			queue = queue[1:]

			// 收集当前节点的值
			currentLevelValues = append(currentLevelValues, node.Val)

			// 把下一层的节点加入队列
			// 注意:是通过 node.Left 和 node.Right 指针访问,不是数组索引!
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 当前层处理完毕,加入结果
		result = append(result, currentLevelValues)
	}

	return result
}
