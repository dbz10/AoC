input_file = "inputs/day16.txt"
import numpy as np
from concurrent.futures import ProcessPoolExecutor, as_completed
from tqdm import tqdm

grid = np.array([list(l) for l in open(input_file).read().split("\n")])

Lx, Ly = grid.shape


def check_bounds(p):
    return (0 <= p[0] < Lx) and (0 <= p[1] < Ly)


dx = np.array([0, 1])
dy = np.array([-1, 0])

dirs = {"r": dx, "u": dy, "d": -dy, "l": -dx}


def next(ray):
    trial_position = ray[0] + dirs[ray[1]]
    if not check_bounds(trial_position):
        return []
    cur_dir = ray[1]
    tile = grid[trial_position[0], trial_position[1]]
    if (
        (tile == ".")
        or (cur_dir in ("l", "r") and tile == "-")
        or (cur_dir in ("u", "d") and tile == "|")
    ):
        return [(trial_position, cur_dir)]

    if cur_dir in ("l", "r") and tile == "|":
        return [
            (trial_position, "u"),
            (trial_position, "d"),
        ]

    if cur_dir in ("u", "d") and tile == "-":
        return [
            (trial_position, "l"),
            (trial_position, "r"),
        ]

    if cur_dir == "r":
        if tile == "/":
            return [(trial_position, "u")]
        if tile == "\\":
            return [(trial_position, "d")]
    if cur_dir == "l":
        if tile == "/":
            return [(trial_position, "d")]
        if tile == "\\":
            return [(trial_position, "u")]
    if cur_dir == "u":
        if tile == "/":
            return [(trial_position, "r")]
        if tile == "\\":
            return [(trial_position, "l")]
    if cur_dir == "d":
        if tile == "/":
            return [(trial_position, "l")]
        if tile == "\\":
            return [(trial_position, "r")]


def go(start):
    energized = np.zeros(grid.shape)
    hist = {}
    while start:
        z = next(start.pop())
        for n in z:
            energized[n[0][0], n[0][1]] = 1
            if hist.get(str(n)) is None:
                hist[str(n)] = 1
                start.append(n)
    return int(energized.sum())


if __name__ == "__main__":
    print("Part 1:", go([(np.array([0, -1]), "r")]))

    with ProcessPoolExecutor() as e:
        futures = []
        for i in range(Lx):
            futures.append(e.submit(go, [(np.array([i, -1]), "r")]))
            futures.append(e.submit(go, [(np.array([i, Ly]), "l")]))
        for i in range(Ly):
            futures.append(e.submit(go, [(np.array([-1, i]), "d")]))
            futures.append(e.submit(go, [(np.array([Lx, i]), "u")]))

        futures = list(tqdm(as_completed(futures), total=len(futures)))

    res = [f.result() for f in futures]

    print("Part 2:", max(res))
