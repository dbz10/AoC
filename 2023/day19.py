from dataclasses import dataclass
from typing import Any, Callable


input_file = "inputs/day19.txt"

workflows_specs, inputs = map(
    lambda x: x.split("\n"), open(input_file).read().split("\n\n")
)


@dataclass
class MachinePart:
    x: int
    m: int
    a: int
    s: int

    @classmethod
    def from_string(cls, input):
        parts = list(
            map(
                lambda x: int(x.split("=")[1]),
                input.replace("{", "").replace("}", "").split(","),
            )
        )
        return MachinePart(parts[0], parts[1], parts[2], parts[3])

    def sum(self):
        return self.x + self.m + self.a + self.s


@dataclass
class Flow:
    condition: Callable
    target: str


class Workflow:
    def __init__(self, spec: str) -> None:
        self.id = spec.split("{")[0]
        flows = spec.split("{")[1].replace("}", "").split(",")
        self.flows = []

        for flow in flows[:-1]:
            condition, target = flow.split(":")
            c = condition[1]
            p, t = condition.split(c)
            self.flows.append(Flow(eval(f"lambda z: z.{p} {c} {int(t)}"), target))

        self.flows.append(Flow(lambda x: True, flows[-1]))

    def evaluate(self, p: MachinePart):
        for flow in self.flows:
            if flow.condition(p):
                next = flow.target
                break
        return next


workflows = {w.id: w for w in [Workflow(w) for w in workflows_specs]}
inputs = [MachinePart.from_string(p) for p in inputs]


acc = 0
for part in inputs:
    wf = workflows["in"]
    while True:
        next = wf.evaluate(part)
        if next in ["A", "R"]:
            break
        wf = workflows[next]

    if next == "A":
        acc += part.sum()

print("Part 1:", acc)


@dataclass
class Range:
    min: int
    max: int

    def size(self):
        return self.max - self.min


@dataclass
class QuantumMachinePart:
    x: Range
    m: Range
    a: Range
    s: Range

    def split(self, field, threshold, direction):
        v = getattr(self, field)
        if direction == "<":
            if threshold > v.max:
                return [self, None]
            if threshold <= v.min:
                return [None, self]

            go = QuantumMachinePart(
                **(self.__dict__ | {field: Range(v.min, threshold - 1)})
            )
            stay = QuantumMachinePart(
                **(self.__dict__ | {field: Range(threshold - 1, v.max)})
            )
            return [go, stay]
        if direction == ">":
            if threshold > v.max:
                return [None, self]
            if threshold <= v.min:
                return [self, None]

            go = QuantumMachinePart(
                **(self.__dict__ | {field: Range(threshold, v.max)})
            )
            stay = QuantumMachinePart(
                **(self.__dict__ | {field: Range(v.min, threshold)})
            )
            return [go, stay]

    def size(self):
        return self.x.size() * self.m.size() * self.a.size() * self.s.size()


class Sieve:
    def __init__(self, spec: str) -> None:
        self.id = spec.split("{")[0]
        flows = spec.split("{")[1].replace("}", "").split(",")
        self.flows = []

        for flow in flows[:-1]:
            condition, target = flow.split(":")

            c = condition[1]
            p, t = condition.split(c)

            def split(p, t, c):
                return lambda z: z.split(p, t, c)

            self.flows.append(Flow(split(p, int(t), c), target))

        self.flows.append(Flow(lambda x: [x, None], flows[-1]))

    def percolate(self, p: QuantumMachinePart) -> list[tuple[QuantumMachinePart, str]]:
        out = []
        for flow in self.flows:
            go, stay = flow.condition(p)
            out.append((go, flow.target))
            if stay is None:
                return out
            p = stay
        return out


start = [
    (
        QuantumMachinePart(
            Range(0, 4000),
            Range(0, 4000),
            Range(0, 4000),
            Range(0, 4000),
        ),
        "in",
    )
]


sieves = {w.id: w for w in [Sieve(w) for w in workflows_specs]}

acc = 0
while start:
    qmp, sid = start.pop()

    nexts = [n for n in sieves[sid].percolate(qmp) if n[0] is not None]
    for next in nexts:
        if next[1] == "A":
            acc += next[0].size()
        elif next[1] == "R":
            continue
        else:
            start.append(next)

print("Part 2:", acc)
