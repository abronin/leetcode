from typing import List

class Solution:
    def busyStudent(self, startTime: List[int], endTime: List[int], queryTime: int) -> int:
        res = 0
        for i in range(len(startTime)):
            if startTime[i] <= queryTime and queryTime <= endTime[i]:
                res += 1
        return res

def test_first():
    s = Solution()
    startTime = [1,2,3]
    endTime = [3,2,7]
    queryTime = 4
    expected = 1
    given = s.busyStudent(startTime, endTime, queryTime)
    assert given == expected
