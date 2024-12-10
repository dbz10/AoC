class Topo:
    def __init__(self, topo: list[list[int]]):
        self.topo = topo
        self.lx = len(topo[0])
        self.ly = len(topo)

    def neighbors(self, x: int, y: int) -> list[tuple[int, int]]:
        n = [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]
        return [p for p in n if 0 <= p[0] < self.lx and 0 <= p[1] < self.ly]

    def height(self, x: int, y: int) -> int:
        return self.topo[y][x]

    def __repr__(self) -> str:
        return "\n".join(["".join(list(map(str, r))) for r in self.topo])


def main(input_file="sample.txt"):
    input = [[int(ch) for ch in row] for row in open(input_file).read().splitlines()]
    topo = Topo(input)
    print(f"Part 1: {part1(topo)}")
    print(f"Part 2: {part2(topo)}")


def part1(topo: Topo):
    lows = [
        (x, y) for x in range(topo.lx) for y in range(topo.ly) if topo.height(x, y) == 0
    ]
    return sum([trace(s, topo) for s in lows])


def part2(topo):
    lows = [
        (x, y) for x in range(topo.lx) for y in range(topo.ly) if topo.height(x, y) == 0
    ]
    return sum([trace(s, topo, p2=True) for s in lows])


def trace(start: tuple[int, int], topo: Topo, p2=False) -> int:
    """
    The number of trailheads reachable from start
    """
    front = [[start]]
    trails = set()
    while front:
        path = front.pop()
        (x, y) = path[-1]
        current_height = topo.height(x, y)
        neighbors = topo.neighbors(x, y)
        for nx, ny in neighbors:
            nh = topo.height(nx, ny)
            if nh == current_height + 1:
                if nh == 9:
                    if p2:
                        path_next = path + [(nx, ny)]
                        trails.add(tuple(path_next))
                    else:
                        trails.add((nx, ny))
                else:
                    path_next = path + [(nx, ny)]
                    front.append(path_next)
    return len(trails)


if __name__ == "__main__":
    main()
