Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

 

Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []
 

Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000

使用廣度優先演算法(BFS)
4层满二叉树

             1
          /     \
         2       3
        / \     / \
       4   5   6   7
      /|\ /|\ /|\ /|\
     8 9 A B C D E F

  用数字表示:
             1
          /     \
         2       3
        / \     / \
       4   5   6   7
      /|\ /|\ /|\ /|\
     8 9 10 11 12 13 14 15

  内存中的结构

  node1 := &TreeNode{Val: 1,
      Left: &TreeNode{Val: 2,
          Left: &TreeNode{Val: 4,
              Left: &TreeNode{Val: 8},
              Right: &TreeNode{Val: 9}},
          Right: &TreeNode{Val: 5,
              Left: &TreeNode{Val: 10},
              Right: &TreeNode{Val: 11}}},
      Right: &TreeNode{Val: 3,
          Left: &TreeNode{Val: 6,
              Left: &TreeNode{Val: 12},
              Right: &TreeNode{Val: 13}},
          Right: &TreeNode{Val: 7,
              Left: &TreeNode{Val: 14},
              Right: &TreeNode{Val: 15}}}}

  ---
  逐循环分析

  初始状态:

  queue = [node1]
  result = []

  ---
  第1次循环 - 处理第1层

  进入循环:
  levelSize = len(queue) = 1     // 第1层有1个节点
  currentLevel = []

  内层循环 (i=0):
  node = queue[0] = node1        // 取出 node1
  queue = []                     // 队列变空

  currentLevel = [1]             // 收集值1

  // node1.Left 不为 nil
  queue = [node2]                // 加入 node2

  // node1.Right 不为 nil  
  queue = [node2, node3]         // 加入 node3

  本层处理完毕:
  result = [[1]]
  queue = [node2, node3]         // 第2层的所有节点

  ---
  第2次循环 - 处理第2层

  进入循环:
  levelSize = len(queue) = 2     // 第2层有2个节点
  currentLevel = []

  内层循环 (i=0) - 处理 node2:
  node = queue[0] = node2
  queue = [node3]

  currentLevel = [2]

  // node2.Left 不为 nil
  queue = [node3, node4]

  // node2.Right 不为 nil
  queue = [node3, node4, node5]

  内层循环 (i=1) - 处理 node3:
  node = queue[0] = node3
  queue = [node4, node5]

  currentLevel = [2, 3]

  // node3.Left 不为 nil
  queue = [node4, node5, node6]

  // node3.Right 不为 nil
  queue = [node4, node5, node6, node7]

  本层处理完毕:
  result = [[1], [2,3]]
  queue = [node4, node5, node6, node7]    // 第3层的所有节点

  ---
  第3次循环 - 处理第3层

  进入循环:
  levelSize = len(queue) = 4     // 第3层有4个节点
  currentLevel = []

  内层循环 (i=0) - 处理 node4:
  node = node4
  queue = [node5, node6, node7]
  currentLevel = [4]
  queue = [node5, node6, node7, node8, node9]    // 加入 node8, node9

  内层循环 (i=1) - 处理 node5:
  node = node5
  queue = [node6, node7, node8, node9]
  currentLevel = [4, 5]
  queue = [node6, node7, node8, node9, node10, node11]

  内层循环 (i=2) - 处理 node6:
  node = node6
  queue = [node7, node8, node9, node10, node11]
  currentLevel = [4, 5, 6]
  queue = [node7, node8, node9, node10, node11, node12, node13]

  内层循环 (i=3) - 处理 node7:
  node = node7
  queue = [node8, node9, node10, node11, node12, node13]
  currentLevel = [4, 5, 6, 7]
  queue = [node8, node9, node10, node11, node12, node13, node14, node15]

  本层处理完毕:
  result = [[1], [2,3], [4,5,6,7]]
  queue = [node8, node9, node10, node11, node12, node13, node14, node15]  // 第4层的8个节点

  ---
  第4次循环 - 处理第4层

  进入循环:
  levelSize = len(queue) = 8     // 第4层有8个节点
  currentLevel = []

  内层循环 (i=0 到 i=7):
  依次取出 node8, node9, ..., node15:
  // 每个节点都没有子节点(叶子节点)
  // 只收集值,不加入新节点到队列

  currentLevel = [8, 9, 10, 11, 12, 13, 14, 15]
  queue = []                     // 队列变空

  本层处理完毕:
  result = [[1], [2,3], [4,5,6,7], [8,9,10,11,12,13,14,15]]
  queue = []                     // 队列空了

  ---
  第5次检查:

  len(queue) == 0  // 不满足循环条件,退出

  最终返回:
  [[1], [2,3], [4,5,6,7], [8,9,10,11,12,13,14,15]]

  ---