mapping = {"L": 0, "R": 1}

directions = "LRLLLRRLRRLRRLRRLLRRLRRLLRRRLLRRLRRLRRLRRLRLRLLLLLRRLRRLLRLRRRLLRRLRLLLLLLLRRLRLRRRLRRLRRRLRRLLLRRLLRRRLLRRRLRRLRLRRRLRRRLRLRLLRRRLRRRLRRLLRRRLRLRRLLRLLRRLLRRLRRRLRRLRLRRLLRRRLRRRLRRRLRLRLRLRRRLLRRRLRLRRLLRRLRRLRRLRLLRRLLRRRLRRRLRRLRRLRLLRRLRLRRLRRRLRRRLRRLRLRRRLRRRLRLLLRRLRLLRRRR"
# directions = "LR"
wrap_len = len(directions)
paths = {line[0:3]: (line[7:10], line[12:15]) for line in open("input.txt")}

steps = 0

node = "AAA"
while True:
    node = paths[node][mapping[directions[steps % wrap_len]]]
    steps += 1
    if node == "ZZZ":
        break

print("Part 1:", steps)

# Part 2
steps = 0
start_nodes = [n for n in paths.keys() if n.endswith("A")]
next_nodes = [n for n in paths.keys() if n.endswith("A")]
finish_lens = set()
while len(next_nodes) > 0:
    for i in range(len(next_nodes)):
        next_nodes[i] = paths[next_nodes[i]][mapping[directions[steps % wrap_len]]]
    steps += 1
    for n in next_nodes:
        if n.endswith("Z"):
            finish_lens.add(steps)
            next_nodes.remove(n)


from collections import Counter


def factorize(n):
    f = []
    i = 2
    t = n
    while i <= n / 2:
        if t % i:
            i += 1
        else:
            f.append(i)
            t = t / i
    return Counter(f)


prime_factors = [factorize(n) for n in finish_lens]
c = prime_factors[0]
for z in prime_factors:
    c = c | z


acc = 1
for k, v in c.items():
    acc *= k**v

print("Part 2:", acc)
