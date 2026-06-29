import pytest
from container_with_most_water import Solution

CASES = [
    ([1, 8, 6, 2, 5, 4, 8, 3, 7], 49),
    ([1, 1], 1),
    ([4, 3, 2, 1, 4], 16),
    ([1, 2, 1], 2),
    ([2, 3, 4, 5, 18, 17, 6], 17),
    ([1, 2, 4, 3], 4),
]


@pytest.fixture
def s() -> Solution:
    return Solution()


@pytest.mark.parametrize(("height", "expected"), CASES)
def test_max_area(s: Solution, height: list[int], expected: int) -> None:
    assert s.maxArea(height) == expected
