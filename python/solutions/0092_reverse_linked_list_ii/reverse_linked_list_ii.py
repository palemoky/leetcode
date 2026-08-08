from common.list_node import ListNode


class Solution:
    # Solution 1:
    # Time: O(), Space: O()
    def reverseBetween(self, head: ListNode | None, left: int, right: int) -> ListNode | None:
        dummy = ListNode(next=head)

        prev = dummy
        for _ in range(1, left):
            assert prev.next is not None
            prev = prev.next

        curr = prev.next
        assert curr is not None
        for _ in range(left, right):
            nxt = curr.next
            assert nxt is not None
            curr.next = nxt.next
            nxt.next = prev.next
            prev.next = nxt

        return dummy.next
