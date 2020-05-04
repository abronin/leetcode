from typing import List

class Solution:
    def longestSubarray(self, nums: List[int], limit: int) -> int:
        n = len(nums)
        if n == 0:
            return 0
        mindex = 0
        maxdex = 0
        start = 0
        minval = nums[mindex]
        maxval = nums[maxdex]
        counter = 0
        for end in range(0, n):
            current = nums[end]
            if minval > current:
                mindex = end
                minval = current
            if maxval < current:
                maxdex = end
                maxval = current
            if maxval-minval <= limit and counter < end - start + 1:
                counter = end - start + 1
            else:
                for newstart in range(start, end+1):
                    sl = nums[slice(newstart, end+1)]
                    if maxdex < newstart:
                        maxval = max(sl)
                        maxdex = newstart + sl.index(maxval)
                    minval = min(sl)
                    mindex = newstart + sl.index(minval)
                    if maxval-minval <= limit:
                        start = newstart
                        break
        return counter

def test_first():
    s = Solution()
    nums = [8,2,4,7]
    limit = 4
    expected = 2
    given = s.longestSubarray(nums, limit)
    assert given == expected

def test_second():
    s = Solution()
    nums = [10,1,2,4,7,2]
    limit = 5
    expected = 4
    given = s.longestSubarray(nums, limit)
    assert given == expected

def test_third():
    s = Solution()
    nums = [4,2,2,2,4,4,2,2]
    limit = 0
    expected = 3
    given = s.longestSubarray(nums, limit)
    assert given == expected

# s = Solution()
# nums = [8,2,4,7]
# limit = 4
# expected = 2
# given = s.longestSubarray(nums, limit)
