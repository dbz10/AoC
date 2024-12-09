from dataclasses import dataclass


def main(input_file="sample.txt"):
    input = open(input_file).read().strip()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")


def part1(input):
    # try a straightforward yet possibly inefficient way?
    disk = []
    nnz = 0
    for i, v in enumerate(input):
        if i % 2 == 0:
            disk.extend([int(i / 2)] * int(v))
            nnz += int(v)
        else:
            disk.extend([None] * int(v))

    left, right = 0, len(disk) - 1
    while left < right:
        while disk[right] is None:
            right -= 1
        while disk[left] is not None:
            left += 1
        if left >= right:
            break
        disk[left] = disk[right]
        disk[right] = None

    checksum = sum(i * v for i, v in enumerate(disk) if v is not None)
    return checksum


@dataclass
class Data:
    id: int | None
    size: int


def part2(input):
    # should have known
    disk = []
    for i, v in enumerate(input):
        if i % 2 == 0:
            disk.append(Data(int(i / 2), int(v)))
        else:
            disk.append(Data(None, int(v)))

    # Attempt to move each file exactly once in order of decreasing file ID number starting with the file with the highest file ID number. If there is no span of free space to the left of a file that is large enough to fit the file, the file does not move.
    # sigh

    right = len(disk) - 1
    while right > 0:
        # right points at the rightmost non-empty data
        while disk[right].id is None:
            right -= 1

        data_block = disk[right]
        empty_ptr = 0
        # scan from the first empty block towards the end for a region which can hold the data
        while (
            (disk[empty_ptr].size < data_block.size) or (disk[empty_ptr].id is not None)
        ) and empty_ptr < right:
            empty_ptr += 1
        if empty_ptr >= right:
            # couldn't find a compatible block of empty space
            right -= 1
            continue

        # ptr should now be pointing at the block of empty space which can hold `data_block`
        empty_block = disk[empty_ptr]
        if data_block.size == empty_block.size:  # the easy case, equal size
            empty_block.id, data_block.id = data_block.id, empty_block.id
        if data_block.size < empty_block.size:
            # split the empty space to insert the data
            disk.insert(empty_ptr, Data(data_block.id, data_block.size))  # copy it
            empty_block.size -= data_block.size
            data_block.id = None
            right += 1

        right -= 1

    disk = unfurl(disk)
    checksum = sum(i * v for i, v in enumerate(disk) if v is not None)
    return checksum


def unfurl(data):
    return [d.id for d in data for _ in range(d.size)]


if __name__ == "__main__":
    main()
