input_file = "inputs/day10.txt"

grid = {}
for y, line in enumerate(open(input_file)):
    for x, char in enumerate(line):
        if char == "\n":
            continue
        grid[(x, y)] = char
        if char == "S":
            start = (x, y)

next_ok = {
    # (dx, dy) -> allowable target tiles
    (0, 1): ["|", "L", "J", "S"],  # down is +
    (0, -1): ["|", "F", "7", "S"],
    (1, 0): ["-", "7", "J", "S"],
    (-1, 0): ["-", "L", "F", "S"],
}


cur = start
prev = cur
steps = 0
path = [cur]

area = 0
while (cur != start) or (steps == 0):
    t = cur
    cur = next(
        step
        for dx, dy in next_ok.keys()
        if grid.get(step := (cur[0] + dx, cur[1] + dy)) in next_ok[(dx, dy)]
        and grid[cur] in next_ok[(-dx, -dy)]
        and step != prev
    )

    prev = t
    steps += 1
    path.append(cur)


print("Part 1:", int(steps / 2))

# Part 2: i dunno, brute force check for every grid position whether inside or outside?

acc = 0
spots = []

# this feels kind of brittle and unsatisfying, but a star is a star
crossings_contribution = {
    "|": 1,
    "L": -1 / 2,
    "J": 1 / 2,
    "F": 1 / 2,
    "7": -1 / 2,
    "S": -1 / 2,
    "-": 0,
}

from tqdm import tqdm

# an extra loop could have been skipped here, but it's already done
for xy in tqdm(grid.keys()):
    xyo = xy
    if xy in path:
        continue
    crossings = 0
    while xy[0] >= -1:
        xy = (xy[0] - 1, xy[1])
        if xy in path:
            crossings += crossings_contribution[grid[xy]]

    if crossings % 2:
        spots.append(xyo)
        acc += 1


print("Part 2:", acc)
