from typing import List

class Solution:
    def buildArray(self, target: List[int], n: int) -> List[str]:
        res = []
        current = 1
        for i in target:
            if current < i:
                for j in range(i - current):
                    res.append("Push")
                    res.append("Pop")
                res.append("Push")
            elif i == current:
                res.append("Push")
            current = i + 1
        return res

def test_first():
    s = Solution()
    target = [1,3]
    n = 3
    expected = ["Push","Push","Pop","Push"]
    given = s.buildArray(target, n)
    assert given == expected
