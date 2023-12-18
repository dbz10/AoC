from functools import cache
import math

input_file = "inputs/day17_test.txt"

grid = {
    x + y * 1j: int(s)
    for (y, line) in enumerate(open(input_file).read().split("\n"))
    for x, s in enumerate(line)
}

Lx = int(math.sqrt(len(grid)))
Ly = int(math.sqrt(len(grid)))

# print(grid)
import time


# def go(pdc, memo, visited):
#     p, d, c = pdc
#     # print(p)

#     visited.add(pdc)
#     if memo.get(pdc) is not None:
#         return memo[pdc]

#     # time.sleep(0.2)
#     if p == 0:
#         return 0

#     tl = d * 1j
#     tr = -d * 1j
#     trials = [
#         (p + tl, tl, 0),
#         (p + tr, tr, 0),
#     ]
#     # next = []
#     if c < 3:
#         trials.append((p + d, d, c + 1))

#     next = [z for z in trials if z not in visited and grid.get(z[0])]
#     if len(next) == 0:
#         return 1e6

#     out = grid.get(p) + min(go(z, memo, visited) for z in next)
#     memo[pdc] = out
#     return out


# start = (Lx - 1) + 1j * (Ly - 1)

# memo = {}
# visited = set()
# print(go((Lx - 1 + 1j * (Ly - 1), -1, 0), memo, visited))

# DFS it is
cmin = 1e6

from copy import copy


p, d, c = (0, 1, 0)
visited = set()
current_heat = 0
while p != Lx - 1 + 1j * (Ly - 1):
    tl = d * 1j
    tr = -d * 1j
    try_next = [
        (p + tl, tl, 1),
        (p + tr, tr, 1),
    ]
    if c < 3:
        try_next.append(((p + d, d, c + 1)))

    filtered = [z for z in try_next if z not in visited and grid.get(z[0])]
    print(filtered)

    p, d, c = filtered.pop()
    visited.add((p, d, c, l))
    print(p)
    current_heat += grid.get(p)

print(current_heat)
