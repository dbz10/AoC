from __future__ import annotations
from dataclasses import dataclass
import os
import numpy as np
from concurrent.futures import ProcessPoolExecutor

input_file = os.path.join(os.path.dirname(__file__), "inputs/day05.txt")

data = open(input_file).read()

sections = data.split("\n\n")

seeds = [int(s) for s in sections[0].split(":")[1].split()]
converters = [
    [[int(x) for x in z.split()] for z in s.split("\n")[1:]] for s in sections[1:]
]


def dive(seed, converters):
    seeds = [seed]
    for converter in converters:
        for mapping in converter:
            source = mapping[1]
            dest = mapping[0]
            range = mapping[2]
            if source <= seed < source + range:
                seed = dest + (seed - source)
                break
        seeds.append(seed)
    return seeds[-1]


res = min(dive(seed, converters) for seed in seeds)
print("Part 1:", res)

# TODO: clean me up later
# def process_pool(start, length):
#     return min(
#         (dive(seed, converters) for seed in range(start, start + length)),
#         key=lambda z: z[1],
#     )


# pairs = [(seeds[2 * i], seeds[2 * i + 1]) for i in range(int(len(seeds) / 2))]


# def go():
#     with ProcessPoolExecutor() as pool:
#         futures = [pool.submit(process_pool, pair[0], pair[1]) for pair in pairs]

#     results = [f.result() for f in futures]
#     return results


# if __name__ == "__main__":

# #     results = go()
# #     print(results)


@dataclass
class Range:
    low: int
    high: int

    def overlap(self, other: Range) -> Range | None:
        if (
            (other.low <= self.low <= other.high)
            or (other.low <= self.high <= other.high)
            or (self.low <= other.low and self.high >= other.high)
        ):
            return Range(max(self.low, other.low), min(self.high, other.high))
        return None

    def map_through(self, m: Mapping) -> tuple[list[Range], list[Range]]:
        # left part: unmapped. right part: mapped
        overlap = self.overlap(m.source)
        # print(overlap)
        if overlap is None:
            return [self], []

        mapped_overlap = Range(
            m.target.low + (overlap.low - m.source.low),
            m.target.low + (overlap.high - m.source.low),
        )

        if self.low < overlap.low and self.high <= overlap.high:
            return [Range(self.low, overlap.low - 1)], [mapped_overlap]
        elif self.low >= overlap.low and self.high <= overlap.high:
            return [], [mapped_overlap]
        elif self.low >= overlap.low and self.high > overlap.high:
            return [Range(overlap.high + 1, self.high)], [mapped_overlap]
        else:
            return [
                Range(self.low, overlap.low - 1),
                Range(overlap.high + 1, self.high),
            ], [mapped_overlap]


@dataclass
class Mapping:
    source: Range
    target: Range


seeds_part_2 = [
    Range(seeds[2 * i], seeds[2 * i] + seeds[2 * i + 1] - 1)
    for i in range(int(len(seeds) / 2))
]

channels = [
    [
        Mapping(Range(z[1], z[1] + z[2] - 1), Range(z[0], z[0] + z[2] - 1))
        for z in mappings
    ]
    for mappings in converters
]


def dive_p2(seed: Range, channels: list[list[Mapping]]):
    x = [seed]
    for channel in channels:
        tmp = (x, [])
        # print(x)
        for mapping in channel:
            mapping_result = [s.map_through(mapping) for s in tmp[0]]
            unmapped_part = [z for part in mapping_result for z in part[0]]
            mapped_part = [z for part in mapping_result for z in part[1]]
            tmp = (unmapped_part, mapped_part + tmp[1])
        x = tmp[0] + tmp[1]
    # print(x)
    return x


res = sorted(
    [x for seed in seeds_part_2 for x in dive_p2(seed, channels)], key=lambda z: z.low
)[0].low
print("Part 2:", res)
