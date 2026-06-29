from common.list_node import ListNode


class Solution:
    # Solution 1: 迭代，逐个把指针反向
    # Time: O(n), Space: O(1)
    def reverseList(self, head: ListNode | None) -> ListNode | None:
        prev = None
        while head:
            next = head.next
            head.next = prev
            prev = head
            head = next

        return prev
