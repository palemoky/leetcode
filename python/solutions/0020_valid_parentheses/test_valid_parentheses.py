import pytest
from valid_parentheses import Solution


@pytest.fixture
def s() -> Solution:
    return Solution()


@pytest.mark.parametrize(
    ("text", "expected"),
    [
        ("()", True),
        ("()[]{}", True),
        ("{[()]}", True),
        ("(]", False),
        ("([)]", False),
        ("(", False),
        ("]", False),
        ("", True),
        ("(((", False),
        ("())", False),
        ("([])", True),
    ],
)
def test_is_valid_stack(s: Solution, text: str, expected: bool) -> None:
    assert s.isValidStack(text) == expected
