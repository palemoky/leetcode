from collections.abc import Iterable


class ListNode:
    def __init__(self, val: int = 0, next: ListNode | None = None) -> None:
        self.val = val
        self.next = next


def build_list(values: Iterable[int]) -> ListNode | None:
    """由可迭代对象构造链表，返回头节点（空时返回 None）。"""
    dummy = ListNode()
    curr = dummy
    for v in values:
        curr.next = ListNode(v)
        curr = curr.next
    return dummy.next


def to_list(head: ListNode | None) -> list[int]:
    """把链表展开成 Python list，便于断言比较。"""
    result: list[int] = []
    while head:
        result.append(head.val)
        head = head.next
    return result
