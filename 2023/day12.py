from itertools import groupby
from collections import Counter

input_file = "inputs/day12_test.txt"

conditions = []
blocks = []
expansion_factor = 5
for line in open(input_file):
    l, r = line.split(" ")
    conditions.append("?".join([l] * expansion_factor))
    blocks.append([int(z.strip()) for z in r.split(",")] * expansion_factor)


def is_valid(proposal, blocks):
    decomposition = [len(list(g)) for k, g in groupby(proposal) if k == "#"]
    return decomposition == blocks


def is_possible(proposal: str, blocks: list[int]):
    pivot = proposal.find("?")
    if pivot == -1:
        return is_valid(proposal, blocks)

    fixed = list(proposal)[:pivot]
    remainder = list(proposal)[pivot:]
    decomposition = [len(list(g)) for k, g in groupby(fixed) if k == "#"]

    if len(decomposition) == 0:
        return True
    if len(decomposition) > len(blocks):
        return False

    # possible_remaining_broken = len([z for z in remainder if z == "#" or z == "?"])
    # needed_remaining_broken = sum(block) - sum(decomposition)
    # if possible_remaining_broken < needed_remaining_broken:
    #     return False

    ld = len(decomposition)
    completed_islands_match = decomposition[:-1] == blocks[: ld - 1]
    latest_island_is_possible = decomposition[-1] <= blocks[ld - 1]

    return completed_islands_match and latest_island_is_possible


def gen(condition: str, blocks) -> list[str]:
    if "?" in condition:
        idx = condition.find("?")
        l = list(condition)
        l[idx] = "."
        l = "".join(l)
        r = list(condition)
        r[idx] = "#"
        r = "".join(r)
        res = []
        if is_possible(l, blocks):
            # print(l)
            res += gen(l, blocks)
        if is_possible(r, blocks):
            # print(r)
            res += gen(r, blocks)
        return res
    if is_valid(condition, block):
        return [condition]
    return []


# bfs?
def gen_bfs(go, blocks):
    while any("?" in x for x in go):
        r = []
        for cur in go:
            s = list(cur)
            for i, ch in enumerate(s):
                if ch == "?":
                    cur[i] = "."
                    if is_possible("".join(cur), blocks):
                        r.append(cur)
                    cur[i] = "#"
                    if is_possible("".join(cur), blocks):
                        r.append(cur)
        return go(r)


acc = 0
acc_2 = 0
for condition, block in zip(conditions, blocks):
    combs = len([c for c in gen(condition, block) if is_valid(c, block)])

    print(condition, combs)
    acc += combs


print("Part 1:", acc)
