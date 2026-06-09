import pytest
from two_sum import Solution

CASES = [
    ([2, 7, 11, 15], 9, [0, 1]),
    ([3, 2, 4], 6, [1, 2]),
    ([3, 3], 6, [0, 1]),
    ([1, 2], 3, [0, 1]),
    ([0, 4, 3, 0], 0, [0, 3]),
]


@pytest.fixture
def s() -> Solution:
    return Solution()


@pytest.mark.parametrize(("nums", "target", "expected"), CASES)
def test_two_sum_brute_force(
    s: Solution, nums: list[int], target: int, expected: list[int]
) -> None:
    assert s.twoSumBruteForce(nums, target) == expected


@pytest.mark.parametrize(("nums", "target", "expected"), CASES)
def test_two_sum_hash_map(s: Solution, nums: list[int], target: int, expected: list[int]) -> None:
    assert s.twoSumHashMap(nums, target) == expected
