import numpy as np

input_file = "inputs/day14.txt"

data = np.array([list(l) for l in open(input_file).read().split("\n")])


def process_line(line):
    block = -1
    for i in range(len(line)):
        if line[i] == "#":
            block = i
        if line[i] == "O":
            line[block + 1] = "O"
            if i != block + 1:
                line[i] = "."
            block += 1
    return line


def push(data):
    return np.vstack([process_line(l) for l in np.copy(data)])


def push_north(data):
    return push(data.transpose()).transpose()


def push_west(data):
    return push(data)


def push_east(data):
    return push(data[:, ::-1])[:, ::-1]


def push_south(data):
    return push(data[::-1].transpose()).transpose()[::-1]


def score_data(d):
    return sum([sum(np.arange(1, L + 1)[::-1][line == "O"]) for line in d.transpose()])


def stringify(d):
    return "".join("".join(i) for i in d)


acc = 0
L = data.shape[0]
for line in np.copy(data.transpose()):
    acc += np.sum(np.arange(1, L + 1)[::-1][process_line(line) == "O"])


print("Part 1:", acc)

first = 0
round = 0
cur = np.copy(data)
cycle = []
scores = []
while round < 10000:
    round += 1
    next = push_north(cur)
    next = push_west(next)
    next = push_south(next)
    next = push_east(next)
    cur = np.copy(next)

    sc = stringify(cur)
    if sc in cycle:
        first = cycle.index(sc)
        break
    cycle.append(sc)
    scores.append(score_data(cur))


cycle_length = round - first - 1
cycle_start = first


final_state = (1000000000 - cycle_start - 1) % cycle_length + cycle_start
print("Part 2:", scores[final_state])
