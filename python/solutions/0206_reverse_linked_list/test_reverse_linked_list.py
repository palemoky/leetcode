from collections.abc import Callable

import pytest
from reverse_linked_list import Solution

from common.list_node import ListNode, build_list, to_list

METHODS = [
    pytest.param(Solution().reverseList, id="reverseList"),
]

CASES = [
    ([1, 2, 3, 4, 5], [5, 4, 3, 2, 1]),
    ([1, 2], [2, 1]),
    ([], []),
]


@pytest.mark.parametrize("solution", METHODS)
@pytest.mark.parametrize(("input", "expected"), CASES)
def test_solution(
    solution: Callable[[ListNode | None], ListNode | None],
    input: list[int],
    expected: list[int],
) -> None:
    assert to_list(solution(build_list(input))) == expected
