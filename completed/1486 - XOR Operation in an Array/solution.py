from typing import List

class Solution:
    def xorOperation(self, n: int, start: int) -> int:
        res = 0
        for i in range(n):
            res = res^(start+i*2)
        return res

def test_first():
    s = Solution()
    n = 4
    start = 3
    expected = 8
    given = s.xorOperation(n, start)
    assert given == expected
