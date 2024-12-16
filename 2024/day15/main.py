from __future__ import annotations
from dataclasses import dataclass
from queue import SimpleQueue


WALL = "#"
BOX = "O"
ROBOT = "@"
EMPTY = "."
LEFTBOX = "["
RIGHTBOX = "]"
DIRECTIONS = {"v": 1j, ">": 1, "<": -1, "^": -1j}

EXPANSION_RULES = {"#": "##", "@": "@.", ".": "..", "O": "[]"}

WIDEBOX = ["[", "]"]


def main(input_file="day15/sample.txt"):
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
    gps_sum = int(sum([p.real + 100 * p.imag for p, c in map.items() if c == BOX]))

    return gps_sum


def render(map):
    lx = int(max(p.real for p in map.keys()))
    ly = int(max(p.imag for p in map.keys()))
    rendered = ""
    for y in range(ly + 1):
        for x in range(lx + 1):
            rendered += str(map[x + 1j * y])
        rendered += "\n"
    print(rendered)


def collision_rules(
    map: dict[complex, str],
    p: complex,
    d: complex,
) -> tuple[dict[complex, str], list[complex]] | None:
    if map[p + d] == WALL:
        return None

    symbol = map[p]
    update = {p: EMPTY, p + d: map[p]}
    if symbol == LEFTBOX:
        if d.imag != 0:
            check_next = list(set([p + 1, p + d]))
        else:
            check_next = [p + d]
    elif symbol == RIGHTBOX:
        if d.imag != 0:
            check_next = list(set([p - 1, p + d]))
        else:
            check_next = [p + d]
    elif symbol == ROBOT:
        check_next = [p + d]
    else:
        check_next = []

    check_next = [c for c in check_next if map[c] != EMPTY]
    return update, check_next


def part2(map, directions):
    pos = [k for k, v in map.items() if v == ROBOT][0]

    for i, dir in enumerate(directions):
        d = DIRECTIONS[dir]

        to_check = SimpleQueue()
        to_check.put(pos)
        visited = set()
        update_rules = []
        blocked = False
        while not to_check.empty():
            check_me = to_check.get()
            move = collision_rules(map, check_me, d)
            if move is None:
                blocked = True
                break
            (update_rule, check_next) = move
            update_rules.append(update_rule)
            for c in [c for c in check_next if c not in visited]:
                to_check.put(c)
                visited.add(c)
        if blocked:
            continue

        # put all the updates together in the reverse order
        update = {}
        for update_rule in update_rules[::-1]:
            update = update | update_rule

        map = map | update
        pos += d

    render(map)
    gps_sum = int(sum([p.real + 100 * p.imag for p, c in map.items() if c == LEFTBOX]))

    return gps_sum


if __name__ == "__main__":
    main()
