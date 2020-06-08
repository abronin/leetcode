from typing import List

class BrowserHistory:
    def __init__(self, homepage: str):
        self.history = []
        self.position = 0
        self.visit(homepage)
    def visit(self, url: str) -> None:
        self.history = self.history[0 : self.position]
        self.history.append(url)
        self.position += 1
    def back(self, steps: int) -> str:
        if steps > self.position - 1:
            self.position = 1
        else:
            self.position -= steps
        return self.history[self.position-1]
    def forward(self, steps: int) -> str:
        if steps > len(self.history) - self.position:
            self.position = len(self.history)
        else:
            self.position += steps
        return self.history[self.position-1]
