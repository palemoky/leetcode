from collections.abc import Callable

import pytest
from group_anagrams import Solution

METHODS = [
    pytest.param(Solution().groupAnagramsSort, id="sort"),
    pytest.param(Solution().groupAnagramsCounter, id="counter"),
]


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


@pytest.mark.parametrize("group_anagrams", METHODS)
@pytest.mark.parametrize(("strs", "expected"), CASES)
def test_group_anagrams(
    group_anagrams: Callable[[list[str]], list[list[str]]],
    strs: list[str],
    expected: list[list[str]],
) -> None:
    assert normalize(group_anagrams(strs)) == normalize(expected)
