from functools import cache


def main(input_file="sample.txt"):
    input = open(input_file).read().split("\n\n")
    towels = tuple(input[0].strip().split(", "))
    designs = input[1].splitlines()
    print(f"Part 1: {part1(towels, designs)}")
    print(f"Part 2: {part2(towels, designs)}")


def part1(towels, designs):
    return len([d for d in designs if designable(d, towels)])


@cache
def designable(design: str, towels: tuple[str]):
    if design == "":
        return 1
    into = 0
    for towel in towels:
        if design.endswith(towel):
            subdesign = design[: -len(towel)]
            into += designable(subdesign, towels)
    return into


def part2(towels, designs):
    return sum([designable(d, towels) for d in designs])


if __name__ == "__main__":
    main()
