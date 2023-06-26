import sys
from main import Checkov


class CheckovWrapper(Checkov):
    def __init__(self, argv: list[str] = sys.argv[1:]):
        print("Hello from CheckovWrapper")
        print(argv)
        super().__init__(argv)
