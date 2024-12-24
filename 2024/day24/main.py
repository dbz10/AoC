from copy import copy
import graphviz


def main(input_file="sample.txt"):
    input = open(input_file).read().split("\n\n")

    vars = {v.split(":")[0]: int(v.split(":")[1]) for v in input[0].splitlines()}
    equations = {
        v.split("->")[1].strip(): v.split("->")[0].strip()
        for v in (
            input[1]
            .replace("XOR", "^")
            .replace("AND", "&")
            .replace("OR", "|")
            .splitlines()
        )
    }

    print(f"Part 1: {part1(copy(vars), equations)}")
    print(f"Part 2: {part2(vars, equations)}")


def part1(vars, equations):
    target_len = len(vars) + len(equations)
    while True:
        for output, eq in equations.items():
            if output not in vars:
                [x1, op, x2] = eq.split()
                if x1 not in vars or x2 not in vars:
                    continue
                vars[output] = eval(f"{vars[x1]} {op} {vars[x2]}")
        if len(vars) == target_len:
            break

    zval = build_int(vars, "z")
    return zval


def part2(vars, equations):
    dot = graphviz.Graph()

    target_len = len(vars) + len(equations)

    # nvv ok to switch to a zero
    replacement_rules = {
        "z05": "hdt",
        "hdt": "z05",
        #
        "z09": "gbf",
        "gbf": "z09",
        #
        "mht": "jgt",
        "jgt": "mht",
        #
        "z30": "nbf",
        "nbf": "z30",
    }
    equations = {
        replacement_rules.get(output, output): eq for output, eq in equations.items()
    }
    while True:
        for output, eq in equations.items():
            if output not in vars:
                [x1, op, x2] = eq.split()
                if x1 not in vars or x2 not in vars:
                    continue
                vars[output] = eval(f"{vars[x1]} {op} {vars[x2]}")
        if len(vars) == target_len:
            break

    xval = build_int(vars, "x")
    yval = build_int(vars, "y")

    target_zval = xval + yval
    target_zstring = bin(target_zval)[2:]

    actual_zval = build_int(vars, "z")
    actual_ztring = bin(actual_zval)[2:]

    wrong_output_bits = [
        f"z{i:02}"
        for i, (tz, az) in enumerate(zip(target_zstring[::-1], actual_ztring[::-1]))
        if tz != az
    ]
    print(actual_zval, target_zval)
    print(sorted([(w, vars[w]) for w in wrong_output_bits]))

    for w in wrong_output_bits:
        print([{k: v} for k, v in equations.items() if k == w])

    def shapecolor(node):
        if node in wrong_output_bits:
            return "red"
        elif node.startswith("x"):
            return "green"
        elif node.startswith("y"):
            return "blue"
        return "black"

    for node in sorted(vars.keys()):
        dot.node(
            node,
            f"{node}: {vars[node]}",
            shape="box",
            color=shapecolor(node),
        )
    for output, eq in sorted(equations.items()):
        [x1, op, x2] = eq.split()
        int_node = f"{output}_int"
        dot.node(
            int_node, op, shape={"|": "circle", "^": "diamond", "&": "invtriangle"}[op]
        )
        dot.edge(x1, int_node)
        dot.edge(x2, int_node)
        dot.edge(int_node, output)

    dot.render()

    return ",".join(sorted(list(replacement_rules.keys())))


def build_int(vars, letter):
    xs = sorted([k for k in vars.keys() if k.startswith(letter)], reverse=True)
    sx = "".join([str(vars[x]) for x in xs])
    xval = int(sx, base=2)
    return xval


if __name__ == "__main__":
    main()
