from __future__ import annotations
from dataclasses import dataclass
from copy import copy


@dataclass
class GardenPatch:
    label: str
    cluster: int

    def listen(self, other: GardenPatch):
        if other.label == self.label and other.cluster <= self.cluster:
            self.cluster = other.cluster

    # def __eq__(self, other: GardenPatch):
    #     return self.label == other.label and self.cluster == other.cluster


def main(input_file="sample.txt"):
    input = list(map(list, open(input_file).read().splitlines()))
    garden = {}

    cluster_id = 0
    # message propagation?
    for y, row in enumerate(input):
        for x, ch in enumerate(row):
            garden[(x, y)] = GardenPatch(ch, cluster_id)
            cluster_id += 1

    for _ in range(100):
        new = copy(garden)
        for (x, y), v in new.items():
            for xp, yp in [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]:
                if (xp, yp) in new:
                    new[(xp, yp)].listen(v)
        # if new==garden:
        #     print([v == garden[k] for k, v in new.items()])
        #     print(_, "?")
        #     garden = new
        #     break

    # remap the clusters into continguous range just for simplicity
    remapping = {
        original_id: new_id
        for new_id, original_id in enumerate(
            sorted(set([p.cluster for p in garden.values()]))
        )
    }
    for k in garden.keys():
        garden[k].cluster = remapping[garden[k].cluster]

    # for y in range(len(input)):
    #     for x in range(len(input[0])):
    #         print(
    #             "({}, {:2d}) ".format(garden[(x, y)].label, garden[(x, y)].cluster),
    #             end="",
    #         )
    #     print("")
    print(f"Part 1: {part1(garden)}")
    print(f"Part 2: {part2(input)}")


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


def part2(input):
    return


if __name__ == "__main__":
    main()
