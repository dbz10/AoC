from __future__ import annotations
from dataclasses import dataclass
import re


@dataclass
class Point:
    x: int
    y: int

    def __add__(self, other: Point) -> Point:
        Point(self.x + other.x, self.y + other.y)


@dataclass
class Game:
    A: Point
    B: Point
    target: Point

    def solve(self) -> Point | None:
        # (ax bx) (d1) = (tx)
        # (ay by) (d2)   (ty)

        det = self.A.x * self.B.y - self.B.x * self.A.y
        if det == 0:
            # underspecified.
            # actually this never happened
            if self.target.x / self.b.x <= 100:
                return self.target.x / self.b.x
            else:
                db = 100
                da = (self.target.x - self.b.x * 100) / self.a.x

        else:
            # fp precision :(
            da = round(
                (1 / det) * (self.B.y * self.target.x - self.B.x * self.target.y), 2
            )
            db = round(
                (1 / det) * (-self.A.y * self.target.x + self.A.x * self.target.y), 2
            )

        if da % 1 != 0 or db % 1 != 0:
            return None

        return int(3 * da + db)


pattern_bx = "X\+(\d*)"
pattern_by = "Y\+(\d*)"
pattern_tx = "X=(\d*)"
pattern_ty = "Y=(\d*)"


def main(input_file="sample.txt"):
    input = [l.splitlines() for l in open(input_file).read().split("\n\n")]
    games = []
    for [button_a, button_b, target] in input:
        ax = int(re.findall(pattern_bx, button_a)[0])
        ay = int(re.findall(pattern_by, button_a)[0])

        bx = int(re.findall(pattern_bx, button_b)[0])
        by = int(re.findall(pattern_by, button_b)[0])

        tx = int(re.findall(pattern_tx, target)[0])
        ty = int(re.findall(pattern_ty, target)[0])

        games.append(Game(Point(ax, ay), Point(bx, by), Point(tx, ty)))

    print(f"Part 1: {part1(games)}")
    print(f"Part 2: {part2(games)}")


def part1(games):
    return sum(tokens for game in games if (tokens := game.solve()) is not None)


def part2(games: list[Game]):
    for game in games:
        game.target.x += 10000000000000
        game.target.y += 10000000000000
    return sum(tokens for game in games if (tokens := game.solve()) is not None)


if __name__ == "__main__":
    main()
