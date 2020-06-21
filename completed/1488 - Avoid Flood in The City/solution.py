from typing import List

class Solution:
    def avoidFlood(self, rains: List[int]) -> List[int]:
        m = {}
        p = []
        res = []
        i = 0
        for i in range(len(rains)):
            lake = rains[i]
            if lake == 0:
                p.append(i)
                res.append(1)
                continue
            if lake in m and m[lake] > -1:
                if len(p) == 0: return []
                pos = self.findPosition(p, m[lake])
                if pos < 0: return []
                position = p.pop(pos)
                res[position] = lake
                res.append(-1)
                m[lake] = i
            else:
                res.append(-1)
                m[lake] = i
        return res

    def findPosition(self, positions: List[int], minimal: int) -> int:
        for i in range(len(positions)):
            if positions[i] > minimal:
                return i
        return -1

def test_first():
    s = Solution()
    rains = [1,2,0,0,2,1]
    expected = [-1,-1,2,1,-1,-1]
    # rains = [1,2,3,4]
    # expected = [-1,-1,-1,-1]
    # rains = [1,2,0,1,2]
    # expected = []
    # rains = [0,1,1]
    # expected = []
    # rains = [2,3,0,0,3,1,0,1,0,2,2]
    # expected = []
    given = s.avoidFlood(rains)
    assert given == expected
