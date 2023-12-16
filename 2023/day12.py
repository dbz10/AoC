from itertools import groupby
from collections import Counter

input_file = "inputs/day12.txt"

conditions = []
blocks = []
expansion_factor = 5
for line in open(input_file):
    l, r = line.split(" ")
    conditions.append("?".join([l] * expansion_factor))
    blocks.append([int(z.strip()) for z in r.split(",")] * expansion_factor)


def is_valid(proposal, blocks):
    if "?" in proposal:
        return False
    decomposition = [len(list(g)) for k, g in groupby(proposal) if k == "#"]
    return decomposition == blocks


def is_possible(proposal: str, blocks: list[int]):
    pivot = proposal.find("?")
    if pivot == -1:
        return is_valid(proposal, blocks)

    fixed = list(proposal)[:pivot]

    decomposition = [len(list(g)) for k, g in groupby(fixed) if k == "#"]
    ld = len(decomposition)

    if ld == 0:
        return True
    if ld > len(blocks):
        return False

    ld = len(decomposition)
    completed_islands_match = decomposition[:-1] == blocks[: ld - 1]
    latest_island_is_possible = decomposition[-1] <= blocks[ld - 1]

    return completed_islands_match and latest_island_is_possible


hits = 0
calls = 0


def gen(condition: str, blocks, memo) -> list[str]:
    # print("")
    global hits, calls
    calls += 1

    idx = condition.find("?")
    if idx < 0:
        return int(is_valid(condition, block))
    # print("c", condition)
    res = 0

    if not is_possible(condition, blocks):
        # print("not possible\n")
        return 0

    # key = (
    #     condition[idx:]
    #     # + str(Counter(condition[:]))
    #     # + str([len(list(g)) for k, g in groupby(condition[:idx]) if k == "#"])
    #     # + str(Counter(condition[idx:]))
    #     # + "".join(sorted(condition[:idx]))
    #     # + condition[idx - 1]
    # )
    # print(condition)
    # print(key)
    # check = memo.get(key)
    # if check:
    #     hits += 1
    #     return check
    l = [x for x in condition]
    l[idx] = "."

    key_l = (
        str([len(list(g)) for k, g in groupby(l[: idx + 1]) if k == "#"])
        + "."
        + condition[idx + 1 :]
    )

    # print("l", key_l)
    check_l = memo.get(key_l)
    if check_l is None:
        l = "".join(l)
        memo[key_l] = gen(l, blocks, memo)

    r = [x for x in condition]
    r[idx] = "#"
    key_r = (
        str([len(list(g)) for k, g in groupby(r[: idx + 1]) if k == "#"])
        + "#"
        + condition[idx + 1 :]
    )
    # print("c", condition)
    # print("r", key_r)
    # print(condition, idx, condition[idx])
    # print("".join(r))
    check_r = memo.get(key_r)
    if check_r is None:
        r = "".join(r)
        memo[key_r] = gen(r, blocks, memo)

    res = memo[key_l] + memo[key_r]
    # print("")
    return res


acc = 0
acc_2 = 0
for condition, block in zip(conditions, blocks):
    memo = {}
    hits = 0
    # print(condition)
    combs = gen(condition, block, memo)

    print(condition, combs)
    acc += combs
    # break


print("Part 1:", acc)
