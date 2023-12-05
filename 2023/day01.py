import os
import re

DIGITS = ["0","1","2","3","4","5","6","7","8","9"]
def find_digit(line: str) -> str:
    return next((i for i in line if i in DIGITS))

input_file = os.path.join(os.path.dirname(__file__), "inputs/day01.txt")

mappings = {
    "0" : "0",
    "1" : "1",
    "2" : "2",
    "3" : "3",
    "4" : "4",
    "5" : "5",
    "6" : "6",
    "7" : "7",
    "8" : "8",
    "9" : "9",
    "zero" : "0",
    "one" : "1",
    "two" : "2",
    "three" : "3",
    "four" : "4",
    "five" : "5",
    "six" : "6",
    "seven" : "7",
    "eight" : "8",
    "nine" : "9"
}

match_pattern = '|'.join(f"{k}" for k in mappings.keys()) 
regex = re.compile(match_pattern)
regex_reversed = re.compile(match_pattern[::-1])

acc_1 = 0
acc_2 = 0
with open(input_file,'r') as f:
    for line in f.readlines():
        acc_1 += int(
            find_digit(line) + find_digit(line[::-1])
        )
        first = regex.search(line)[0]
        last = regex_reversed.search(line[::-1])[0][::-1]
        res = int(mappings[first] + mappings[last])
        acc_2 += res


print("Part One:", acc_1)
print("Part Two:", acc_2)

