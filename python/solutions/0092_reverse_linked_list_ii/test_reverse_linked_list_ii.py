from collections.abc import Callable

import pytest
from reverse_linked_list_ii import Solution

from common.list_node import ListNode, build_list, to_list

METHODS = [
    pytest.param(Solution().reverseBetween, id="reverse_between"),
]

CASES = [
    ([1, 2, 3, 4, 5], 2, 4, [1, 4, 3, 2, 5]),
    ([5], 1, 1, [5]),
    ([1, 2, 3], 1, 3, [3, 2, 1]),
    ([1, 2, 3, 4], 1, 2, [2, 1, 3, 4]),
    ([1, 2, 3, 4], 3, 4, [1, 2, 4, 3]),
]


@pytest.mark.parametrize("solution", METHODS)
@pytest.mark.parametrize(("input", "left", "right", "expected"), CASES)
def test_solution(
    solution: Callable[[ListNode | None, int, int], ListNode | None],
    input: list[int],
    left: int,
    right: int,
    expected: list[int],
) -> None:
    assert to_list(solution(build_list(input), left, right)) == expected
