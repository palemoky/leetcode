import pytest
from group_anagrams import Solution


@pytest.fixture
def s() -> Solution:
    return Solution()


def normalize(groups: list[list[str]]) -> set[frozenset[str]]:
    # Group order and within-group order are not guaranteed.
    return {frozenset(group) for group in groups}


CASES = [
    (
        ["eat", "tea", "tan", "ate", "nat", "bat"],
        [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]],
    ),
    ([""], [[""]]),
    (["a"], [["a"]]),
    ([], []),
    (["abc", "bca", "cab"], [["abc", "bca", "cab"]]),
    (["abc", "def", "ghi"], [["abc"], ["def"], ["ghi"]]),
    (["", "", "b"], [["", ""], ["b"]]),
]


@pytest.mark.parametrize(("strs", "expected"), CASES)
@pytest.mark.parametrize("method", ["groupAnagramsSort", "groupAnagramsCounter"])
def test_group_anagrams(
    s: Solution, method: str, strs: list[str], expected: list[list[str]]
) -> None:
    func = getattr(s, method)
    assert normalize(func(strs)) == normalize(expected)
