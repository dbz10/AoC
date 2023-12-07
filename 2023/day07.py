import os
from collections import Counter

input_file = os.path.join(os.path.dirname(__file__), "inputs/day07.txt")


def hand_type(hand: str) -> int:
    c = Counter(hand)
    if len(c) == 1:
        return 5
    nj = c.get("J")
    if nj:
        most_common = c.most_common(1)[0][0]
        if most_common == "J":
            most_common = c.most_common(2)[1][0]
        c[most_common] += nj
        c["J"] = 0
    c = sorted(list(c.values()), reverse=True)
    if c[0] == 5:
        return 5
    elif c[0] == 4:
        return 4
    elif c[:2] == [3, 2]:
        return 3
    elif c[:2] == [3, 1]:
        return 2
    elif c[:2] == [2, 2]:
        return 1
    elif c[0] == 2:
        return 0
    return -1


face_card_mapping = {"T": 10, "J": 1, "Q": 12, "K": 13, "A": 14}


line_gen = (line.split() for line in open(input_file))
hand_gen = (
    ([hand_type(hand)] + [int(face_card_mapping.get(x, x)) for x in hand], bid)
    for hand, bid in line_gen
)
hands_ranked = sorted(hand_gen, key=lambda z: z[0])

acc = 0
for rank, (hand, bid) in enumerate(hands_ranked, start=1):
    acc += rank * int(bid)


print("Part 1/2:", acc)
