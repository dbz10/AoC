import os

DIGITS = ["0","1","2","3","4","5","6","7","8","9"]
def find_digit(line: str) -> str:
    return next((i for i in line if i in DIGITS))

input_file = os.path.join(os.path.dirname(__file__), "inputs/day01.txt")

acc = 0
with open(input_file,'r') as f:
    for line in f.readlines():
        acc += int(
            find_digit(line) + find_digit(line[::-1])
        )

print("Part One:", acc)