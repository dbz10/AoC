import re
import numpy as np
from copy import copy
import matplotlib.pyplot as plt

# LX, LY = 11, 7
LX, LY = 101, 103


def main(input_file="sample.txt"):
    input = np.array(
        [
            list(map(int, re.findall("-?\d+", line)))
            for line in open(input_file).read().splitlines()
        ]
    )
    print(f"Part 1: {part1(copy(input))}")
    print(f"Part 2: {part2(input)}")


def part1(input):
    for _ in range(100):
        input[:, 0] = np.mod(input[:, 0] + input[:, 2], LX)
        input[:, 1] = np.mod(input[:, 1] + input[:, 3], LY)

    input = input[(input[:, 0] != LX // 2) & (input[:, 1] != LY // 2)]
    # breakpoint()
    prod = 1
    for qx in [0, 1]:
        for qy in [0, 1]:
            prod *= len(
                input[
                    (LX // 2 * qx <= input[:, 0])
                    & (input[:, 0] <= LX // 2 * (qx + 1))
                    & (LY // 2 * qy <= input[:, 1])
                    & (input[:, 1] <= LY // 2 * (qy + 1))
                ]
            )

    return prod


def part2(input):
    # what the hell
    visited_configurations = set()
    for iteration in range(10_000):
        config_memo = tuple(sorted(set(zip(input[:, 0], input[:, 1]))))
        if config_memo in visited_configurations:
            break
        visited_configurations.add(config_memo)
        input[:, 0] = np.mod(input[:, 0] + input[:, 2], LX)
        input[:, 1] = np.mod(input[:, 1] + input[:, 3], LY)

        tree_maybe = np.zeros((LY, LX))
        tree_maybe[input[:, 1], input[:, 0]] = 1

        fig, ax = plt.subplots()
        ax.imshow(tree_maybe)
        plt.savefig(f"day14/imgs/{iteration}.png")
        plt.close()
    return


if __name__ == "__main__":
    main()


# 2149 too low
# try 2857
