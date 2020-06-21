from typing import List
import re

class Solution:
    def __init__(self):
        self.counter = {}

    def getFolderNames(self, names: List[str]) -> List[str]:
        res = []
        for name in names:
            if name in self.counter:
                self.findNum(name)
                new_name = self.nameWithNum(name, self.counter[name])
                res.append(new_name)
                self.counter[name] += 1
                self.counter[new_name] = 1
            else:
                self.counter[name] = 1
                res.append(name)
        return res

    def findNum(self, name: str):
        new_name = self.nameWithNum(name, self.counter[name])
        while new_name in self.counter:
            self.counter[name] += 1
            new_name = self.nameWithNum(name, self.counter[name])

    def nameWithNum(self, name: str, num: int) -> str:
        return name + '(' + str(num) + ')'

def test_first():
    s = Solution()
    names = ["pes","fifa","gta","pes(2019)"]
    expected = ["pes","fifa","gta","pes(2019)"]
    given = s.getFolderNames(names)
    assert given == expected

def test_second():
    s = Solution()
    names = ["gta","gta(1)","gta","avalon"]
    expected = ["gta","gta(1)","gta(2)","avalon"]
    given = s.getFolderNames(names)
    assert given == expected

def test_third():
    s = Solution()
    names = ["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece"]
    expected = ["onepiece","onepiece(1)","onepiece(2)","onepiece(3)","onepiece(4)"]
    given = s.getFolderNames(names)
    assert given == expected

def test_forth():
    s = Solution()
    names = ["wano","wano","wano","wano"]
    expected = ["wano","wano(1)","wano(2)","wano(3)"]
    given = s.getFolderNames(names)
    assert given == expected

def test_fifth():
    s = Solution()
    names = ["kaido","kaido(1)","kaido","kaido(1)"]
    expected = ["kaido","kaido(1)","kaido(2)","kaido(1)(1)"]
    given = s.getFolderNames(names)
    assert given == expected

def test_six():
    s = Solution()
    names = ["kaido","kaido(1)","kaido","kaido(1)","kaido(2)"]
    expected = ["kaido","kaido(1)","kaido(2)","kaido(1)(1)","kaido(2)(1)"]
    given = s.getFolderNames(names)
    assert given == expected

