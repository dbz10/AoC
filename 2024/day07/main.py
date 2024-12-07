def main(input_file="sample.txt"):
    input = open(input_file).read().splitlines()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")


def part1(input):
    cnt = 0
    for line in input:
        s = line.split(":")
        test_value = int(s[0])
        components = list(map(int, s[1].split()))
        if check(test_value, components):
            cnt += test_value
    return cnt


def check(rem, cs):
    if rem < 0:
        return False
    if len(cs) == 0:
        return rem == 0
    return check(rem - cs[-1], cs[:-1]) or check(rem / cs[-1], cs[:-1])


def check2(rem, cs):
    if rem < 0:
        return False
    if len(cs) == 0:
        return rem == 0
    head, tail = cs[:-1], cs[-1]  # not the typical def of head and tail...
    possible_mods = [check2(int(rem - tail), head)]
    if rem / tail % 1 == 0:
        possible_mods.append(check2(int(rem / tail), head))
    if str(rem).endswith(str(tail)):
        if str(rem) == str(tail):
            return True
        else:
            deconcatenated = int(str(rem)[: -len(str(tail))])
            possible_mods.append(check2(deconcatenated, head))
    return any(possible_mods)


def part2(input):
    cnt = 0
    for line in input:
        s = line.split(":")
        test_value = int(s[0])
        components = list(map(int, s[1].split()))
        if check2(test_value, components):
            cnt += test_value
    return cnt


if __name__ == "__main__":
    main()
