from queue import SimpleQueue


def main(input_file="sample.txt"):
    input = open(input_file).read().splitlines()
    grid = {x + 1j * y: ch for y, row in enumerate(input) for x, ch in enumerate(row)}
    print(f"Part 1: {part1(grid)}")
    print(f"Part 2: {part2(grid)}")


def part1(grid):
    start_dir = 1
    start_position = [k for k, v in grid.items() if v == "S"][0]
    end_position = [k for k, v in grid.items() if v == "E"][0]

    # dijkstra?
    lengths = {(k, d): 1e12 for k in grid.keys() for d in [1, -1, 1j, -1j]}
    lengths[start_position] = 0
    front = SimpleQueue()
    front.put((start_position, start_dir, 0))
    while not front.empty():
        p, d, l = front.get()
        # print(p, d, l)
        # if grid.get(p) == "E":
        #     break
        # check forward
        if grid.get(p + d) != "#":
            l_new = l + 1
            if l_new < lengths[(p + d, d)]:
                lengths[(p + d, d)] = l_new
                front.put((p + d, d, l_new))
        # check side to side
        for turn in [1j, -1j]:
            dnew = d * turn
            if grid.get(p + dnew) != "#":
                l_new = l + 1001
                if l_new < lengths[(p + dnew, dnew)]:
                    lengths[(p + dnew, dnew)] = l_new
                    front.put((p + dnew, dnew, l_new))

    return min(lengths[(end_position, d)] for d in [1, -1, 1j, -1j])


def part2(grid):
    start_dir = 1
    start_position = [k for k, v in grid.items() if v == "S"][0]
    end_position = [k for k, v in grid.items() if v == "E"][0]

    # dijkstra?
    lengths = {(k, d): 1e12 for k in grid.keys() for d in [1, -1, 1j, -1j]}
    front = SimpleQueue()
    front.put([(start_position, start_dir, 0)])
    finishing_paths = {}
    while not front.empty():
        path = front.get()
        p, d, l = path[-1]
        # print(p, d, l)
        if grid.get(p) == "E":
            finishing_paths[tuple(path)] = l
        # check forward
        if grid.get(p + d) != "#":
            l_new = l + 1
            if l_new <= lengths[(p + d, d)]:
                lengths[(p + d, d)] = l_new
                front.put(path + [(p + d, d, l_new)])
        # check side to side
        for turn in [1j, -1j]:
            dnew = d * turn
            if grid.get(p + dnew) != "#":
                l_new = l + 1001
                if l_new <= lengths[(p + dnew, dnew)]:
                    lengths[(p + dnew, dnew)] = l_new
                    front.put(path + [(p + dnew, dnew, l_new)])

    min_length = min(finishing_paths.values())
    best_paths = [k for k, v in finishing_paths.items() if v == min_length]
    tiles_on_best_paths = set([t[0] for path in best_paths for t in path])
    return len(tiles_on_best_paths)


if __name__ == "__main__":
    main()
