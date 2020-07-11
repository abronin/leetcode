from typing import List

class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        res = []
        sum = 0
        for i in nums:
            sum += i
            res.append(sum)
        return res

def test_first():
    s = Solution()
    nums = [1,2,3,4]
    expected = [1,3,6,10]
    given = s.runningSum(nums)
    assert given == expected
