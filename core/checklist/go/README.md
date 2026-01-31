# Go 刷题清单

## 链表

=== "反转链表"

    ```go
    func reverseList(head *ListNode) *ListNode {
        var prev *ListNode
        // 暂存、反转、移动
        for head != nil {
            next := head.next
            head.Next = prev
            prev = head
            head = next
        }

        return prev
    }
    ```

=== "局部反转链表"

    头插法，将每个要反转的节点连接到 `prev` 后

    ```go
    func reverseBetween(head *ListNode, left int, right int) *ListNode {
    	dummy := &ListNode{Next: head}

    	prev := dummy
    	for range left - 1 {
    		prev = prev.Next
    	}

    	cur := prev.Next
    	for range right - left {
    		next := cur.Next
    		cur.Next = next.Next
    		next.Next = prev.Next
    		prev.Next = next
    	}

    	return dummy.Next
    }
    ```

=== "K个一组反转链表"

    分组+局部反转

    ```go
    func reverseKGroup(head *ListNode, k int) *ListNode {
    	dummy := &ListNode{Next: head}
    	pre := dummy

    	for {
    		tail := pre
    		for range k {
    			tail = tail.Next
    			if tail == nil {
    				return dummy.Next
    			}
    		}

    		nextGroup := tail.Next

    		cur := pre.Next
    		for range k - 1 {
    			next := cur.Next
    			cur.Next = next.Next
    			next.Next = pre.Next
    			pre.Next = next
    		}

    		cur.Next = nextGroup
    		pre = cur
    	}
    }
    ```

=== "判断链表是否有环"

    快慢指针解法：快指针走两步，慢指针走一步，相遇则有环

    ```go
    func hasCycle(head *ListNode) bool {
        slow, fast := head, head
        for fast != nil && fast.Next != nil {
            slow, fast = slow.Next, fast.Next.Next
            if slow == fast {
                return true
            }
        }

        return false
    }
    ```

=== "找到环形链表的入口"

    ```go
    func detectCycle(head *ListNode) *ListNode {
        seen := map[*ListNode]struct{}{}
        for head != nil {
            if _, ok := seen[head]; ok {
                return head
            }
            seen[head] = struct{}{}
            head = head.Next
        }

        return nil
    }
    ```

## 二叉树

### 遍历方式

=== "层序遍历"

    队列+双层循环（外循环控制深度，内循环控制宽度）

    ```go
    func levelOrder(root *TreeNode) [][]int {
      ans := [][]int{}
      if root == nil {
        return ans
      }

      queue := []*TreeNode{root}
      for len(queue) > 0 {
        level := make([]int, 0, len(queue))
        for range len(queue) {
          node := queue[0]
          queue = queue[1:]

          level = append(level, node.Val)
          if node.Left != nil {
            queue = append(queue, node.Left)
          }
          if node.Right != nil {
            queue = append(queue, node.Right)
          }
        }
        ans = append(ans, level)
      }

      return ans
    }
    ```

=== "前序遍历（最简单）"

    栈：先压右再压左

    ```go
    func preorderTraversal(root *TreeNode) []int {
      ans := []int}{}
      if root == nil {
        return ans
      }

    	stack := []*TreeNode{root}
      for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        ans = append(ans, node.Val)
        // 因为栈是先进后出，所以先压右节点
        if node.Right != nil {
          stack = append(stack, node.Right)
        }
        if node.Lfet != nil {
          stack = append(stack, node.Left)
        }
      }

    	return ans
    }
    ```

=== "中序遍历（仅限于二叉树）"

    栈：一路向左，先处理完左子树再处理右子树

    ```go
    func inorderTraversal(root *TreeNode) []int {
      ans := []int{}
      stack := []*TreeNode{}

      curr := root
      for curr != nil || len(stack) > 0 {
        // 把节点一路向左压入栈
        for curr != nil {
          stack = append(stack, curr)
          curr = curr.Left
        }

        // 开始倒序处理栈中的节点（即从下往上遍历树）
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        ans = append(ans, curr.Val)

        // 左子树处理完处理右子树
        curr = curr.Right
      }

      return ans
    }
    ```

=== "后续遍历（最复杂）"

    后序遍历和中序遍历有点像，主要不同点在于，中序遍历对每个节点只访问一次，而后序遍历则可能访问两次，且通过 `prev` 来记录是否已经访问。后续遍历的`prev` 类似于回溯中的 `used[]`

    !!! Note

        以 `[1,2,3,4,5]`为例，

        ```
            1
           / \
          2   3
         / \
        4   5
        ```

        第1次遇到节点2（检查阶段）：

        ```go
        // 一路向左：1 → 2 → 4
        stack = [1, 2, 4]

        // 访问完4后，回到节点2
        curr = stack[1] = 2  // ← 第1次遇到节点2

        // 判断
        if curr.Right == nil || curr.Right == prev {
            // 2.Right = 5, prev = 4
            // 5 != 4 不满足条件
        } else {
            curr = curr.Right  // ← 转向右子树5，暂时不访问2
        }
        ```

        第2次遇到节点2（访问阶段）：

        ```go
        // 访问完5后，再次回到节点2
        stack = [1, 2]
        curr = stack[1] = 2  // ← 第2次遇到节点2

        // 判断
        if curr.Right == nil || curr.Right == prev {
            // 2.Right = 5, prev = 5
            // 5 == 5 ✅ 满足条件（右子树已访问）

            stack = [1]
            nums = append(nums, 2)  // ← 现在才真正访问节点2
            prev = 2
            curr = nil
        }
        ```

        每个节点可能被访问两次：

        - 检查右子树

        - 右子树访问完后，才真正访问当前节点

    ```go
    func postorderTraversal(root *TreeNode) []int {
      ans := []int{}
      stack := []*TreeNode{}
      var prev *TreeNode

      curr := root
      for curr != nil || len(stack) > 0 {
        for curr != nil {
          stack = append(stack, curr)
          curr = curr.Left
        }

        // 查看栈顶不弹出
        curr = stack[Len(stack)-1]

        if curr.Right == nil || curr.Right == prev {
          stack = stack[:len(stack)-1]
          ans = append(ans, curr.Val)
          prev = curr
          curr = nil
        } else {
          curr = curr.Right
        }
      }

      return ans
    }
    ```

=== "递归解法"

    ```go
    func traversal(root *TreeNode) []int {
        if root == nil {
            return []int{}
        }

        // Preorder: 根 -> 左 -> 右
        // nums := []int{root.Val}
        // nums = append(nums, traversal(root.Left)...)
        // nums = append(nums, traversal(root.Right)...)

        // Inorder: 左 -> 根 -> 右
        nums := traversal(root.Left)
        nums = append(nums, root.Val)
        nums = append(nums, traversal(root.Right)...)

        // Postorder: 左 -> 右 -> 根
        // nums := traversal(root.Left)
        // nums = append(nums, traversal(root.Right)...)
        // nums = append(nums, root.Val)

        return nums
    }
    ```

### 后序遍历题目

当节点需要依赖左右子树的信息时，使用后序遍历

=== "最大深度"

    后序遍历的递归解法：**先钻到最底下，在返回时再做处理**

    ```go
    func maxDepthDFS(root *TreeNode) int {
    	if root == nil {
    		return 0
    	}

    	left := maxDepthDFS(root.Left)
    	right := maxDepthDFS(root.Right)

    	return max(left, right) + 1
    }
    ```

=== "树的直径"

=== "平衡树"

### 前序遍历题目

当节点需要依赖左右子树的信息时，使用前序遍历，这样不仅代码简单，而且高效

=== "对称树"

    对称树的定义：

    - 左子树的左节点 == 右子树的右节点

    - 左子树的右节点 == 右子树的左节点

    镜像递归
    ```go
    func isSymmetricMirrorRecursive(root *utils.TreeNode) bool {
    	if root == nil {
    		return true
    	}
    	return isMirror(root.Left, root.Right)
    }

    func isMirror(left, right *utils.TreeNode) bool {
    	// 递归终止条件
      // 检查节点存在的对称性
    	if left == nil && right == nil {
    		return true
    	}
    	if left == nil || right == nil {
    		return false
    	}

      // 递归处理逻辑
    	// 检查节点值的对称性
    	if left.Val != right.Val {
    		return false
    	}

    	// 递归处理：交叉比较子树（镜像对称）
    	return isMirror(left.Left, right.Right) &&
    		isMirror(left.Right, right.Left)
    }
    ```

=== "路径和"

=== "复制树"

### 层序遍历题目

## 数组

## 字符串

## 栈

## 队列

## 二分查找

## 递归

编写递归步骤：

1. 明确输入输出
2. 明确递归终止条件（什么时候触底反弹？）
3. 明确递归处理逻辑（每层要做什么事？）
4. 明确递归过程（下楼做还是返回做？）

## 回溯

## 动态规划

## 图论
