L = 70
p1_pixels = 1024


def main(input_file="sample.txt"):
    input = [tuple(l.split(",")) for l in open(input_file).read().splitlines()]

    print(f"Part 1: {part1(input[:p1_pixels])}")
    print(f"Part 2: {part2(input)}")


def part1(input):
    grid = {int(x) + int(y) * 1j: "#" for x, y in input}
    s = ""
    for y in range(L + 1):
        for x in range(L + 1):
            if (x + 1j * y) in grid:
                s += "#"
            else:
                s += "."
        s += "\n"
    if len(input) % 1024 == 0:
        print(s)

    front = [(0, 0)]
    end = L + 1j * (L)
    distances = {0: 0}
    came_from = {}

    while front:
        (p, d) = front.pop(0)

        neighbors = [p + 1, p - 1, p + 1j, p - 1j]
        neighbors = [
            n
            for n in neighbors
            if 0 <= n.real <= L and 0 <= n.imag <= L and n not in grid
        ]
        d_new = d + 1
        for n in neighbors:
            if n not in distances.keys() or d_new < distances[n]:
                front.append((n, d_new))
                distances[n] = d_new
                came_from[n] = p

    return distances[end]


def part2(input):
    # i dunno, it's early and this way is obvious
    for npixels, blocker in enumerate(input):
        try:
            part1(input[: npixels + 1])
        except KeyError:
            return (npixels, blocker)


if __name__ == "__main__":
    main()
