WALL = "#"
BOX = "O"
ROBOT = "@"
EMPTY = "."
DIRECTIONS = {"v": 1j, ">": 1, "<": -1, "^": -1j}

EXPANSION_RULES = {"#": "##", "@": "@.", ".": "..", "O": "[]"}

WIDEBOX = ["[", "]"]


def main(input_file="sample.txt"):
    input = open(input_file).read().split("\n\n")

    map_input = input[0].splitlines()
    map = {
        x + 1j * y: ch for y, row in enumerate(map_input) for x, ch in enumerate(row)
    }
    directions = input[1].replace("\n", "")
    print(f"Part 1: {part1(map, directions)}")

    expanded_input = [
        [k for ch in row for k in EXPANSION_RULES[ch]] for row in map_input
    ]
    map = {
        x + 1j * y: ch
        for y, row in enumerate(expanded_input)
        for x, ch in enumerate(row)
    }

    print(f"Part 2: {part2(map, directions)}")


def part1(map, directions):
    pos = [k for k, v in map.items() if v == ROBOT][0]
    for dir in directions:
        d = DIRECTIONS[dir]
        update = {}

        if map[pos + d] == EMPTY:
            update[pos + d] = ROBOT
            update[pos] = EMPTY
        elif map[pos + d] == BOX:
            l = 1
            while map[pos + d * l] == BOX:
                l += 1
            if map[pos + d * l] == EMPTY:
                update[pos + d * l] = BOX
                update[pos + d] = ROBOT
                update[pos] = EMPTY

        if len(update) > 0:
            map = map | update
            pos += d
        # render(map)
    gps_sum = sum([p.real + 100 * p.imag for p, c in map.items() if c == BOX])

    return gps_sum


def render(map):
    lx = int(max(p.real for p in map.keys()))
    ly = int(max(p.imag for p in map.keys()))
    rendered = ""
    for y in range(ly + 1):
        for x in range(lx + 1):
            rendered += map[x + 1j * y]
        rendered += "\n"
    print(rendered)


def collision_rules(map, start, direction):
    # jeez
    return "thanks a lot eric"


def contains_box(map, point):
    return map[point] in WIDEBOX


def part2(input):
    pos = [k for k, v in map.items() if v == ROBOT][0]
    for dir in directions:
        d = DIRECTIONS[dir]
        update = {}

        if map[pos + d] == EMPTY:
            update[pos + d] = ROBOT
            update[pos] = EMPTY
        elif map[pos + d] == BOX:
            l = 1
            while map[pos + d * l] == BOX:
                l += 1
            if map[pos + d * l] == EMPTY:
                update[pos + d * l] = BOX
                update[pos + d] = ROBOT
                update[pos] = EMPTY

        if len(update) > 0:
            map = map | update
            pos += d
        # render(map)
    gps_sum = sum([p.real + 100 * p.imag for p, c in map.items() if c == BOX])

    return gps_sum


if __name__ == "__main__":
    main()
