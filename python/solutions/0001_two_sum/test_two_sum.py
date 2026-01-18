import unittest

from two_sum import Solution


class TestTwoSum(unittest.TestCase):
    def setUp(self) -> None:
        self.s = Solution()

    def test_two_sum_brute_force(self) -> None:
        self.assertEqual(self.s.twoSumBruteForce([2, 7, 11, 15], 9), [0, 1])
        self.assertEqual(self.s.twoSumBruteForce([3, 2, 4], 6), [1, 2])
        self.assertEqual(self.s.twoSumBruteForce([3, 3], 6), [0, 1])
        self.assertEqual(self.s.twoSumBruteForce([1, 2], 3), [0, 1])
        self.assertEqual(self.s.twoSumBruteForce([0, 4, 3, 0], 0), [0, 3])

    def test_two_sum_hash_map(self) -> None:
        self.assertEqual(self.s.twoSumHashMap([2, 7, 11, 15], 9), [0, 1])
        self.assertEqual(self.s.twoSumHashMap([3, 2, 4], 6), [1, 2])
        self.assertEqual(self.s.twoSumHashMap([3, 3], 6), [0, 1])
        self.assertEqual(self.s.twoSumHashMap([1, 2], 3), [0, 1])
        self.assertEqual(self.s.twoSumHashMap([0, 4, 3, 0], 0), [0, 3])


if __name__ == "__main__":
    unittest.main()
