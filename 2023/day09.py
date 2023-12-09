import numpy as np

seqs = [np.array(list(map(int, l.split()))) for l in open("inputs/day09.txt")]

acc = 0
for seq in seqs:
    lasts = [seq[-1]]
    while any(d != 0 for d in seq):
        seq = np.diff(seq)
        lasts.append(seq[-1])

    ups = [0]
    rev = lasts[::-1]
    for i in range(1, len(lasts)):
        ups.append(ups[-1] + rev[i])

    acc += ups[-1]

print("Part 1:", acc)


acc = 0
for seq in seqs:
    firsts = [seq[0]]
    while any(d != 0 for d in seq):
        seq = np.diff(seq)
        firsts.append(seq[0])

    ups = [0]
    rev = firsts[::-1]
    for i in range(1, len(firsts)):
        ups.append(rev[i] - ups[-1])

    acc += ups[-1]

print("Part 2:", acc)
