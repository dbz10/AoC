from itertools import permutations
from typing import Literal
from tqdm import tqdm

KEYPAD = {
    "7": (0, 0),
    "8": (1, 0),
    "9": (2, 0),
    "4": (0, 1),
    "5": (1, 1),
    "6": (2, 1),
    "1": (0, 2),
    "2": (1, 2),
    "3": (2, 2),
    "0": (1, 3),
    "A": (2, 3),
}

DIRPAD = {
    "^": (1, 0),
    "A": (2, 0),
    "<": (0, 1),
    "v": (1, 1),
    ">": (2, 1),
}

MOVEMENT_MAP = {
    (1, 0): ">",
    (-1, 0): "<",
    (0, -1): "^",
    (0, 1): "v",
}

INVERSE_MOVEMENT_MAP = {v: k for k, v in MOVEMENT_MAP.items()}
KEYPAD_START = (2, 3)
DIRPAD_START = (2, 0)


# Basic movement optimization heuristics:
# move in as long as possible straight lines
# finish as close to 'A' as possible?

# that means that on the dir pad:
# if there is a '<' component, move


def main(input_file="sample.txt"):
    input = open(input_file).read().splitlines()
    print(f"Part 1: {part1(input,2)}")
    print(f"Part 2: {part1(input,25)}")


def part1(input, len_robot_chain: int = 2):
    outs = []
    for code in input:
        # print(code)
        targets = [KEYPAD[t] for t in list(code)]
        full_sequences = []
        keypad_sequences = derive_sequence(
            target_positions=targets,
            start_position=KEYPAD_START,
            pad_type="keypad",
        )
        target_seq = keypad_sequences
        for i in range(len_robot_chain):
            print(i, len(target_seq))
            target_seq = evolve(target_seq)

        ss = target_seq[0]

        # if any(
        #     "".join(s)
        #     == "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A"
        #     for s in d2_sequences
        # ):
        #     print("found something")
        score = int(code[:-1]) * len(ss)
        outs.append((code, len(ss), score))
    print(outs)
    return sum(o[2] for o in outs)


def evolve(target_sequences: list[list[str]]):
    output_sequences = []
    mcur = 1e12
    for target_sequence in target_sequences:
        targets = [DIRPAD[t] for t in target_sequence]
        o_sequences_int = derive_sequence(targets, DIRPAD_START, pad_type="dirpad")
        for o1s in o_sequences_int:
            if len(o1s) <= mcur:
                output_sequences.append(o1s)
                mcur = len(o1s)
    lmin = min(len(d) for d in output_sequences)
    output_sequences = [d for d in output_sequences if len(d) == lmin]
    optimal_sequence_contiguity = max(
        score_sequence_contiguity(d) for d in output_sequences
    )
    output_sequences = [
        d
        for d in output_sequences
        if score_sequence_contiguity(d) == optimal_sequence_contiguity
    ]
    optimal_sequence_contiguity = max(score_contiguous_A(d) for d in output_sequences)
    output_sequences = [
        d
        for d in output_sequences
        if score_contiguous_A(d) == optimal_sequence_contiguity
    ]
    return output_sequences


def derive_sequence(
    target_positions: list[tuple[int, int]],
    start_position: tuple[int, int],
    pad_type: Literal["keypad", "dirpad"],
) -> list[list[str]]:
    sequences: list[list[str]] = [[]]
    current_position = start_position
    for target_position in target_positions:
        button_sequence = operate_pad(current_position, target_position, pad_type)
        # maybe?
        lmin = min([len(d) for d in button_sequence])
        button_sequence = [d for d in button_sequence if len(d) == lmin]
        optimal_sequence_contiguity = max(
            score_sequence_contiguity(d) for d in button_sequence
        )
        button_sequence = [
            d
            for d in button_sequence
            if score_sequence_contiguity(d) == optimal_sequence_contiguity
        ]
        optimal_sequence_contiguity = max(
            score_contiguous_A(d) for d in button_sequence
        )
        button_sequence = [
            d
            for d in button_sequence
            if score_contiguous_A(d) == optimal_sequence_contiguity
        ]
        new_sequences = []
        for full_seq in sequences:
            for intermediate_d1_seq in button_sequence:
                new_sequences.append(full_seq + intermediate_d1_seq + ["A"])
        sequences = new_sequences
        current_position = target_position
    return sequences


def score_sequence_contiguity(s: list[str]) -> int:
    # i have no idea why i'm utilizing this logic at three separate places
    return sum([c == n for c, n in zip(s[:-1], s[1:])])


def score_contiguous_A(s: list[str]) -> int:
    return sum([c == n == "A" for c, n in zip(s[:-1], s[1:])])


def operate_pad(
    current_position: tuple[int, int],
    target_position: tuple[int, int],
    pad_type: Literal["keypad", "dirpad"],
) -> list[list[str]]:
    dx, dy = (
        target_position[0] - current_position[0],
        target_position[1] - current_position[1],
    )
    abs_dx, abs_dy, sgn_x, sgn_y = (
        abs(dx),
        abs(dy),
        int(dx / abs(dx)) if dx != 0 else 1,
        int(dy / abs(dy)) if dy != 0 else 1,
    )
    base_move_set = [(sgn_x, 0)] * abs_dx + [(0, sgn_y)] * abs_dy
    base_move_buttons = [MOVEMENT_MAP[m] for m in base_move_set]

    # try to efficiently construct unique button press sequences?
    # probably not worth it since the sequence would be at most 3-4 elements long.

    move_perms = permutations(base_move_buttons)
    button_sequences = list(map(list, set(permutations(base_move_buttons))))
    # enforce that the arm needs to stay within the domain of the board at all times
    # i.e. it can't go over the blank space
    pad_locations = KEYPAD.values() if pad_type == "keypad" else DIRPAD.values()
    for move_maybe in move_perms:
        button_sequence = list(move_maybe)
        if not validate_keypad_button_sequence(
            current_position, button_sequence, pad_locations
        ):
            continue
    allowed_button_sequences = [
        b
        for b in button_sequences
        if validate_keypad_button_sequence(current_position, b, pad_locations)
    ]
    button_sequence_contiguity_scores = [
        score_sequence_contiguity(s) for s in allowed_button_sequences
    ]
    optimal_contiguity = max(button_sequence_contiguity_scores)
    filtered_button_sequences = [
        a
        for a in allowed_button_sequences
        if score_sequence_contiguity(a) == optimal_contiguity
    ]
    return filtered_button_sequences


def validate_keypad_button_sequence(
    current_position: tuple[int, int],
    button_sequence: list[str],
    allowed_hovers: list[tuple[int, int]],
):
    positions_sequence = [current_position]
    for button_press in button_sequence:
        delta = INVERSE_MOVEMENT_MAP[button_press]
        c = positions_sequence[-1]
        next_position = (c[0] + delta[0], c[1] + delta[1])
        if next_position not in allowed_hovers:
            return False
        positions_sequence.append(next_position)
    return True


def part2(input):
    return


if __name__ == "__main__":
    main()
