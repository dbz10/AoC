from collections import Counter


def main(input_file="sample.txt"):
    input = open(input_file).read().splitlines()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")


def part1(input):
    unzipped = list(zip(*[z.split("   ") for z in input]))
    l = sorted(unzipped[0])
    r = sorted(unzipped[1])
    return sum(abs(int(x) - int(y)) for x, y in zip(l, r))


def part2(input):
    unzipped = list(zip(*[z.split("   ") for z in input]))
    cnt_r = Counter(unzipped[1])
    similarity_score = sum(int(i) * cnt_r.get(i, 0) for i in unzipped[0])
    return similarity_score


if __name__ == "__main__":
    main()
