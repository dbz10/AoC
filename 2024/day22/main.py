from collections import defaultdict


def main(input_file="sample.txt"):
    input = list(map(int, open(input_file).read().splitlines()))
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")


def part1(input):
    res = {seed: evolve_n(seed, 2000)[-1] for seed in input}
    return sum(res.values())


def evolve(number: int) -> int:
    v1 = prune(mix(number, number * 64))
    v2 = prune(mix(v1, int(v1 / 32)))
    v3 = prune(mix(v2, v2 * 2048))
    return v3


def evolve_n(number: int, n: int) -> list[int]:
    res = [number]
    for _ in range(n):
        number = evolve(number)
        res.append(number)
    return res


def mix(x, y):
    return x ^ y


def prune(x):
    return x % 16777216


def part2(input):
    # idk lol
    # some heuristics get it to be not horrible in terms of time.
    all_monkeys_sequences = [[x % 10 for x in evolve_n(seed, 2000)] for seed in input]
    all_deltas = [[a - b for a, b in zip(v[1:], v)] for v in all_monkeys_sequences]
    banana_prices = [z[1:] for z in all_monkeys_sequences]
    price_getting_sequence = defaultdict(dict)
    for monkey, delta in enumerate(all_deltas):
        for idx, seq in enumerate(zip(delta, delta[1:], delta[2:], delta[3:]), start=3):
            if seq not in price_getting_sequence[monkey]:
                price_getting_sequence[monkey][seq] = banana_prices[monkey][idx]
            if len(price_getting_sequence[monkey]) == 10:
                continue

    max_bananas = -1
    opt_seq = None
    possibly_positive_sequences = set(
        k for k, v in price_getting_sequence[0].items() if v == 9
    )
    for i in range(1, len(input)):
        possibly_positive_sequences = possibly_positive_sequences | set(
            k for k, v in price_getting_sequence[i].items() if v == 9
        )

    for s in possibly_positive_sequences:
        bananas = sum(
            [price_getting_sequence[monkey].get(s, 0) for monkey in range(len(input))]
        )
        if bananas > max_bananas:
            max_bananas = bananas
            opt_seq = s
    return max_bananas


if __name__ == "__main__":
    main()
