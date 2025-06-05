use std::collections::HashMap;
impl Solution {
    pub fn two_sum_bruteforce(nums: Vec<i32>, target: i32) -> Vec<i32> {
        for i in 0..nums.len() { 
            for j in i + 1..nums.len() { 
                if nums[i] + nums[j] == target { 
                    return vec![i as i32, j as i32]; 
                }
            }
        }
        unreachable!() 
    }

     pub fn two_sum_hashmap(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut idx = HashMap::new(); 
        for (j, &x) in nums.iter().enumerate() { 
            if let Some(&i) = idx.get(&(target - x)) { 
                return vec![i as i32, j as i32];
            }
            idx.insert(x, j);
        }
        unreachable!() 
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    struct TestCase {
        nums: Vec<i32>,
        target: i32,
        want: Vec<i32>,
    }

    #[test]
    fn test_two_sum_all() {
        let test_cases = [
            TestCase { nums: vec![2, 7, 11, 15], target: 9, want: vec![0, 1] },
            TestCase { nums: vec![3, 2, 4], target: 6, want: vec![1, 2] },
            TestCase { nums: vec![3, 3], target: 6, want: vec![0, 1] },
            TestCase { nums: vec![1, 2], target: 3, want: vec![0, 1] },
            TestCase { nums: vec![0, 4, 3, 0], target: 0, want: vec![0, 3] },
        ];

        for case in test_cases.iter() {
            let got1 = Solution::two_sum_bruteforce(case.nums.clone(), case.target);
            let got2 = Solution::two_sum_hashmap(case.nums.clone(), case.target);
            assert_eq!(got1, case.want, "bruteforce failed for nums={:?}, target={}", case.nums, case.target);
            assert_eq!(got2, case.want, "hashmap failed for nums={:?}, target={}", case.nums, case.target);
        }
    }
}
