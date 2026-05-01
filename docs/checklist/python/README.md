# Python 刷题清单

## 哈希表

=== "#1 两数之和"

    ```python
    def twoSum(nums: list[int], target: int) -> list[int]:
        # 哈希表记录出现的元素及索引
        # 遍历 nums，如果 target - num 在哈希表中，返回索引
        # 否则将 num 加入哈希表
        seen = {}
        for i, num in enumerate(nums):
            if target - num in seen:
                return [i, seen[target - num]]
            seen[num] = i
        return []
    ```

=== "#49 字母异位词分组"

    <div class="grid cards" markdown>
    - <figure>
        ![排序解法](../imgs/group_anagrams/sort.webp)
        <figcaption>排序解法</figcaption>
    </figure>
    - <figure>
        ![计数解法](../imgs/group_anagrams/count.webp)
        <figcaption>计数解法</figcaption>
    </figure>
    </div>

    ```python
    # 排序解法
    # Time: O(m*nlogn), Space: O(m*n)
    def groupAnagrams(strs: list[str]) -> list[list[str]]:
        groups = {}
        for s in strs:
            key = ''.join(sorted(s))
            groups.setdefault(key, []).append(s)
        return list(groups.values())
    ```

    ```python
    # 计数解法
    # Time: O(m*n), Space: O(m*n)
    def groupAnagrams(strs: list[str]) -> list[list[str]]:
        groups = {}
        for s in strs:
            count = [0] * 26
            for ch in s:
                count[ord(ch) - ord('a')] += 1
            key = tuple(count)
            groups.setdefault(key, []).append(s)
        return list(groups.values())
    ```

## 双指针

=== "#3 无重复字符的最长子串"

    ```python
    def lengthOfLongestSubstring(s: str) -> int:
        # 哈希表记录出现的元素及索引
        # 哈希表出现的元素，left 移动到哈希表中索引 +1 位置，否则右指针右移扩大窗口
        seen = {}
        left = right = max_win = 0
        while right < len(s):
            if s[right] in seen:
                left = seen[s[right]] + 1
            seen[s[right]] = right
            right += 1
            max_win = max(max_win, right - left)
        return max_win
    ```

=== "#15 三数之和"

    ```python
    def threeSum(nums: list[int]) -> list[list[int]]:
    	if len(nums) < 3:
    		return []

    	nums.sort()

    	ans = []
    	# 先固定第一个数
    	for i in range(len(nums) - 2):
    		# 提前剪枝：如果第一个数已经大于0，后面都是正数，不可能和为0
    		if nums[i] > 0:
    			break

    		if i > 0 and nums[i] == nums[i-1]:
    			continue  # 跳过重复的第一个数

    		# 用双指针让剩余两数之和与第一个数的和为0
    		left, right = i + 1, len(nums) - 1
    		while left < right:
    			s = nums[i] + nums[left] + nums[right]

    			# 注意此时已经排序
    			if s < 0:  # 和太小，需要右移靠近较大数
    				left += 1
    			elif s > 0:  # 和太大，需要左移靠近较小数
    				right -= 1
    			else:  # 和为 0
    				ans.append([nums[i], nums[left], nums[right]])

    				# 跳过重复的第二个数
    				while left < right and nums[left] == nums[left+1]:
    					left += 1

    				# 跳过重复的第三个数
    				while left < right and nums[right] == nums[right-1]:
    					right -= 1

    				left += 1
    				right -= 1

    	return ans
    ```

=== "#5 最长回文子串"

    由于本题查找最长回文子串，我们并不清楚回文边界，因此 **无法使用对撞指针**，只能用回文中心对称的特点，从中心向两侧扩展求解。

    ```python
    # Time: O(n²), Space: O(1)
    def longestPalindrome(s: str) -> str:
    	start = max_len = 0

    	# 从中心向两侧扩展，返回回文长度
    	def expand(l: int, r: int) -> None:
    		nonlocal start, max_len
    		while l >= 0 and r < len(s) and s[l] == s[r]:
    			l -= 1
    			r += 1
    		# 此时 [l+1, r-1] 是回文
    		if r - l - 1 > max_len:
    			start = l + 1
    			max_len = r - l - 1

    	for i in range(len(s)):
    		# 无法预知以某个位置为中心的最长回文是奇数还是偶数长度，因此需要同时遍历两种情况，取最长的那个
    		expand(i, i)      # 奇数长度回文（如 "aba"）
    		expand(i, i + 1)  # 偶数长度回文（如 "abba"）

    	return s[start:start + max_len]
    ```

=== "#31 下一个排列"

    题目解释：将数组视为一个数字，在所有排列中找到**恰好比当前排列大的下一个排列**；若已是最大排列，则回绕到最小排列。

    - `nums = [1,2,3]`：对应数字 123，所有排列按大小排序为 123 → **132** → 213 → …，下一个排列为 `[1,3,2]`
    - `nums = [3,2,1]`：对应数字 321，已是最大排列，回绕到最小排列 `[1,2,3]`

    解题思路：

    1. 从右往左，找到第一个满足 `nums[k] < nums[k+1]` 的位置 `k`（即第一个"下降点"）。若不存在，说明已是最大排列，直接翻转整个数组即可。（可以通过折线图可视化查找过程）
    2. 再从右往左，找到第一个满足 `nums[l] > nums[k]` 的位置 `l`。
    3. 交换 `nums[k]` 与 `nums[l]`。
    4. 翻转 `nums[k+1:]`，使其从降序变为升序，得到恰好大一点的排列。

    以 `nums = [1,2,7,4,3,1]` 为例：

    - 找下降点：从右往左扫描，`nums[1]=2 < nums[2]=7`，故 `k=1`
    - 找交换点：从右往左找第一个大于 2 的数，`nums[4]=3`，故 `l=4`
    - 交换：`nums = [1,3,7,4,2,1]`
    - 翻转 `nums[2:]`：`nums = [1,3,1,2,4,7]` ✅

    ```python
    # Time: O(n), Space: O(1)
    def nextPermutation(nums: list[int]) -> None:
    	# 从右往左找到第一个下降点 i
    	i = len(nums) - 2
    	while i >= 0 and nums[i] >= nums[i+1]:
    		i -= 1

    	if i >= 0:
    		# 从右往左找到第一个大于 nums[i] 的数 j
    		j = len(nums) - 1
    		while nums[j] <= nums[i]:
    			j -= 1
    		# 交换 i 和 j
    		nums[i], nums[j] = nums[j], nums[i]

    	# 反转 i 之后的部分
    	left, right = i + 1, len(nums) - 1
    	while left < right:
    		nums[left], nums[right] = nums[right], nums[left]
    		left += 1
    		right -= 1
    ```

=== "#88 合并两个有序数组"

    双指针倒序填充

    ```python
    # Time: O(m+n), Space: O(1)
    def merge(nums1: list[int], m: int, nums2: list[int], n: int) -> None:
    	p1, p2, tail = m - 1, n - 1, m + n - 1

    	# 只需检查 p2 >= 0，因为 nums2 处理完后，nums1 剩余元素已在正确位置
    	while p2 >= 0:
    		if p1 >= 0 and nums1[p1] > nums2[p2]:
    			nums1[tail] = nums1[p1]
    			p1 -= 1
    		else:
    			nums1[tail] = nums2[p2]
    			p2 -= 1
    		tail -= 1
    ```

=== "合并区间"

=== "#165 比较版本号"

    ```python
    # Time: O(n+m), Space: O(1)
    def compareVersion(version1: str, version2: str) -> int:
    	n, m = len(version1), len(version2)
    	i = j = 0
    	while i < n or j < m:
    		x = 0
    		while i < n and version1[i] != '.':
    			x = x * 10 + int(version1[i])
    			i += 1
    		i += 1  # 跳过点号

    		y = 0
    		while j < m and version2[j] != '.':
    			y = y * 10 + int(version2[j])
    			j += 1
    		j += 1  # 跳过点号

    		if x > y:
    			return 1
    		if x < y:
    			return -1

    	return 0
    ```

## 链表

=== "#206 反转链表"

    ```python
    def reverseList(head: ListNode | None) -> ListNode | None:
        prev = None
        # 暂存、反转、移动
        while head:
            next_node = head.next
            head.next = prev
            prev = head
            head = next_node
        return prev
    ```

=== "#92 局部反转链表"

    头插法，将每个要反转的节点连接到 `prev` 后

    ```python
    def reverseBetween(head: ListNode | None, left: int, right: int) -> ListNode | None:
        dummy = ListNode(next=head)  # 可能从头结点开始反转

        prev = dummy
        for _ in range(left - 1):
            prev = prev.next

        curr = prev.next
        for _ in range(right - left):
            next_node = curr.next      # 暂存操作节点的移动路径
            curr.next = next_node.next # 摘下节点
            next_node.next = prev.next # 插到区间头部
            prev.next = next_node      # 移动操作的节点

        return dummy.next
    ```

    ![reverse_by_head_insert](../../data_structures/linked_list/reverse_by_head_insert.webp)

=== "#25 K个一组反转链表"

    分组+局部反转

    ```python
    def reverseKGroup(head: ListNode | None, k: int) -> ListNode | None:
        dummy = ListNode(next=head)
        prev = dummy

        while True:
            tail = prev
            for _ in range(k):
                tail = tail.next
                if not tail:
                    return dummy.next

            next_group = tail.next

            # 局部反转链表
            cur = prev.next
            for _ in range(k - 1):
                next_node = cur.next
                cur.next = next_node.next
                next_node.next = prev.next
                prev.next = next_node

            cur.next = next_group
            prev = cur
    ```

=== "#141 判断环形链表"

    快慢指针解法：快指针走两步，慢指针走一步，相遇则有环

    ```python
    def hasCycle(head: ListNode | None) -> bool:
        slow = fast = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow is fast:
                return True
        return False
    ```

=== "#142 找到环形链表的入口"

    ```python
    def detectCycle(head: ListNode | None) -> ListNode | None:
        seen = set()
        while head:
            if head in seen:
                return head
            seen.add(head)
            head = head.next
        return None
    ```

=== "#21 合并两个有序链表"

    ```python
    def mergeTwoLists(l1: ListNode | None, l2: ListNode | None) -> ListNode | None:
    	dummy = ListNode()
    	current = dummy

    	# 注意该遍历需同时操作3个链表
    	while l1 and l2:  # 注意是 and
    		if l1.val < l2.val:  # 将较小的值挂载在新链表上
    			current.next = l1
    			l1 = l1.next  # 移动原链表
    		else:
    			current.next = l2
    			l2 = l2.next
    		current = current.next  # 移动新链表

    	# 将链表剩余部分挂载，同时处理原链表为空
    	current.next = l1 if l1 else l2

    	# 注意返回的是 dummy.next
    	return dummy.next
    ```

=== "#160 相交链表"

=== "#143 重排链表"

=== "#148 排序链表"

    归并排序

=== "#146 LRU缓存"

    双向链表 + 哈希表

    ```python
    class Node:
        def __init__(self, key: int = 0, value: int = 0):
            self.key = key
            self.value = value
            self.prev: 'Node | None' = None
            self.next: 'Node | None' = None

    class LRUCache:
        def __init__(self, capacity: int):
            self.capacity = capacity
            self.cache: dict[int, Node] = {}  # 哈希表：key -> 链表节点
            self.head = Node()  # 虚拟头节点（最近使用）
            self.tail = Node()  # 虚拟尾节点（最久未使用）
            self.head.next = self.tail
            self.tail.prev = self.head

        def get(self, key: int) -> int:
            if key in self.cache:
                node = self.cache[key]
                # 将节点移到头部（标记为最近使用）
                self._move_to_head(node)
                return node.value
            return -1

        def put(self, key: int, value: int) -> None:
            if key in self.cache:
                # 更新已存在的节点
                node = self.cache[key]
                node.value = value
                self._move_to_head(node)
            else:
                # 创建新节点
                new_node = Node(key, value)
                self.cache[key] = new_node
                self._add_to_head(new_node)

                # 检查容量，必要时删除最久未使用的节点
                if len(self.cache) > self.capacity:
                    removed = self._remove_tail()
                    del self.cache[removed.key]

        # 辅助方法：将节点添加到头部
        def _add_to_head(self, node: Node) -> None:
            node.prev = self.head
            node.next = self.head.next
            self.head.next.prev = node
            self.head.next = node

        # 辅助方法：移除节点
        def _remove_node(self, node: Node) -> None:
            node.prev.next = node.next
            node.next.prev = node.prev

        # 辅助方法：将节点移到头部
        def _move_to_head(self, node: Node) -> None:
            self._remove_node(node)
            self._add_to_head(node)

        # 辅助方法：移除尾部节点
        def _remove_tail(self) -> Node:
            node = self.tail.prev
            self._remove_node(node)
            return node
    ```

## 二叉树

### 遍历方式

=== "#102 层序遍历"

    队列+双层循环（外循环控制深度，内循环控制宽度）

    ```python
    from collections import deque

    def levelOrder(root: TreeNode | None) -> list[list[int]]:
      ans = []
      if not root:
        return ans

      queue = deque([root])
      while queue:
        level = []
        for _ in range(len(queue)):
          node = queue.popleft()

          level.append(node.val)
          if node.left:
            queue.append(node.left)
          if node.right:
            queue.append(node.right)
        ans.append(level)

      return ans
    ```

=== "#144 前序遍历（最简单）"

    栈：先压右再压左

    ```python
    def preOrderTraversal(root: TreeNode | None) -> list[int]:
      ans = []
      if not root:
        return ans

      stack = [root]
      while stack:
        node = stack.pop()

        ans.append(node.val)
        # 因为栈是先进后出，所以先压右节点
        if node.right:
          stack.append(node.right)
        if node.left:
          stack.append(node.left)

      return ans
    ```

=== "#94 中序遍历（仅限于二叉树）"

    栈：一路向左，先处理完左子树再处理右子树

    ```python
    def inOrderTraversal(root: TreeNode | None) -> list[int]:
      ans = []
      stack = []

      curr = root
      while curr or stack:
        # 把节点一路向左压入栈
        while curr:
          stack.append(curr)
          curr = curr.left

        # 开始倒序处理栈中的节点（即从下往上遍历树）
        curr = stack.pop()

        ans.append(curr.val)

        # 左子树处理完处理右子树
        curr = curr.right

      return ans
    ```

=== "#145 后序遍历（最复杂）"

    后序遍历的迭代实现由中序遍历演化而来，主要区别在于后序遍历需在右子树遍历完成后才处理根节点。因此，引入 `prev` 指针用于记录历史状态，防止右子树被重复访问。这一机制类似于回溯算法中利用 `used[]` 数组来标记路径是否已被使用。

    !!! Note

        以 `[1,2,3,4,5]` 为例，

        ```
            1
           / \
          2   3
         / \
        4   5
        ```

        第 1 次遇到节点 2（检查阶段）：

        ```python
        # 一路向左：1 → 2 → 4
        stack = [1, 2, 4]

        # 访问完4后，回到节点2
        curr = stack[-1]  # = 2  ← 第1次遇到节点2

        # 判断
        if curr.right is None or curr.right is prev:
            # curr.right = 5, prev = 4
            # 5 is not 4，不满足条件
            pass
        else:
            curr = curr.right  # ← 转向右子树5，暂时不访问2
        ```

        第 2 次遇到节点 2（访问阶段）：

        ```python
        # 访问完5后，再次回到节点2
        stack = [1, 2]
        curr = stack[-1]  # = 2  ← 第2次遇到节点2

        # 判断
        if curr.right is None or curr.right is prev:
            # curr.right = 5, prev = 5
            # 5 is 5 ✅ 满足条件（右子树已访问）

            stack.pop()
            ans.append(curr.val)  # ← 现在才真正访问节点2
            prev = curr
            curr = None
        ```

        每个节点可能被访问两次：

        - 检查右子树

        - 右子树访问完后，才真正访问当前节点

    ```python
    def postOrderTraversal(root: TreeNode | None) -> list[int]:
      ans = []
      stack = []
      prev = None

      curr = root
      while curr or stack:
        while curr:
          stack.append(curr)
          curr = curr.left

        # 查看栈顶不弹出
        curr = stack[-1]

        if curr.right is None or curr.right is prev:
          stack.pop()
          ans.append(curr.val)
          prev = curr
          curr = None
        else:
          curr = curr.right

      return ans
    ```

=== "DFS 递归解法"

    ```python
    def traversal(root: TreeNode | None) -> list[int]:
        if not root:
            return []

        # Preorder: 根 -> 左 -> 右
        # return [root.val] + traversal(root.left) + traversal(root.right)

        # Inorder: 左 -> 根 -> 右
        return traversal(root.left) + [root.val] + traversal(root.right)

        # Postorder: 左 -> 右 -> 根
        # return traversal(root.left) + traversal(root.right) + [root.val]
    ```

### 后序遍历题目

当节点需要依赖左右子树的信息时，使用后序遍历

=== "#104 最大深度"

    后序遍历的递归解法：**先钻到最底下，在返回时再做处理**

    ```python
    def maxDepth(root: TreeNode | None) -> int:
        if not root:
            return 0

        left = maxDepth(root.left)
        right = maxDepth(root.right)

        return max(left, right) + 1
    ```

=== "#110 平衡树"

    平衡树就是左右子树高度差不超过 1，所以要基于后序遍历的树深度来求解

    ```python
    def isBalanced(root: TreeNode | None) -> bool:
        def check_height(node: TreeNode | None) -> int:
            if not node:
                return 0

            # 后序遍历：先检查左子树
            left_height = check_height(node.left)
            if left_height == -1:
                return -1  # 左子树不平衡，提前终止

            # 后序遍历：再检查右子树
            right_height = check_height(node.right)
            if right_height == -1:
                return -1  # 右子树不平衡，提前终止

            # 返回时处理：检查当前节点是否平衡
            if abs(left_height - right_height) > 1:
                return -1  # 当前节点不平衡

            # 返回当前节点的高度（自底向上汇总信息）
            return max(left_height, right_height) + 1

        return check_height(root) != -1
    ```

=== "#543 树的直径"

    ```python
    def diameterOfBinaryTree(root: TreeNode | None) -> int:
        max_diameter = 0

        def depth(node: TreeNode | None) -> int:
            nonlocal max_diameter
            if not node:
                return 0

            # 后序遍历：先递归计算左右子树的深度
            left_depth = depth(node.left)
            right_depth = depth(node.right)

            # 当前节点的直径 = 左子树深度 + 右子树深度
            max_diameter = max(max_diameter, left_depth + right_depth)

            # 计算当前节点的深度
            return max(left_depth, right_depth) + 1

        depth(root)

        return max_diameter
    ```

=== "#124 最大路径和"

    与树的直径类似，对于每个节点，`路径和 = 左子树贡献 + 右子树贡献 + 当前节点值`

    ```python
    import math

    def maxPathSum(root: TreeNode | None) -> int:
        max_sum = -math.inf  # 初始化为最小值，因为节点值可能为负

        def max_gain(node: TreeNode | None) -> int:
            nonlocal max_sum
            if not node:
                return 0

            # 后序遍历：先递归计算左右子树的最大贡献
            # 如果子树贡献为负，则不选择该子树（取 0）
            left_gain = max(max_gain(node.left), 0)
            right_gain = max(max_gain(node.right), 0)

            # 以当前节点为"拐点"的路径和
            # 路径和 = 左子树贡献 + 右子树贡献 + 当前节点值
            current_path_sum = left_gain + right_gain + node.val
            max_sum = max(max_sum, current_path_sum)

            # 返回给父节点的最大贡献：只能选择左或右其中一条路径
            # 贡献 = 当前节点值 + max(左子树贡献, 右子树贡献)
            return node.val + max(left_gain, right_gain)

        max_gain(root)
        return max_sum
    ```

=== "#236 最近公共祖先"

    <div class="grid cards" markdown>
    - <figure>
        ![LCA 示例](../imgs/LCA/lowest_common_ancestor_of_a_binary_tree.webp)
        <figcaption>最近公共祖先示例</figcaption>
    </figure>
    - <figure>
        ![LCA 逻辑](../imgs/LCA/lca_logic.webp)
        <figcaption>LCA 判断逻辑</figcaption>
    </figure>
    </div>

    ```python
    def lowestCommonAncestor(root: TreeNode | None, p: TreeNode, q: TreeNode) -> TreeNode | None:
        # 递归终止条件：
        # 1. 搜到底了（None）
        # 2. 找到目标节点（p 或 q）
        if not root or root is p or root is q:
            return root

        # 后序遍历：先递归左右子树
        left = lowestCommonAncestor(root.left, p, q)
        right = lowestCommonAncestor(root.right, p, q)

        # 根据左右子树的返回值判断 LCA 位置
        # 情况1：p 和 q 分散在左右两侧 → 当前节点就是 LCA
        if left and right:
            return root

        # 情况2：p 和 q 都在左子树 → 返回左子树的结果
        if left:
            return left

        # 情况3：p 和 q 都在右子树（或右子树找到一个）→ 返回右子树的结果
        return right
    ```

=== "#235 二叉搜索树的最近公共祖先"

    由于 BST 的有序性，其解法与普通的 LCA 相似：如果 $p$ 和 $q$ 都小于当前节点，则 LCA 在左子树；如果 $p$ 和 $q$ 都大于当前节点，则 LCA 在右子树；否则当前节点就是 LCA。

    ```python
    # 迭代解法(推荐)
    # Time: O(h), Space: O(1)
    def lowestCommonAncestorIterative(root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode | None:
        while root:
            if p.val < root.val and q.val < root.val:
                root = root.left
            elif p.val > root.val and q.val > root.val:
                root = root.right
            else:
                return root
        return None
    ```

    ```python
    # 递归解法
    # Time: O(h), Space: O(h)
    def lowestCommonAncestorRecursive(root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        if p.val < root.val and q.val < root.val:
            return lowestCommonAncestorRecursive(root.left, p, q)
        if p.val > root.val and q.val > root.val:
            return lowestCommonAncestorRecursive(root.right, p, q)
        return root
    ```

### 中序遍历题目

=== "#98 验证二叉搜索树"

    由于需要捕获外部变量，因此可以使用闭包避免二级指针

    ```python
    def isValidBST(root: TreeNode | None) -> bool:
        # prev 记录中序遍历中上一个访问的节点
        prev = None

        def inorder(node: TreeNode | None) -> bool:
            nonlocal prev
            # 递归终止条件：空节点视为有效
            if not node:
                return True

            # 中序遍历：左 -> 根 -> 右
            # 1. 先递归检查左子树
            if not inorder(node.left):
                return False  # 左子树不合法，提前终止

            # 2. 检查当前节点：BST 的中序遍历必须严格递增
            if prev is not None and node.val <= prev.val:
                return False  # 不满足严格递增，不是 BST
            prev = node  # 更新 prev 为当前节点，供下一个节点比较

            # 3. 递归检查右子树
            return inorder(node.right)

        return inorder(root)
    ```

### 前序遍历题目

当节点需要依赖左右子树的信息时，使用前序遍历，这样不仅代码简单，而且高效

=== "#226 翻转二叉树"

    ```python
    def invertTree(root: TreeNode | None) -> TreeNode | None:
      if not root:
          return None

      # 交换当前节点的左右子树
      root.left, root.right = root.right, root.left

      # 递归翻转子树
      invertTree(root.left)
      invertTree(root.right)

      return root
    ```

=== "#101 对称树"

    <div class="grid" markdown>
    <div markdown>
    对称树的定义：

    - 左子树的左节点 == 右子树的右节点

    - 左子树的右节点 == 右子树的左节点
    </div>
    <div markdown>![对称树示意图](../imgs/symmetric_tree.webp)</div>
    </div>

    镜像递归

    ```python
    def isSymmetric(root: TreeNode | None) -> bool:
        if not root:
            return True
        return is_mirror(root.left, root.right)

    def is_mirror(left: TreeNode | None, right: TreeNode | None) -> bool:
        # 递归终止条件
        # 检查节点存在的对称性
        if not left and not right:
            return True
        if not left or not right:
            return False

        # 递归处理逻辑
        # 检查节点值的对称性
        if left.val != right.val:
            return False

        # 递归处理：交叉比较子树（镜像对称）
        return is_mirror(left.left, right.right) and \
            is_mirror(left.right, right.left)
    ```

=== "路径和"

    判断给定的树中是否有和为 targetSum 的路径存在。

    ```python
    def hasPathSum(root: TreeNode | None, target_sum: int) -> bool:
        if not root:
            return False

        # 到达叶子节点
        if not root.left and not root.right:
            return root.val == target_sum

        remaining_sum = target_sum - root.val
        # 只需要左、右子树其中一个满足条件即可
        # 短路求值提前终止
        return hasPathSum(root.left, remaining_sum) or hasPathSum(root.right, remaining_sum)
    ```

=== "复制树"

=== "序列化与反序列化"

### 层序遍历题目

=== "#103 锯齿形层序遍历"

    ```python
    from collections import deque

    # Time: O(n), Space: O(n)
    def zigzagLevelOrder(root: TreeNode | None) -> list[list[int]]:
    	result = []
    	if not root:
    		return result

    	left_to_right = True
    	queue = deque([root])
    	while queue:
    		level_size = len(queue)
    		level = [0] * level_size

    		for i in range(level_size):
    			node = queue.popleft()

    			# 根据方向决定插入位置
    			index = i if left_to_right else level_size - 1 - i
    			level[index] = node.val

    			if node.left:
    				queue.append(node.left)
    			if node.right:
    				queue.append(node.right)

    		result.append(level)
    		left_to_right = not left_to_right

    	return result
    ```

=== "#199 右视图"

    ```python
    from collections import deque

    def rightSideView(root: TreeNode | None) -> list[int]:
        ans = []
        if not root:
            return ans

        queue = deque([root])
        while queue:
            width = len(queue)
            for i in range(width):
                node = queue.popleft()

                # 只把每层的最后一个节点加入结果集
                if i == width - 1:
                    ans.append(node.val)

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

        return ans
    ```

=== "#111 最小深度"

    ```python
    from collections import deque

    def minDepth(root: TreeNode | None) -> int:
        if not root:
            return 0

        depth = 1  # 非空节点的最小深度为 1
        queue = deque([root])

        while queue:
            for _ in range(len(queue)):
                node = queue.popleft()

                # 找到第一个叶子节点，立即返回
                if not node.left and not node.right:
                    return depth

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

            depth += 1

        return depth
    ```

=== "#513 左下角"

    ```python
    from collections import deque

    # Time: O(n), Space: O(n)
    def findBottomLeftValue(root: TreeNode) -> int:
        first_node = 0
        queue = deque([root])
        while queue:
            for i in range(len(queue)):
                node = queue.popleft()

                # 只记录每层的第一个节点
                if i == 0:
                    first_node = node.val

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

        return first_node
    ```

## DFS

=== "#200 岛屿数量"

    可用DFS、BFS、并查集3种解法，但推荐DFS

    ```python
    # Time: O(m×n), Space: O(m×n) 递归栈
    def numIslands(grid: list[list[str]]) -> int:
        if not grid:
            return 0

        def dfs(i: int, j: int) -> None:
            # 边界检查
            if i < 0 or i >= len(grid) or j < 0 or j >= len(grid[0]) or grid[i][j] == '0':
                return

            # 标记为已访问
            grid[i][j] = '0'

            # 递归访问四个方向
            dfs(i-1, j)  # 上
            dfs(i+1, j)  # 下
            dfs(i, j-1)  # 左
            dfs(i, j+1)  # 右

        count = 0
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == '1':
                    count += 1
                    dfs(i, j)

        return count
    ```

## 堆/优先队列

=== "#215 数组中的第K个最大元素"

    使用小顶堆，维护K个最大元素

    ```python
    import heapq

    def findKthLargest(nums: list[int], k: int) -> int:
        # 使用小顶堆，维护 K 个最大元素
        h: list[int] = []
        for num in nums:
            heapq.heappush(h, num)
            if len(h) > k:
                heapq.heappop(h)  # 保持堆大小为 k
        return h[0]
    ```

=== "#23 合并K个升序链表"

=== "#347 前K个高频元素"

## 栈

=== "#20 有效的括号"

    ```python
    # Time: O(n), Space: O(n)
    def isValid(s: str) -> bool:
    	pairs = {')': '(', ']': '[', '}': '{'}

    	stack = []
    	for ch in s:
    		if ch in '([{':
    			stack.append(ch)
    		else:
    			if not stack or stack[-1] != pairs[ch]:
    				return False
    			stack.pop()

    	return len(stack) == 0
    ```

=== "#155 最小栈"

=== "#739 每日温度"

    单调递减栈，当前温度高于栈顶时不断弹出并更新等待天数（出栈时更新结果）。建议手动模拟栈的变化过程来辅助理解。

    ![温度折线图](../imgs/daily_temperatures/example.webp)
    ![单调递减栈](../imgs/daily_temperatures/monotonic_stack.webp)

    ```python
    def dailyTemperatures(temperatures: list[int]) -> list[int]:
    	ans = [0] * len(temperatures)

    	stack: list[int] = []  # 栈内存放 temperatures 的索引
    	for i in range(len(temperatures)):
    		# 如果当前元素 > 栈顶元素，则不断弹出栈中的元素，直至当前元素 < 栈顶元素
    		while stack and temperatures[i] > temperatures[stack[-1]]:
    			top = stack.pop()   # 取栈顶值，即 temperatures 数组中 i 之前的索引
    			ans[top] = i - top  # 将差值更新到对应的位置
    		stack.append(i)  # push

    	return ans
    ```

=== "#84 柱状图中最大的矩形"

=== "#394 字符串解码"

## 队列

=== "滑动窗口最大值"

## 二分查找

=== "#704 二分查找"

    ```python
    def search(nums: list[int], target: int) -> int:
    	left, right = 0, len(nums) - 1

    	# 这里的二分查找的核心在于每次搜索都是以 mid 为单位跳跃
    	while left <= right:
    		mid = left + (right - left) // 2
    		if target < nums[mid]:
    			right = mid - 1  # 左侧区间
    		elif target > nums[mid]:
    			left = mid + 1  # 右侧区间
    		else:
    			return mid  # 找到

    	return -1
    ```

=== "#35 搜索插入位置"

    ```python
    def searchInsert(nums: list[int], target: int) -> int:
    	left, right = 0, len(nums)
    	while left < right:
    		mid = left + (right - left) // 2
    		if target > nums[mid]:
    			left = mid + 1
    		else:
    			right = mid

    	return left  # 返回第一个 >= target 的位置
    ```

=== "#33 搜索旋转排序数组"

=== "在排序数组中查找元素的第一个和最后一个位置"

## 动态规划

=== "#70 爬楼梯"

    ```python
    # Time: O(n), Space: O(1)
    def climbStairs(n: int) -> int:
    	if n <= 2:
    		return n

    	x, y = 1, 2
    	for _ in range(3, n + 1):
    		x, y = y, x + y

    	return y
    ```

=== "#53 最大子数组和"

=== "#300 最长递增子序列"

=== "#1143 最长公共子序列"

=== "最小路径和"

=== "不同路径"

=== "打家劫舍"

=== "#322 零钱兑换"

=== "#42 接雨水"

=== "编辑距离"

=== "最大正方形"

=== "最长有效括号"

## 回溯

=== "#46 全排列"

=== "组合总和"

=== "子集"

=== "括号生成"

## 贪心

=== "#121 买卖股票的最佳时机"

    ```python
    import math

    def maxProfit(prices: list[int]) -> int:
        min_price, max_profit = math.inf, 0
        for price in prices:
            min_price = min(min_price, price)
            max_profit = max(max_profit, price - min_price)
        return max_profit
    ```

=== "#55 跳跃游戏"

    解题思路：可跳跃范围覆盖最后一个元素的索引即可

    ```python
    def canJump(nums: list[int]) -> bool:
        cover = 0
        for i, jump in enumerate(nums):
            # 不可达位置
            if i > cover:
                return False

            # 更新最大覆盖范围
            cover = max(cover, i + jump)  # 从 i 跳 jump 步

            # 是否可以覆盖最后一个元素
            if cover >= len(nums) - 1:
                break

        return True
    ```

=== "#134 加油站"

=== "#11 盛最多水的容器"

=== "分发糖果"

## 递归

编写递归步骤：

1. 明确输入输出
2. 明确递归终止条件（什么时候触底反弹？）
3. 明确递归处理逻辑（每层要做什么事？）
4. 明确递归过程（下楼做还是返回做？）

## 图论

=== "课程表"
