from dataclasses import dataclass, field
from typing import Self


@dataclass
class SafetyManualPage:
    page_number: int
    comes_before: set[int] = field(default_factory=set)

    def __lt__(self, other: Self) -> bool:
        return self.page_number not in other.comes_before


def parse_rules(rules: list[list[int]]) -> dict[int, SafetyManualPage]:
    d = {}
    for rule_pair in rules:
        l = rule_pair[0]
        r = rule_pair[1]
        v = d.get(l, SafetyManualPage(page_number=l))
        v.comes_before.add(r)
        d[l] = v
    return d


def main(input_file="input.txt"):
    input = open(input_file).read().split("\n\n")
    rules = [list(map(int, i.split("|"))) for i in input[0].splitlines()]
    page_lists_ = [list(map(int, i.split(","))) for i in input[1].splitlines()]
    pagedefs = parse_rules(rules)

    page_lists = [
        [pagedefs.get(p, SafetyManualPage(page_number=p)) for p in pl_]
        for pl_ in page_lists_
    ]

    # part 1
    sorteds = [pl for pl in page_lists if sorted(pl) == pl]
    print("Part 1:", sum(pl[len(pl) // 2].page_number for pl in sorteds))

    # part 2
    not_sorteds = [pl for pl in page_lists if sorted(pl) != pl]
    sort_them = [sorted(pl) for pl in not_sorteds]
    print("Part 2:", sum(sorted(pl)[len(pl) // 2].page_number for pl in sort_them))


if __name__ == "__main__":
    main()
