from collections import defaultdict
import math


def main(input_file="sample.txt"):
    input = list(map(list, open(input_file).read().splitlines()))
    lx = len(input[0])
    ly = len(input)
    antenna = defaultdict(list)
    for y, row in enumerate(input):
        for x, ch in enumerate(row):
            if ch != ".":
                antenna[ch].append((x, y))

    print(f"Part 1: {part1(antenna, lx, ly, False)}")
    print(f"Part 2: {part1(antenna, lx, ly, True)}")


def part1(antenna, lx, ly, p2=False):
    antinodes_found = set()
    if p2:  # lol
        # if its stupid and it works, then ... well it's still stupid but it works
        span = range(-100, 100)
    else:
        span = [1]

    for like_nodes in antenna.values():
        for i in range(len(like_nodes)):
            for j in range(i + 1, len(like_nodes)):
                p1 = like_nodes[i]
                p2 = like_nodes[j]
                dx, dy = p2[0] - p1[0], p2[1] - p1[1]
                r = math.gcd(dx, dy)
                dx, dy = dx / r, dy / r

                for n in span:  # lol
                    a1 = (p1[0] + n * 2 * r * dx, p1[1] + n * 2 * r * dy)
                    a2 = (p2[0] - n * 2 * r * dx, p2[1] - n * 2 * r * dy)
                    for a in [a1, a2]:
                        if 0 <= a[0] < lx and 0 <= a[1] < ly:
                            antinodes_found.add((int(a[0]), int(a[1])))
    return len(antinodes_found)


if __name__ == "__main__":
    main()
