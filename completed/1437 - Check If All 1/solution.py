from typing import List

class Solution:
    def kLengthApart(self, nums: List[int], k: int) -> bool:
        counter = k
        for n in nums:
            if n == 0:
                counter += 1
            if n == 1:
                if counter < k:
                    return False
                counter = 0
            print(n, counter)
        return True

def test_first():
    s = Solution()
    nums = [1,0,0,0,1,0,0,1]
    k = 2
    expected = True
    given = s.kLengthApart(nums, k)
    assert given == expected

def test_zero():
    s = Solution()
    nums = [1,1,1,1,1]
    k = 0
    expected = True
    given = s.kLengthApart(nums, k)
    assert given == expected

def test_second():
    s = Solution()
    nums = [1,0,0,0,1,0,0,1]
    k = 2
    expected = True
    given = s.kLengthApart(nums, k)
    assert given == expected

# s = Solution()
# nums = [1,0,0,0,1,0,0,1]
# k = 2
# expected = True
# given = s.kLengthApart(nums, k)
