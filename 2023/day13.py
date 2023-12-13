import numpy as np

input_file = "inputs/day13.txt"


data = [l.strip() for l in open(input_file).read().split("\n\n")]


arrs = [
    np.array(
        list(
            map(
                lambda x: list(map(int, x)),
                map(list, d.replace("#", "1").replace(".", "0").split("\n")),
            )
        )
    )
    for d in data
]


def find_pivot_contribution(arr, skip=None):
    Lx, Ly = arr.shape
    for i in range(1, Lx):
        p = min(i, Lx - i)
        if i < Lx / 2:
            left = arr[:p]
            right = arr[p : 2 * p][::-1]
        else:
            left = arr[::-1][:p]
            right = arr[::-1][p : 2 * p][::-1]

        if np.array_equal(left, right) and ("h", i) != skip:
            return ("h", i, 100 * i)

    for i in range(1, Ly):
        p = min(i, Ly - i)
        if i < Ly / 2:
            left = arr[:, :p]
            right = arr[:, p : 2 * p][:, ::-1]
        else:
            left = arr[:, ::-1][:, :p]
            right = arr[:, ::-1][:, p : 2 * p][:, ::-1]

        if np.array_equal(left, right) and ("v", i) != skip:
            return ("v", i, i)


acc = 0
acc_2 = 0
for d in arrs:
    original_sol = find_pivot_contribution(d)
    acc += original_sol[2]
    sh = d.shape
    for i in range(d.size):
        dt = np.copy(d)
        idx = np.unravel_index(i, sh)
        dt[idx] = 1 - dt[idx]
        tentative = find_pivot_contribution(dt, skip=original_sol[:2])
        if tentative and tentative != original_sol:
            acc_2 += tentative[2]
            break

print(acc)
print(acc_2)
