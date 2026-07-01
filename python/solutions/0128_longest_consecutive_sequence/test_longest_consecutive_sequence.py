import pytest
from longest_consecutive_sequence import Solution


@pytest.fixture
def s() -> Solution:
    return Solution()


@pytest.mark.parametrize(
    ("nums", "expected"),
    [
        ([100, 4, 200, 1, 3, 2], 4),
        ([0, 3, 7, 2, 5, 8, 4, 6, 0, 1], 9),
        ([], 0),
        ([1], 1),
        ([1, 2, 3, 4, 5], 5),
        ([5, 4, 3, 2, 1], 5),
        ([1, 3, 5, 7], 1),
        ([1, 1, 2, 2], 2),
    ],
)
def test_longest_consecutive(s: Solution, nums: list[int], expected: int) -> None:
    assert s.longestConsecutive(nums) == expected
