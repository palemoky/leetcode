import unittest

from valid_parentheses import Solution


class TestValidParentheses(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_is_valid_stack(self) -> None:
        self.assertTrue(self.s.isValidStack("()"))
        self.assertTrue(self.s.isValidStack("()[]{}"))
        self.assertTrue(self.s.isValidStack("{[()]}"))
        self.assertFalse(self.s.isValidStack("(]"))
        self.assertFalse(self.s.isValidStack("([)]"))
        self.assertFalse(self.s.isValidStack("("))
        self.assertFalse(self.s.isValidStack("]"))
        self.assertTrue(self.s.isValidStack(""))
        self.assertFalse(self.s.isValidStack("((("))
        self.assertFalse(self.s.isValidStack("())"))
        self.assertTrue(self.s.isValidStack("([])"))


if __name__ == "__main__":
    unittest.main()
