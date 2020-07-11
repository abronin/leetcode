from typing import List

class Solution:
    def minDays(self, bloomDay: List[int], m: int, k: int) -> int:
        # print(bloomDay)

        n = len(bloomDay)
        if n < m*k:
            return -1

        tm = {}
        for i in bloomDay:
            if i in tm:
                tm[i] += 1
            else:
                tm[i] = 1
        # print(tm)

        days = sorted(tm.keys())
        # print(days)

        start_day = 0
        limit = m*k
        for i in days:
            limit -= tm[i]
            if limit <= 0:
                break
            start_day += 1
        
        if self.checkDay(bloomDay, m, k, days[start_day]):
            return days[start_day]
        days = days[start_day:]

        day = self.findDay(bloomDay, m, k, days)
        return day

    def findDay(self, bloomDay, m, k, days) -> int:
        n = len(days)
        if n == 1:
            return days[0]
        if n == 2:
            return days[1]
        med = int(n/2)
        if self.checkDay(bloomDay, m, k, days[med]):
            days = days[:med+1]
        else:
            days = days[med:]
        return self.findDay(bloomDay, m, k, days)

    def checkDay(self, bloomDay: List[int], m: int, k: int, day: int) -> bool:
        # print(day)
        segment = 0
        for i in bloomDay:
            if i > day:
                m -= self.segmentValue(segment, k)
                segment = 0
            else:
                segment += 1
            if m <= 0:
                return True
        m -= self.segmentValue(segment, k)
        if m <= 0:
            return True
        return False

    def segmentValue(self, segment: int, k: int) -> int:
        return int(segment/k)


def test_first():
    s = Solution()
    bloomDay = [1,10,3,10,2]
    m = 3
    k = 1
    expected = 3
    given = s.minDays(bloomDay, m, k)
    assert given == expected

def test_second():
    s = Solution()
    bloomDay = [1,10,3,10,2]
    m = 3
    k = 2
    expected = -1
    given = s.minDays(bloomDay, m, k)
    assert given == expected
