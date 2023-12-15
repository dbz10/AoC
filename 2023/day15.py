input_file = "inputs/day15.txt"
from collections import OrderedDict, defaultdict

codes = [c for c in open(input_file).read().split(",")]


def hash(chars):
    s = 0
    for ch in list(chars):
        s = ((s + ord(ch)) * 17) % 256
    return s


acc = 0
for c in codes:
    acc += hash(c)

print("Part 1:", acc)


class Cache:
    def __init__(self, source) -> None:
        self.cache = {}
        self.source = source

    def get(self, key):
        check = self.cache.get(key)
        if check:
            return check
        res = self.source(key)
        self.cache[key] = res
        return res


lens_library = defaultdict(dict)  # python dicts remember insertion order
label_cache = Cache(hash)
for c in codes:
    if "=" in c:
        l, r = c.split("=")
        label = label_cache.get(l)
        lens_library[label][l] = int(r)
    if "-" in c:
        l = c[:-1]
        label = label_cache.get(l)
        lens_library[label].pop(l, None)

acc = 0
for k, v in lens_library.items():
    for i, f in enumerate(v.values(), start=1):
        acc += (k + 1) * i * f

print("Part 2:", acc)
