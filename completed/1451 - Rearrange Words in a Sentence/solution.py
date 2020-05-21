from typing import List

class Solution:
    def arrangeWords(self, text: str) -> str:
        arr = sorted(text.split(), key=lambda str: len(str))
        res = " ".join(arr).lower().capitalize()
        return res

def test_first():
    s = Solution()
    text = "Leetcode is cool"
    expected = "Is cool leetcode"
    given = s.arrangeWords(text)
    assert given == expected
