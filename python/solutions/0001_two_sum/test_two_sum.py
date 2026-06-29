from collections.abc import Callable

import pytest
from two_sum import Solution

CASES = [
    ([2, 7, 11, 15], 9, [0, 1]),
    ([3, 2, 4], 6, [1, 2]),
    ([3, 3], 6, [0, 1]),
    ([1, 2], 3, [0, 1]),
    ([0, 4, 3, 0], 0, [0, 3]),
]


@pytest.fixture(params=["twoSumBruteForce", "twoSumHashMap"])
def two_sum(request: pytest.FixtureRequest) -> Callable[[list[int], int], list[int]]:
    return getattr(Solution(), request.param)


@pytest.mark.parametrize(("nums", "target", "expected"), CASES)
def test_two_sum(
    two_sum: Callable[[list[int], int], list[int]],
    nums: list[int],
    target: int,
    expected: list[int],
) -> None:
    assert two_sum(nums, target) == expected
