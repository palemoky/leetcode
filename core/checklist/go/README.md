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

=== "DFS 递归解法"

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
    func maxDepth(root *TreeNode) int {
    	if root == nil {
    		return 0
    	}

    	left := maxDepth(root.Left)
    	right := maxDepth(root.Right)

    	return max(left, right) + 1
    }
    ```

=== "平衡树"

    平衡树就是左右子树高度差不超过 1，所以要基于后序遍历的树深度来求解

    ```go
    func abs(x int) int {
      if x < 0 {
    		return -x
    	}
    	return x
    }

    func isBalanced(root *utils.TreeNode) bool {
    	return checkHeight(root) != -1
    }

    // 返回树的高度，如果不平衡则返回 -1
    // 后序遍历：先递归左右子树，返回时处理当前节点
    func checkHeight(root *utils.TreeNode) int {
    	if root == nil {
    		return 0
    	}

    	// 后序遍历：先检查左子树
    	leftHeight := checkHeight(root.Left)
    	if leftHeight == -1 {
    		return -1 // 左子树不平衡，提前终止
    	}

    	// 后序遍历：再检查右子树
    	rightHeight := checkHeight(root.Right)
    	if rightHeight == -1 {
    		return -1 // 右子树不平衡，提前终止
    	}

    	// 返回时处理：检查当前节点是否平衡
    	if abs(leftHeight-rightHeight) > 1 {
    		return -1 // 当前节点不平衡
    	}

    	// 返回当前节点的高度（自底向上汇总信息）
    	return max(leftHeight, rightHeight) + 1
    }
    ```

=== "树的直径"

    ```go
    func diameterOfBinaryTree(root *utils.TreeNode) int {
    	maxDiameter := 0

    	// 使用闭包捕获 maxDiameter
    	var depth func(*utils.TreeNode) int
    	depth = func(node *utils.TreeNode) int {
    		if node == nil {
    			return 0
    		}

    		// 后序遍历：先递归计算左右子树的深度
    		leftDepth := depth(node.Left)
    		rightDepth := depth(node.Right)

    		// 当前节点的直径 = 左子树深度 + 右子树深度
    		maxDiameter = max(maxDiameter, leftDepth+rightDepth)

    		// 计算当前节点的深度
    		return 1 + max(leftDepth, rightDepth)
    	}

    	depth(root)

    	return maxDiameter
    }
    ```

### 中序遍历题目

=== "验证二叉搜索树"

    由于需要捕获外部变量，因此可以使用闭包避免二级指针

    ```go
    func isValidBST(root *TreeNode) bool {
        // prev 记录中序遍历中上一个访问的节点
        // 放在闭包外部，所有递归调用共享同一个变量
        var prev *TreeNode

        // 定义闭包函数，自动捕获外部的 prev 变量
        var inorder func(*TreeNode) bool
        inorder = func(node *TreeNode) bool {
            // 递归终止条件：空节点视为有效
            if node == nil {
                return true
            }

            // 中序遍历：左 -> 根 -> 右
            // 1. 先递归检查左子树
            if !inorder(node.Left) {
                return false // 左子树不合法，提前终止
            }

            // 2. 检查当前节点：BST 的中序遍历必须严格递增
            if prev != nil && node.Val <= prev.Val {
                return false // 不满足严格递增，不是 BST
            }
            prev = node // 更新 prev 为当前节点，供下一个节点比较

            // 3. 递归检查右子树
            return inorder(node.Right)
        }

        return inorder(root)
    }
    ```

### 前序遍历题目

当节点需要依赖左右子树的信息时，使用前序遍历，这样不仅代码简单，而且高效

=== "翻转二叉树"

    ```go
    func invertTree(root *TreeNode) *TreeNode {
      if root == nil {
          return nil
      }

      // 交换当前节点的左右子树
      root.Left, root.Right = root.Right, root.Left

      // 递归翻转子树
      invertTree(root.Left)
      invertTree(root.Right)

      return root
    }
    ```

=== "对称树"

    <div class="grid" markdown>
    <div markdown>
    对称树的定义：

    - 左子树的左节点 == 右子树的右节点

    - 左子树的右节点 == 右子树的左节点
    </div>
    <div markdown>![对称树示意图](symmetric_tree.webp)</div>
    </div>

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

    判断给定的树中是否有和为 targetSum 的路径存在。

    ```go
    func hasPathSum(root *utils.TreeNode, targetSum int) bool {
    	if root == nil {
    		return false
    	}

    	// 到达叶子节点
    	if root.Left == nil && root.Right == nil {
    		return root.Val == targetSum
    	}

    	remainingSum := targetSum - root.Val
    	// 只需要左、右子树其中一个满足条件即可
      // 短路求值提前终止
    	return hasPathSum(root.Left, remainingSum) || hasPathSum(root.Right, remainingSum)
    }
    ```

=== "复制树"

### 层序遍历题目

## 数组

=== "无重复字符的最长子串"

    ```go
    func lengthOfLongestSubstring(s string) int {
        // 哈希表记录出现的元素及索引
        // 哈希表出现的元素，left 移动到哈希表中索引 +1 位置，否则右指针右移扩大窗口

        seen := map[byte]int{}
        left, right, maxWin := 0, 0, 0
        for right < len(s) {
            if prevIndex, ok := seen[s[right]]; ok {
                left = prevIndex + 1
            }

            seen[s[right]] = right
            right++

            maxWin = max(maxWin, right-left)
        }

        return maxWin
    }
    ```

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
