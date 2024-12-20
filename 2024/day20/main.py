from queue import SimpleQueue
from dataclasses import dataclass
from copy import copy
from collections import defaultdict

MIN_IMPROVEMENT = 100


def main(input_file="sample.txt"):
    input = open(input_file).read().splitlines()
    grid = {x + 1j * y: ch for y, row in enumerate(input) for x, ch in enumerate(row)}

    print(f"Part 1: {part1(grid)}")
    print(f"Part 2: {part2(grid)}")


def part1(grid):
    start = [k for k, v in grid.items() if v == "S"][0]
    end = [k for k, v in grid.items() if v == "E"][0]
    noskip_map = lengths_no_skip(grid, start, end)
    original_length = noskip_map[start]
    possible_lengths_with_skips = skipfind(
        grid, start, end, early_stopping_mapping=noskip_map, skip_length=2
    )
    improvements = [
        original_length - v
        for v in possible_lengths_with_skips.values()
        if (original_length - v) >= MIN_IMPROVEMENT
    ]

    return len(improvements)


@dataclass
class State:
    """
    In the end i moved the skipping logic outside, so there's not much reason to have this class
    but i left it anyways
    """

    position: complex
    skip_available: bool = True
    skip_active: bool = False
    skips_remaining: int = 2
    cheat_start: complex | None = None
    cheat_end: complex | None = None

    # def activate_skip(self):
    #     if self.skip_available:
    #         self.skip_active = True
    #         self.skip_available = False
    #         self.cheat_start = self.position

    def move(self, d: complex):
        # if self.skip_active:
        #     self.skips_remaining -= 1
        #     if self.skips_remaining == 0:
        #         self.skip_active = 0
        #         self.cheat_end = self.position + d

        self.position += d


def lengths_no_skip(grid, start, end):
    # first just get the distance to the end from each open space
    front: SimpleQueue[tuple[complex, int]] = SimpleQueue()
    front.put((start, 0))
    distances = {start: 0}
    while not front.empty():
        p, l = front.get()
        distances[p] = l
        if p == end:
            break
        else:
            for d in [1, -1, 1j, -1j]:
                # First just check walking without using a skip
                if p + d not in distances.keys() and grid.get(p + d) in [".", "E"]:
                    front.put((p + d, l + 1))

    l_end = distances[end]
    distance_from_end = {position: l_end - l for position, l in distances.items()}

    return distance_from_end


def skipfind(
    grid,
    start,
    end,
    early_stopping_mapping: dict[complex, int],
    skip_length: int = 2,
):
    """
    We already found the distance to the end from each open patch.
    So I think we just need to walk along open patches and consider jumping to any open patch within
    reach that's closer to the end?
    """
    original_length = early_stopping_mapping[start]

    finishing_paths = {}

    for p, distance_to_end in early_stopping_mapping.items():
        ds = [
            dx + 1j * dy
            for dx in range(-skip_length, skip_length + 1)
            for dy in range(-skip_length, skip_length + 1)
            if p + dx + 1j * dy in early_stopping_mapping
            and abs(dx) + abs(dy) <= skip_length
        ]
        for d in ds:
            traveled_so_far = original_length - distance_to_end
            distance_from_jump_point_to_end = early_stopping_mapping[p + d]
            jump_length = abs(d.real) + abs(d.imag)
            full_distance_if_jump = (
                traveled_so_far + distance_from_jump_point_to_end + jump_length
            )

            if full_distance_if_jump <= original_length - MIN_IMPROVEMENT:
                finishing_paths[(p, p + d)] = (
                    original_length - distance_to_end
                ) + early_stopping_mapping[p + d]

    return finishing_paths


def part2(grid):
    start = [k for k, v in grid.items() if v == "S"][0]
    end = [k for k, v in grid.items() if v == "E"][0]
    noskip_map = lengths_no_skip(grid, start, end)
    original_length = noskip_map[start]

    possible_lengths_with_skips = skipfind(
        grid,
        start,
        end,
        early_stopping_mapping=noskip_map,
        skip_length=20,
    )
    improvements = [
        original_length - v
        for v in possible_lengths_with_skips.values()
        if (original_length - v) >= MIN_IMPROVEMENT
    ]

    return len(improvements)


if __name__ == "__main__":
    main()
