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


import time

hits = 0
calls = 0


def gen(condition: str, blocks, memo) -> list[str]:
    global hits, calls
    calls += 1

    if not is_possible(condition, block):
        return 0
    if is_valid(condition, block):
        return 1

    if "?" in condition:
        # print(condition)
        idx = condition.find("?")
        key = condition[idx:]
        # print(condition)
        # print(key)
        # print("")
        # time.sleep(1)
        check = memo.get(key)
        if check:
            hits += 1
            return check
        # print(idx)
        l = list(condition)
        l[idx] = "."
        l = "".join(l)
        r = list(condition)
        r[idx] = "#"
        r = "".join(r)
        res = gen(l, blocks, memo) + gen(r, blocks, memo)
        memo[key] = res
        return res

    return 0


acc = 0
acc_2 = 0
for condition, block in zip(conditions, blocks):
    memo = {}
    hits = 0
    combs = gen(condition, block, memo)

    print(condition, combs, hits, calls, len(memo))
    acc += combs


print("Part 1:", acc)
