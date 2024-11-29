from __future__ import annotations
from dataclasses import dataclass
from queue import Queue

input_file = "inputs/day20.txt"

node_specs = open(input_file).read().split("\n")


@dataclass
class Transmission:
    source: str
    dest: str
    signal: int


class Node:
    def __init__(self, id: str, child_ids: list[str]) -> None:
        self.id = id
        self.child_ids = child_ids
        self.state = 0
        self.inputs = []  # not really needed here but who knows
        self.input_states = []  # not really needed here but who knows

    def proc(self, transmission: Transmission) -> list[Transmission]:
        pass

    def link_input(self, parent: str):
        self.inputs.append(parent)
        self.input_states.append(0)


class FlipFlop(Node):
    def __init__(self, id: str, child_ids: list[str]) -> None:
        super().__init__(id, child_ids)

    def proc(self, transmission: Transmission) -> list[Transmission]:
        if transmission.signal == 0:
            self.state = 1 - self.state
            return [Transmission(self.id, id, self.state) for id in self.child_ids]
        return []


class Conj(Node):
    def __init__(self, id: str, child_ids: list[str]) -> None:
        super().__init__(id, child_ids)

    def proc(self, transmission: Transmission) -> list[Transmission]:
        self.input_states[self.inputs.index(transmission.source)] = transmission.signal
        if all(self.input_states):
            return [Transmission(self.id, id, 0) for id in self.child_ids]
        return [Transmission(self.id, id, 1) for id in self.child_ids]


class Broadcaster(Node):
    def __init__(self, id: str, child_ids: list[str]) -> None:
        super().__init__(id, child_ids)

    def proc(self, transmission: Transmission) -> list[Transmission]:
        return [Transmission(self.id, id, 0) for id in self.child_ids]


class Output(Node):
    def __init__(self, id: str, child_ids: list[str]) -> None:
        super().__init__(id, child_ids)

    def proc(self, transmission: Transmission) -> list[Transmission]:
        return []


nodes: dict[str, Node] = {}
nodes["output"] = Output("output", [])
nodes["rx"] = Output("rx", [])
for node_spec in node_specs:
    l, r = node_spec.split(" -> ")
    children = r.split(", ")
    if l == "broadcaster":
        nodes["broadcaster"] = Broadcaster("broadcaster", children)

    elif l.startswith("%"):
        nodes[l[1:]] = FlipFlop(l[1:], children)
    elif l.startswith("&"):
        nodes[l[1:]] = Conj(l[1:], children)

extra_outputs = []
for v in nodes.values():
    for c in v.child_ids:
        nodes[c].link_input(v.id)


transmissions = Queue()
acc_low, acc_high = 0, 0
acc_history = [(0, 0)]
gateway_nodes = ["pl", "zv", "sd", "sk"]
for round in range(500000):
    transmissions.put(Transmission("button", "broadcaster", 0))
    acc_low += 1
    while not transmissions.empty():
        t = transmissions.get()
        # print(t)
        relays = nodes[t.dest].proc(t)
        # if len(relays) == 0:
        #     continue
        for relay in relays:
            transmissions.put(relay)
            if relay.signal == 0:
                acc_low += 1
            else:
                acc_high += 1

    acc_history.append((acc_low, acc_high))
    if all(n.state == 0 for n in nodes.values()):
        print(round)
        break

    if any(nodes[k].state == 1 for k in gateway_nodes) or round % 10000 == 0:
        print({k: nodes[k].state for k in gateway_nodes})


n_cycles = 1000 // (round + 1)
rem = 1000 % (round + 1)

res = (n_cycles * acc_low + acc_history[rem][0]) * (
    n_cycles * acc_high + acc_history[rem][1]
)
print("Part 1:", res)

# part 2....
import graphviz

g = graphviz.Digraph()
for k in nodes.keys():
    g.node(k)

status = ["rx"]
seen = set()
while status:
    n = status.pop()
    if n in seen:
        continue
    seen.add(n)
    node = nodes[n]
    nid = node.id
    if isinstance(node, Output):
        inputs = node.inputs
        # print(f"Condition for {n}: Receive low pulse from 1 of {inputs}")
        status.extend(inputs)
        for i in inputs:
            g.edge(i, n)
    if isinstance(node, FlipFlop):
        inputs = node.inputs
        # print(f"Condition for {n}: Receive low pulse from 1 of {inputs}")
        for i in inputs:
            g.edge(i, n)
        status.extend(inputs)
    if isinstance(node, Conj):
        inputs = node.inputs

        # print(f"Condition for {n}: Receive all high pulse from {inputs}")
        status.extend(inputs)
        for i in inputs:
            g.edge(i, n, style="dotted")

g.render(view=True)
# # kh, lz, tg, hn need to all send a positive pulse
