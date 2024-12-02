import numpy as np


def main(input_file="sample.txt"):
    input = [
        np.array(list(map(int, z.split(" "))))
        for z in open(input_file).read().splitlines()
    ]
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")


def part1(input):
    return sum(condition1(row) for row in input)


def part2(input):
    return sum(condition2(row) for row in input)


def condition1(row) -> bool:
    deltas = row[1:] - row[:-1]
    return (
        ((deltas > 0).all() or (deltas < 0).all())
        and (1 <= np.abs(deltas)).all()
        and (np.abs(deltas) <= 3).all()
    )


def condition2(row) -> bool:
    if condition1(row):
        return True
    for removal in range(len(row)):
        depleted = np.concatenate([row[:removal], row[removal + 1 :]])
        if condition1(depleted):
            return True
    return False


if __name__ == "__main__":
    main()
