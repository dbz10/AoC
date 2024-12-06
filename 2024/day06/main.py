def main(input_file="sample.txt"):
    input = open(input_file).read().splitlines()
    ly = len(input)
    lx = len(input[0])
    grid = {x + 1j * y: ch for (y, row) in enumerate(input) for x, ch in enumerate(row)}
    start = [k for k, v in grid.items() if v == "^"][0]
    grid[start] = "."

    print(f"Part 1: {part1(grid, start)}")
    print(f"Part 2: {part2(grid, start, lx, ly)}")


def part1(grid, pos):
    visited = walk_path(grid, pos)
    return len(visited)


def walk_path(grid, pos):
    dir = -1j
    visited = set()
    while grid.get(pos):
        visited.add(pos)
        if grid.get(pos + dir) == "#":
            dir *= 1j
        pos += dir
    return visited


def part2(grid, pos, lx, ly):
    initial_path = walk_path(grid, pos)

    loop_creations = [
        k for k in initial_path if k != pos and is_loop(grid | {k: "#"}, pos, lx, ly)
    ]

    return len(loop_creations)


def is_loop(grid, pos, lx, ly):
    dir = -1j
    visited = set()
    while grid.get(pos):
        if (pos, dir) in visited:
            # render(grid, visited, lx, ly)
            return True
        visited.add((pos, dir))
        if grid.get(pos + dir) == "#":
            dir *= 1j
        else:
            pos += dir
    return False


def render(grid, visited, lx, ly):
    for y in range(ly):
        for x in range(lx):
            p = x + 1j * y
            if p in visited:
                print("o", end="")
            else:
                print(grid[x + 1j * y], end="")
        print("")
    print("")


if __name__ == "__main__":
    main()
