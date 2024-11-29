input_file = "inputs/day21.txt"

grid = {
    x + y * 1j: s
    for (y, line) in enumerate(open(input_file).read().split("\n"))
    for x, s in enumerate(line)
}

Lx = int(max(c.real for c in grid.keys()) + 1)
Ly = int(max(c.imag for c in grid.keys()) + 1)


start = {v: k for k, v in grid.items() if v == "S"}["S"]
grid[start] = "."


n_steps = 6

front = [start]
for i in range(n_steps):
    nexts = []
    while front:
        p = front.pop()
        # visited.add(p)
        candidates = [p + 1j, p - 1j, p + 1, p - 1]
        candidates = [c for c in candidates if grid.get(c) == "."]
        nexts.extend(candidates)
    front = set(nexts)

print(len(front))


n_steps = 500

front = [start]
sf = []
from tqdm import tqdm

for i in tqdm(range(n_steps)):
    nexts = []
    sf.append(len(front))
    while front:
        p = front.pop()
        # visited.add(p)
        candidates = [p + 1j, p - 1j, p + 1, p - 1]
        cwrapped = [(c.real % Lx) + 1j * (c.imag % Ly) for c in candidates]
        # print(candidates)
        candidates = [c for c, cw in zip(candidates, cwrapped) if grid.get(cw) == "."]
        nexts.extend(candidates)
    front = set(nexts)

print(len(front))
import matplotlib.pyplot as plt

fig, ax = plt.subplots(figsize=(11, 7))
ax.plot(range(n_steps), sf)
plt.show()

import pandas as pd

pd.DataFrame({"n": range(n_steps), "y": sf}).to_csv("data.csv", index=False)
