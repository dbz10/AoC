from collections import defaultdict
from copy import copy
from functools import cmp_to_key


def main(input_file="sample.txt"):
    input = open(input_file).read().split("\n\n")
    rules = [i.split("|") for i in input[0].splitlines()]
    page_lists = [i.split(",") for i in input[1].splitlines()]

    rules_packed = defaultdict(set)
    for rule_pair in rules:
        rules_packed[rule_pair[0]].add(rule_pair[1])
    print(f"Part 1: {part1(page_lists, rules_packed)}")
    print(f"Part 2: {part2(page_lists, rules_packed)}")


def part1(pages, rules):
    pages_ok = [p for p in pages if ok1(copy(p), rules)]
    return sum(int(p[len(p) // 2]) for p in pages_ok)


def ok1(pages: list[str], rules):
    while pages:
        page = pages.pop()
        if len(set(pages) & rules[page]) > 0:
            return False
    return True


def part2(pages, rules):
    # brute force? sorting black magic?
    pages_not_ok = [p for p in pages if not ok1(copy(p), rules)]
    sortkey = cmp_to_key(cmp(rules))
    fixed = [sorted(p, key=sortkey) for p in pages_not_ok]
    return sum(int(p[len(p) // 2]) for p in fixed)


def cmp(rules):
    def cmp_s(s1: str, s2: str) -> int:
        if s1 in rules[s2]:
            return -1
        elif s2 in rules[s1]:
            return 1
        return 0

    return cmp_s


if __name__ == "__main__":
    main()
