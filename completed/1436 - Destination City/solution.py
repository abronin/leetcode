from typing import List

class Solution:
    def destCity(self, paths: List[List[str]]) -> str:
        starts = list(map(self.first, paths))
        for pair in paths:
            if pair[1] not in starts:
                return pair[1]
    
    def first(self, arr: List[str]) -> List[str]:
        return arr[0]

def test_first():
    s = Solution()
    input = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
    expected = "Sao Paulo"
    given = s.destCity(input)
    assert given == expected

def test_one():
    s = Solution()
    input = [["A","B"]]
    expected = "B"
    given = s.destCity(input)
    assert given == expected
