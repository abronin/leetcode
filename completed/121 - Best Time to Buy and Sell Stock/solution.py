from typing import List

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) == 0:
            return 0
        min = prices[0]
        max = prices[0]
        res = 0
        for price in prices:
            if price < min:
                min = price
                max = price
            if price > max:
                max = price
            if res < max-min:
                res = max-min
        return res

def test_first():
    s = Solution()
    input = [7,1,5,3,6,4]
    expected = 5
    given = s.maxProfit(input)
    assert given == expected

def test_second():
    s = Solution()
    input = [7,6,4,3,1]
    expected = 0
    given = s.maxProfit(input)
    assert given == expected

def test_one():
    s = Solution()
    input = [7]
    expected = 0
    given = s.maxProfit(input)
    assert given == expected

def test_empty():
    s = Solution()
    input = []
    expected = 0
    given = s.maxProfit(input)
    assert given == expected
