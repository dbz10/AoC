from __future__ import annotations
from dataclasses import dataclass


@dataclass
class GardenPatch:
    label: str
    cluster: int

    def listen(self, other: GardenPatch):
        if other.label == self.label and other.cluster <= self.cluster:
            self.cluster = other.cluster


def main(input_file="day12/sample.txt"):
    input = list(map(list, open(input_file).read().splitlines()))
    garden = {}

    # weirdest floodfill implementation?
    for y, row in enumerate(input):
        for x, ch in enumerate(row):
            garden[(x, y)] = ch

    cluster = {}
    cluster_id = 0
    unvisited = set(garden.keys())
    while unvisited:
        (x, y) = unvisited.pop()
        front = [(x, y)]
        cluster[(x, y)] = cluster_id
        cluster_id += 1
        while front:
            (x, y) = front.pop()
            for xp, yp in [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]:
                if (xp, yp) not in cluster and garden.get((xp, yp)) == garden[(x, y)]:
                    cluster[(xp, yp)] = cluster[(x, y)]
                    front.append((xp, yp))
                    unvisited.remove((xp, yp))

    garden = {k: GardenPatch(garden[k], cluster[k]) for k in garden.keys()}
    print(f"Part 1: {part1(garden)}")
    print(f"Part 2: {part2(garden)}")


def part1(garden: dict[tuple[int, int], GardenPatch]):
    # maybe not efficient but straightforward
    clusters = set([p.cluster for p in garden.values()])
    areas = {c: 0 for c in clusters}
    perimeters = {c: 0 for c in clusters}

    for cluster in clusters:
        patch_locations = [k for k, v in garden.items() if v.cluster == cluster]
        for x, y in patch_locations:
            areas[cluster] += 1
            for xp, yp in [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]:
                if (xp, yp) in garden:
                    if garden[(xp, yp)].label != garden[(x, y)].label:
                        perimeters[cluster] += 1
                else:
                    perimeters[cluster] += 1

    return sum(areas[k] * perimeters[k] for k in clusters)


def part2(garden):
    clusters = set([p.cluster for p in garden.values()])
    cluster_labels = {v.cluster: v.label for v in garden.values()}
    areas = {c: 0 for c in clusters}
    sides = {c: 0 for c in clusters}

    # Ok, corner counting, because apparently I'm too stupid to do wall crawling
    for cluster in clusters:
        patch_locations = [k for k, v in garden.items() if v.cluster == cluster]
        areas[cluster] += len(patch_locations)

        if len(patch_locations) in [1, 2]:
            sides[cluster] = 4
            continue

        for x, y in patch_locations:
            corners_detected = 0
            shape = local_shape(garden, x, y)

            corners_detected += scansum(shape, ALL_CORNERS)

            if scan(shape, ENDPIECES):
                corners_detected += 2
            sides[cluster] += corners_detected

    return sum(areas[k] * sides[k] for k in clusters)


def neighbor_candidates(x, y):
    return [
        (x + dx, y + dy)
        for dx in [-1, 0, 1]
        for dy in [-1, 0, 1]
        if (x + dx, y + dy) != (x, y)
    ]


def local_shape(garden, x, y):
    return [
        [
            int(
                garden.get((x + dx, y + dy)) is not None
                and garden.get((x + dx, y + dy)).cluster == garden[(x, y)].cluster
            )
            for dx in [-1, 0, 1]
        ]
        for dy in [-1, 0, 1]
    ]


def shape_comparison(shape, template):
    for y, r in enumerate(shape):
        for x, v in enumerate(r):
            if template[y][x] != "*" and template[y][x] != shape[y][x]:
                return False
    return True


def scan(shape, template_group):
    return any(shape_comparison(shape, template) for template in template_group)


def scansum(shape, template_group):
    return sum(shape_comparison(shape, template) for template in template_group)


def rotate_shape(shape):
    return [[shape[x][2 - y] for x in range(3)] for y in range(3)]


def rotaten(shape, times=1):
    for _ in range(times):
        shape = rotate_shape(shape)
    return shape


# Manual enumeration of shape templates because I'm too stupid to do anything better
ICBASE = [["*", 1, 0], ["*", 1, 1], ["*", "*", "*"]]
INSIDE_CORNERS = [
    ICBASE,
    rotaten(ICBASE, 1),
    rotaten(ICBASE, 2),
    rotaten(ICBASE, 3),
]

OCBASE = [["*", 0, "*"], [1, 1, 0], ["*", 1, "*"]]
OUTSIDE_CORNERS = [
    OCBASE,
    rotaten(OCBASE, 1),
    rotaten(OCBASE, 2),
    rotaten(OCBASE, 3),
]

ENDPIECE = [["*", 0, "*"], [0, 1, 0], ["*", 1, "*"]]
ENDPIECES = [
    ENDPIECE,
    rotaten(ENDPIECE, 1),
    rotaten(ENDPIECE, 2),
    rotaten(ENDPIECE, 3),
]

ALL_CORNERS = INSIDE_CORNERS + OUTSIDE_CORNERS

if __name__ == "__main__":
    main()
