from typing import List

class Solution:
    def peopleIndexes(self, favoriteCompanies: List[List[str]]) -> List[int]:
        n = len(favoriteCompanies)
        res = []
        s = []
        for i in range(n):
            s.append(set(favoriteCompanies[i]))
        for i in range(n):
            flag = True
            for j in range(n):
                if i != j and len(s[i]) <= len(s[j]) and s[i].issubset(s[j]):
                    flag = False
                    break
            if flag:
                res.append(i)
        return res

def test_first():
    s = Solution()
    favoriteCompanies = [["leetcode","google","facebook"],["google","microsoft"],["google","facebook"],["google"],["amazon"]]
    expected = [0,1,4]
    given = s.peopleIndexes(favoriteCompanies)
    assert given == expected
