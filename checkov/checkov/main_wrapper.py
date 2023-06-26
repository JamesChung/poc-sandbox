import sys
from main import Checkov


class CheckovWrapper(Checkov):
    def __init__(self, argv: list[str] = sys.argv[1:]):
        # TODO
        print("Hello from CheckovWrapper")
        print(argv)
        super().__init__(argv)


def main():
    checkov = CheckovWrapper()
    checkov.run()


if __name__ == "__main__":
    main()
