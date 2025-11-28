def main(input_file="sample.txt"):
    input = open(input_file).read().split("\n\n")
    locks = []
    keys = []

    for key_or_lock in input:
        rows = key_or_lock.splitlines()
        thing = [-1] * len(rows[0])
        for row in rows:
            for idx, ch in enumerate(row):
                if ch == "#":
                    thing[idx] += 1
        if rows[0][0] == "#":
            locks.append(thing)
        else:
            keys.append(thing)

    print(f"Part 1: {part1(locks, keys)}")
    print(f"Part 2: {part2(input)}")


AVAILABLE_SPACE = 5


def part1(locks, keys):
    fits = 0
    for lock in locks:
        for key in keys:
            merged = [a + b for a, b in zip(lock, key)]
            if all(z <= AVAILABLE_SPACE for z in merged):
                fits += 1

    return fits


def part2(input):
    return


if __name__ == "__main__":
    main()
