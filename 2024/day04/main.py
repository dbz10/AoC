def main(input_file="sample.txt"):
    rows = open(input_file).read().splitlines()
    grid = {
        x + y * 1j: ch for (y, row) in enumerate(rows) for (x, ch) in enumerate(row)
    }
    print(f"Part 1: {part1(grid)}")
    print(f"Part 2: {part2(grid)}")


def part1(input):
    x_positions = [k for k, v in input.items() if v == "X"]
    cnt = 0
    for xy in x_positions:
        for dir in [1, -1, 1j, -1j, 1 + 1j, 1 - 1j, -1 + 1j, -1 - 1j]:
            if scan_xmax(input, xy, dir) == ("M", "A", "S"):
                cnt += 1
    return cnt


def scan_xmax(grid, xy, dir):
    return grid.get(xy + dir), grid.get(xy + 2 * dir), grid.get(xy + 3 * dir)


def part2(input):
    a_positions = [k for k, v in input.items() if v == "A"]
    cnt = 0
    for xy in a_positions:
        (ur, ul, dr, dl) = scan_x_mas(input, xy)
        if ((ur == "M" and dl == "S") or (ur == "S" and dl == "M")) and (
            (ul == "M" and dr == "S") or (ul == "S" and dr == "M")
        ):
            cnt += 1
    return cnt


def scan_x_mas(grid, xy):
    return (
        grid.get(xy + 1 + 1j),
        grid.get(xy + 1 - 1j),
        grid.get(xy - 1 + 1j),
        grid.get(xy - 1 - 1j),
    )


if __name__ == "__main__":
    main()
