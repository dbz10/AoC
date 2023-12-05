import os
import numpy as np
import re

input_file = os.path.join(os.path.dirname(__file__), "inputs/day03.txt")

def get_neighbors(x,y, Lx, Ly):
    return [
        (x+dx,y+dy) for dx in [-1,0,1] for dy in [-1,0,1]
        if 0 <= x + dx < Lx
        and 0 <= y + dy < Ly
        and not (dx == 0 and dy == 0)
    ]

arrs = []
number_spans = []
with open(input_file,'r') as f:
    for (idx,line) in enumerate(f.readlines()):
        arrs.append(np.array([l for l in line if l != '\n']))
        number_spans.append(
            (idx, [m for m in re.finditer(r"(\d+)", line)])
        )

grid = np.vstack(arrs)
Ly, Lx = grid.shape  

symbol_re = re.compile(r"[^\d\.]")

acc = 0
for y, matches in number_spans:
    for match in matches:
        if any(
            symbol_re.match(grid[n_y, n_x]) is not None
            for x in range(match.start(), match.end())
            for n_x, n_y in get_neighbors(x,y,Lx,Ly)
        ):
            acc += int(match[0])
        
    
print("Part 1:", acc)

acc = 0
for y in range(Ly):
    for x in range(Lx):
        if grid[y,x] == "*":
            adjacent_numbers = {}
            # check spans in the current row and y+-1
            neighors = get_neighbors(x,y,Lx,Ly)
            for nx, ny in neighors:
                for span in number_spans[ny][1]:
                    if span.start() <= nx < span.end():
                        adjacent_numbers[span] = (nx, ny)
            if len(adjacent_numbers) == 2:
                acc += np.prod([int(s[0]) for s in adjacent_numbers.keys()])
            
print("Part 2:", acc)
