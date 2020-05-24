from typing import List

class Solution:
    def printVertically(self, s: str) -> List[str]:
        arr = s.split()
        m = len(max(arr, key=lambda x: len(x)))
        res = [""]*m
        for word in arr:
            for i in range(m):
                if i < len(word):
                    res[i] += word[i]
                else:
                    res[i] += " "
        for i in range(len(res)):
            res[i] = res[i].rstrip()
        return res

def test_first():
    s = Solution()
    st = "HOW ARE YOU"
    expected = ["HAY","ORO","WEU"]
    given = s.printVertically(st)
    assert given == expected

def test_second():
    s = Solution()
    st = "TO BE OR NOT TO BE"
    expected = ["TBONTB","OEROOE","   T"]
    given = s.printVertically(st)
    assert given == expected

def test_three():
    s = Solution()
    st = "CONTEST IS COMING"
    expected = ["CIC","OSO","N M","T I","E N","S G","T"]
    given = s.printVertically(st)
    assert given == expected
