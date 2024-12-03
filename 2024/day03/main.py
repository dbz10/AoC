import re


def main(input_file="sample.txt"):
    input = open(input_file).read()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")


def part1(input):
    reg = re.compile("mul\([0-9]+,[0-9]+\)")
    groups = reg.findall(input)
    nums = [
        list(map(int, g.replace("mul(", "").replace(")", "").split(",")))
        for g in groups
    ]
    return sum(n[0] * n[1] for n in nums)


def part2(input):
    reg = re.compile("mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)")
    groups = reg.findall(input)
    ops = []
    active = True
    for g in groups:
        if g == "do()":
            active = True
        elif g == "don't()":
            active = False
        elif active:
            ops.append(g)
    nums = [
        list(map(int, g.replace("mul(", "").replace(")", "").split(","))) for g in ops
    ]
    return sum(n[0] * n[1] for n in nums)


if __name__ == "__main__":
    main()
