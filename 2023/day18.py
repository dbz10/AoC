input_file = "inputs/day18.txt"

dirs = {"R": 1, "L": -1, "U": 1j, "D": -1j}
cv = {"0": "R", "1": "D", "2": "L", "3": "U"}


def convert_color(s):
    s = s.replace("(", "").replace(")", "").replace("#", "")
    return (dirs[cv[s[5]]], int(s[:5], 16))


instructions = [
    (dirs[d], int(l), convert_color(c))
    for (d, l, c) in map(lambda x: x.split(), open(input_file).read().split("\n"))
]

area = 0
circumference = 0
p = 0 + 0j
for d, l, _ in instructions:
    step = d * l
    area += step.real * p.imag  # y * dx
    circumference += l
    p += step

# pick's theorem
interior_points = area + 1 - circumference / 2
print("Part 1:", circumference + interior_points)


area = 0
circumference = 0
p = 0 + 0j
for _, _, (d, l) in instructions:
    step = d * l
    area += step.real * p.imag  # y * dx
    circumference += l
    p += step

interior_points = area + 1 - circumference / 2
print("Part 2:", circumference + interior_points)
