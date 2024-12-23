from collections import defaultdict
from concurrent.futures import ProcessPoolExecutor


def main(input_file="sample.txt"):
    input = open(input_file).read().splitlines()

    nodes = defaultdict(set[str])
    for row in input:
        [l, r] = row.split("-")
        nodes[l].add(r)
        nodes[r].add(l)

    print(f"Part 1: {part1(nodes)}")
    print(f"Part 2: {part2(nodes)}")


def part1(nodes):
    all_nodes = list(nodes.keys())
    # idk
    sets_of_three = set()
    for n1 in all_nodes:
        for n2 in reduce_neighbors([n1], nodes):
            for n3 in reduce_neighbors([n1, n2], nodes):
                if not any([n.startswith("t") for n in [n1, n2, n3]]):
                    continue
                sets_of_three.add(tuple(sorted([n1, n2, n3])))
    return len(sets_of_three)


def part2(nodes):
    all_nodes = sorted(list(set(nodes.keys())))
    with ProcessPoolExecutor() as pool:
        nodes_largest_clusters_futures = nodes_largest_clusters = {
            node: pool.submit(build_largest_cluster, node, nodes) for node in all_nodes
        }
    nodes_largest_clusters = {
        node: future.result() for node, future in nodes_largest_clusters_futures.items()
    }
    biggest_cluster_size = max(len(v[0]) for v in nodes_largest_clusters.values())
    a_biggest_cluster = [
        v[0]
        for v in nodes_largest_clusters.values()
        if len(v[0]) == biggest_cluster_size
    ][0]
    return ",".join(sorted(a_biggest_cluster))


def build_largest_cluster(node: str, links: dict[str, set[str]]) -> list[str]:
    # i dunno, seems plausible
    clusters = [[node, n] for n in links[node]]
    while True:
        this_pass = clusters
        possible_next_clusters = []
        for pc in this_pass:
            expansion_candidates = reduce_neighbors(pc, links)
            this_possible_next_clusters = [
                tuple(sorted(pc + [ec])) for ec in expansion_candidates
            ]
            possible_next_clusters.extend(this_possible_next_clusters)

        deduplicated_clusters = [list(x) for x in list(set(possible_next_clusters))]
        if len(deduplicated_clusters) == 0:
            break
        clusters = deduplicated_clusters
    lmax = max(len(c) for c in clusters)
    return [c for c in clusters if len(c) == lmax]


def reduce_neighbors(
    nodes: list[str],
    links: dict[str, set[str]],
) -> list[str]:
    out = links[nodes[0]]
    for n in nodes[1:]:
        out = out & links[n]
    return out


if __name__ == "__main__":
    main()
