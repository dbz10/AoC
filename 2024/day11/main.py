from collections import Counter, defaultdict


def main(input_file="sample.txt"):
    input = list(map(int, open(input_file).read().split()))
    print(f"Part 1: {stonecounter(input, 25)}")
    print(f"Part 2: {stonecounter(input, 75)}")


def blink(i: int) -> list[int]:
    if i == 0:
        return [1]
    # ok yeah definitely never using the walrus operator again
    if ((l := len((s := str(i)))) % 2) == 0:
        left, right = s[: l // 2], s[l // 2 :]
        return [int(left), int(right)]
    return [i * 2024]


def stonecounter(input, n_rounds):
    stonecounter = Counter(input)
    for _ in range(n_rounds):
        update = defaultdict(int)
        for stone_number, amount in stonecounter.items():
            update[stone_number] -= amount
            for next in blink(stone_number):
                update[next] += amount
        stonecounter.update(update)
    return sum(stonecounter.values())


if __name__ == "__main__":
    main()
