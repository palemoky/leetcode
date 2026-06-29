import pytest
from subarray_sum_equals_k import Solution


@pytest.fixture
def s() -> Solution:
    return Solution()


@pytest.mark.parametrize(
    ("nums", "k", "expected"),
    [
        ([1, 1, 1], 2, 2),
        ([1, 2, 3], 3, 2),
        ([1], 1, 1),
        ([1], 0, 0),
        ([0, 0, 0], 0, 6),
        ([1, -1, 1, -1], 0, 4),
        ([-1, -1, 1], 0, 1),
        ([3, 4, 7, 2, -3, 1, 4, 2], 7, 4),
        ([1, 2, 3], 7, 0),
        ([-1, -1, 1], -1, 3),
    ],
)
def test_subarray_sum(s: Solution, nums: list[int], k: int, expected: int) -> None:
    assert s.subarraySum(nums, k) == expected
