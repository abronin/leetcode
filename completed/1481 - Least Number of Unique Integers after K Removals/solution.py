from typing import List

class Solution:
    def findLeastNumOfUniqueInts(self, arr: List[int], k: int) -> int:
        m = {}
        c = {}
        for i in arr:
            if i in m:
                m[i] += 1
                if m[i] in c:
                    c[m[i]] += 1 
                else:
                    c[m[i]] = 1 
                c[m[i] - 1] -= 1
                if c[m[i] - 1] == 0:
                    del c[m[i] - 1]
            else:
                if 1 in c:
                    c[1] += 1
                else:
                    c[1] = 1
                m[i] = 1
        res = len(m)
        ks = sorted(c.keys())
        for i in ks:
            possible = i*c[i]
            if k >= possible:
                res -= c[i]
            else:
                res -= int(k/i)
            k -= possible
            if k <= 0:
                return res
        return res

def test_first():
    s = Solution()
    arr = [5,5,4]
    k = 1
    expected = 1
    given = s.findLeastNumOfUniqueInts(arr, k)
    assert given == expected

def test_second():
    s = Solution()
    arr = [4,3,1,1,3,3,2]
    k = 3
    expected = 2
    given = s.findLeastNumOfUniqueInts(arr, k)
    assert given == expected
