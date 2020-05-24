from typing import List

class Solution:
    def maximum69Number(self, num: int) -> int:
        s = str(num)
        for i in range(len(s)):
            if s[i] == '6':
                s = s[0:i] + '9' + s[i+1:]
                break
        res = int(s)
        return res

def test_first():
    s = Solution()
    num = 9669
    expected = 9969
    given = s.maximum69Number(num)
    assert given == expected
