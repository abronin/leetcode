from typing import List

class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        res = []
        for i in range(n):
            res.append(nums[i])
            res.append(nums[n+i])
        return res

def test_first():
    s = Solution()
    nums = [2,5,1,3,4,7]
    n = 3
    expected = [2,3,5,4,1,7]
    given = s.shuffle(nums, n)
    assert given == expected
