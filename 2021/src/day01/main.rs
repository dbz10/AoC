use aoc_2021::read_content;

fn main() {
    let content = read_content();

    let part1 = part1(&content);
    println!("Part 1: {part1}");

    let part2 = part2(&content);
    println!("Part 2: {part2}");
}

fn part1(data: &str) -> usize {
    data.lines()
        .map(|x| x.parse::<usize>().expect("parse error"))
        .collect::<Vec<usize>>()
        .windows(2)
        .filter(|&w| w[1] > w[0])
        .count()
}

fn part2(data: &str) -> usize {
    data.lines()
        .map(|x| x.parse::<usize>().expect("parse error"))
        .collect::<Vec<usize>>()
        .windows(4)
        .filter(|&w| w[3] > w[0])
        .count()
}
