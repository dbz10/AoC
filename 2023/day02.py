import os
import re
from collections import namedtuple, defaultdict

input_file = os.path.join(os.path.dirname(__file__), "inputs/day02.txt")


game_maxes = []
with open(input_file,'r') as f:
    for idx, line in enumerate(f.readlines(), start = 1):
        gm = defaultdict(int)
        gm['id'] = idx
        rounds = line.split(":")[1].split(';')
        for round in rounds:
            draws = [d.strip() for d in round.split(',')]
            for draw in draws:
                number, color = draw.split(' ')
                gm[color] = max(gm[color], int(number))
        game_maxes.append(gm)

red_cubes = 12
green_cubes = 13
blue_cubes = 14

part_1 = sum(
    g['id'] for g in game_maxes 
    if g['red'] <= red_cubes 
    and g['green'] <= green_cubes 
    and g['blue'] <= blue_cubes
)

print("Part 1:", part_1)

part_2 = sum(
    g['red'] * g['green'] * g['blue']
    for g in game_maxes
)

print("Part 2:", part_2)