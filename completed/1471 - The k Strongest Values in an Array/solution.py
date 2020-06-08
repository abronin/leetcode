from typing import List

class Solution:
    def getStrongest(self, arr: List[int], k: int) -> List[int]:
        s = sorted(arr)
        mindex = int((len(arr) - 1)/2)
        m = s[mindex]
        ss = sorted(s, key=lambda x: abs(x-m)*100000+x, reverse=True)
        return ss[0:k]

def test_first():
    s = Solution()
    arr = [1,2,3,4,5]
    k = 2
    expected = [5,1]
    given = s.getStrongest(arr, k)
    assert given == expected
