package main

import "fmt"

func main() {
	// 测试方案 2（优化版）
	fmt.Println("=== 测试 MyQueue（方案2：优化版）===")
	q := Constructor()

	fmt.Println("Push(1)")
	q.Push(1)
	q.printState()

	fmt.Println("\nPush(2)")
	q.Push(2)
	q.printState()

	fmt.Println("\nPeek():", q.Peek())
	q.printState()

	fmt.Println("\nPop():", q.Pop())
	q.printState()

	fmt.Println("\nPush(3)")
	q.Push(3)
	q.printState()

	fmt.Println("\nPush(4)")
	q.Push(4)
	q.printState()

	fmt.Println("\nPop():", q.Pop())
	q.printState()

	fmt.Println("\nEmpty():", q.Empty())
}

// ========== 方案 1：Push 时倒转（简单但慢）==========
// 每次 Push 都重新排列，保证栈顶永远是队列的 front
// Push: O(n), Pop: O(1)

type MyQueueSimple struct {
	stack1 []int // 主栈
	stack2 []int // 辅助栈（用于倒转）
}

func ConstructorSimple() MyQueueSimple {
	return MyQueueSimple{
		stack1: []int{},
		stack2: []int{},
	}
}

func (q *MyQueueSimple) Push(x int) {
	// 核心思路：每次 Push 都要把元素放到栈底
	// 通过倒转实现

	// 1. 把 stack1 的所有元素倒到 stack2
	//    这样 stack2 的顺序就是反的
	for len(q.stack1) > 0 {
		top := q.stack1[len(q.stack1)-1]
		q.stack1 = q.stack1[:len(q.stack1)-1]
		q.stack2 = append(q.stack2, top)
	}

	// 2. 把新元素 push 到 stack1（现在是空的）
	//    新元素会在最底部
	q.stack1 = append(q.stack1, x)

	// 3. 把 stack2 的元素倒回 stack1
	//    这样新元素就在栈底，旧元素在栈顶
	for len(q.stack2) > 0 {
		top := q.stack2[len(q.stack2)-1]
		q.stack2 = q.stack2[:len(q.stack2)-1]
		q.stack1 = append(q.stack1, top)
	}
}

func (q *MyQueueSimple) Pop() int {
	// 因为 Push 已经保证了顺序，直接 pop 栈顶
	top := q.stack1[len(q.stack1)-1]
	q.stack1 = q.stack1[:len(q.stack1)-1]
	return top
}

func (q *MyQueueSimple) Peek() int {
	return q.stack1[len(q.stack1)-1]
}

func (q *MyQueueSimple) Empty() bool {
	return len(q.stack1) == 0
}

// ========== 方案 2：Pop 时倒转（优化，Amortized O(1)）==========
// 延迟倒转，只在需要 Pop/Peek 且 outStack 为空时才倒转
// Push: O(1), Pop/Peek: 摊销 O(1)

type MyQueue struct {
	inStack  []int // 输入栈：只负责 Push
	outStack []int // 输出栈：只负责 Pop/Peek
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

// Push: 直接放到 inStack，O(1)
func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

// Pop: 从 outStack 取，如果空则先倒转
func (q *MyQueue) Pop() int {
	q.move() // 确保 outStack 有元素

	// Stack 的 pop 操作
	top := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return top
}

// Peek: 查看 outStack 的栈顶
func (q *MyQueue) Peek() int {
	q.move() // 确保 outStack 有元素
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

// move: 关键方法！只在 outStack 为空时才倒转
// 这是 Amortized O(1) 的核心
func (q *MyQueue) move() {
	// 只有 outStack 为空时才需要倒转
	if len(q.outStack) == 0 {
		// 把 inStack 的所有元素倒到 outStack
		// 倒转后，最早进入的元素会在 outStack 的栈顶
		for len(q.inStack) > 0 {
			top := q.inStack[len(q.inStack)-1]
			q.inStack = q.inStack[:len(q.inStack)-1]
			q.outStack = append(q.outStack, top)
		}
	}
}

// 辅助方法：打印状态（用于理解）
func (q *MyQueue) printState() {
	fmt.Printf("  inStack:  %v (top→)", q.inStack)
	if len(q.inStack) > 0 {
		fmt.Printf(" [top=%d]", q.inStack[len(q.inStack)-1])
	}
	fmt.Println()

	fmt.Printf("  outStack: %v (top→)", q.outStack)
	if len(q.outStack) > 0 {
		fmt.Printf(" [top=%d]", q.outStack[len(q.outStack)-1])
	}
	fmt.Println()
}

// ========== Stack 性质详解 ==========
/*
核心理解：

1. Stack 的 LIFO 特性
   Push: [1] → [1,2] → [1,2,3]
   Pop:  3 ← 2 ← 1 (倒序)

2. 为什么需要两个栈？
   单个栈：[1,2,3]
          栈顶是 3，但队列需要先取出 1

   两个栈倒转：
   inStack:  [1,2,3]  (栈顶是 3)
              ↓ 全部倒到 outStack
   outStack: [3,2,1]  (栈顶是 1) ✓

3. 倒转的本质
   Stack1 → pop 顺序：3, 2, 1
   Stack2 → push 这个顺序：stack2 = [3,2,1]
   Stack2 → pop 顺序：1, 2, 3 (恢复了原始顺序！)

4. 为什么方案 2 更快？
   方案 1：每次 Push 都倒转 2 次，O(n)
   方案 2：只在 outStack 空时倒转 1 次
          - 每个元素最多移动 2 次：进 inStack → 倒到 outStack
          - n 次操作总共 O(n)，摊销 O(1)

5. Amortized Analysis（摊销分析）
   例子：Push 1,2,3,4,5 然后 Pop 5 次

   Push(1): inStack=[1]         → O(1)
   Push(2): inStack=[1,2]       → O(1)
   Push(3): inStack=[1,2,3]     → O(1)
   Push(4): inStack=[1,2,4]     → O(1)
   Push(5): inStack=[1,2,3,4,5] → O(1)

   Pop():  move() 倒转 5 个元素 → O(5)
           outStack=[5,4,3,2,1]
           pop 1 → O(1)
   Pop():  outStack=[5,4,3,2]
           pop 2 → O(1) (不需要倒转！)
   Pop():  pop 3 → O(1)
   Pop():  pop 4 → O(1)
   Pop():  pop 5 → O(1)

   总共：10 次操作，O(10) = O(n)
   平均：每次 O(1)

6. Stack 的操作限制
   只能从栈顶操作：
   ✓ Push(top)
   ✓ Pop(top)
   ✓ Peek(top)
   ✗ 不能访问中间元素
   ✗ 不能从底部操作

   这个限制导致了需要"倒转"的技巧

7. 延迟求值（Lazy Evaluation）
   方案 2 的优化思想：
   - 不立即倒转，等到真正需要时再倒
   - 批量倒转比每次倒转更高效
   - 类似于：不提前做所有功课，deadline 前集中完成
*/

// ========== 图解示例 ==========
/*
方案 2 的完整执行过程：

初始状态：
inStack:  []
outStack: []

Push(1):
inStack:  [1] ← top
outStack: []
说明：直接 push 到 inStack

Push(2):
inStack:  [1, 2] ← top
outStack: []
说明：继续 push 到 inStack

Peek():
1. move() 被调用
2. outStack 为空，开始倒转
3. inStack pop: 2 → outStack push: 2
4. inStack pop: 1 → outStack push: 1
结果：
inStack:  []
outStack: [2, 1] ← top (队列的 front)
返回：1 ✓

Pop():
1. move() 被调用
2. outStack 不为空，不倒转
3. outStack pop: 1
结果：
inStack:  []
outStack: [2] ← top
返回：1

Push(3):
inStack:  [3] ← top
outStack: [2] ← top
说明：新元素进 inStack，outStack 保持不变

Push(4):
inStack:  [3, 4] ← top
outStack: [2] ← top

Pop():
1. move() 被调用
2. outStack 不为空（还有 2），不倒转
3. outStack pop: 2
结果：
inStack:  [3, 4] ← top
outStack: []
返回：2

Pop():
1. move() 被调用
2. outStack 为空，开始倒转
3. inStack pop: 4 → outStack push: 4
4. inStack pop: 3 → outStack push: 3
结果：
inStack:  []
outStack: [4, 3] ← top
返回：3

关键观察：
- 每个元素最多被移动 1 次（从 inStack 到 outStack）
- 移动是批量的，不是每次操作都移动
- outStack 有元素时，Pop 是 O(1)
*/
