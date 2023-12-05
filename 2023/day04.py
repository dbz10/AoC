import os
import numpy as np
import re

input_file = os.path.join(os.path.dirname(__file__), "inputs/day04.txt")

with open(input_file,'r') as f:
    card_draws = (
        line.split(':')[1].strip()
        for line in f.readlines()
    )

card_draws = list(
    (cards.split('|')[0].strip().split(), set(cards.split('|')[1].strip().split()))
    for cards in card_draws
)

acc = 0
p2_ncards = np.ones(len(card_draws))

for (idx,draw) in enumerate(card_draws):
    nmatches = len([x for x in draw[1] if x in draw[0]])
    p2_ncards[idx+1:idx+nmatches+1] += p2_ncards[idx]
    if nmatches > 0:
        acc += 2**(nmatches-1)
        


print("Part 1:", acc)
print("Part 2:", p2_ncards.sum())


