def main(input_file="sample.txt"):
    input = open(input_file).read().split("\n\n")

    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")


def part1(input):
    # pure evil!

    exec(input[0].replace(":", "="))

    while True:
        all_set = True
        for eqrow in (
            input[1]
            .replace("XOR", "^")
            .replace("AND", "&")
            .replace("OR", "|")
            .splitlines()
        ):
            [eq, var] = eqrow.split(" -> ")
            try:
                eval(var)
            except NameError:
                try:
                    exec(f"{var} = {eq}".strip())
                except NameError:
                    all_set = False

        if all_set:
            break

    zs = sorted(
        [
            var
            for r in input[1].splitlines()
            if (var := r.split("-> ")[-1].strip()).startswith("z")
        ],
        reverse=True,
    )
    s = ""
    for z in zs:
        s += str(eval(z))

    return int(s, base=2)


def part2(input):
    return


if __name__ == "__main__":
    main()
