import numpy as np

input_file = "inputs/day11.txt"

arrs = []
for line in open(input_file):
    arrs.append(np.array([1 if ch == "#" else 0 for ch in line if ch != "\n"]))

grid = np.vstack(arrs)
Lx, Ly = grid.shape


# i = 0
# while i < Lx:
#     if all(grid[i, :] == 0):
#         grid = np.vstack([grid[:i], np.zeros(Ly), grid[i:]])
#         i += 1
#         Lx = grid.shape[0]
#     i += 1
# i = 0
# while i < Ly:
#     if all(grid[:, i] == 0):
#         grid = np.hstack([grid[:, :i], np.zeros(Lx).reshape(-1, 1), grid[:, i:]])
#         i += 1
#         Ly = grid.shape[1]
#     i += 1

horizontal_expanders = [i for i in range(Lx) if all(grid[i, :] == 0)]
vertical_expanders = [i for i in range(Ly) if all(grid[:, i] == 0)]


galaxies = []
for x in range(Lx):
    for y in range(Ly):
        if grid[x, y] == 1:
            galaxies.append((x, y))

acc_1 = 0
acc_2 = 0
expansion_factor = 1_000_000 - 1
for i, g1 in enumerate(galaxies):
    for g2 in galaxies[i + 1 :]:
        n_expanders = len(
            [
                1
                for e in horizontal_expanders
                if min(g1[0], g2[0]) < e < max(g1[0], g2[0])
            ]
            + [
                1
                for e in vertical_expanders
                if min(g1[1], g2[1]) < e < max(g1[1], g2[1])
            ]
        )
        acc_1 += abs(g1[0] - g2[0]) + abs(g1[1] - g2[1]) + n_expanders
        acc_2 += (
            abs(g1[0] - g2[0]) + abs(g1[1] - g2[1]) + expansion_factor * n_expanders
        )


print("Part 1:", acc_1)
print("Part 2:", acc_2)
