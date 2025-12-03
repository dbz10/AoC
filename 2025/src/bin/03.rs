use aoc2025;

fn main() {
    let content: Vec<Vec<u64>> = aoc2025::read_content()
        .lines()
        .map(|s| {
            s.chars()
                .map(|ch| ch.to_digit(10).unwrap() as u64)
                .collect()
        })
        .collect();

    println!("Part 1: {}", part1(&content));
    println!("Part 2: {}", part2(&content));
}

fn bank_joltage(batteries: &[u64], n_batteries: usize) -> u64 {
    let mut start = 0;
    let mut chosen: Vec<u64> = vec![];

    for remaining in (0..n_batteries).rev() {
        let end = batteries.len() - remaining;

        let (offset, &max) = batteries[start..end]
            .iter()
            .enumerate()
            .max_by_key(|&(_, v)| v)
            .unwrap();

        chosen.push(max);
        start += offset + 1;
    }

    chosen.into_iter().fold(0, |acc, d| acc * 10 + d)
}

fn part1(contents: &[Vec<u64>]) -> u64 {
    contents
        .into_iter()
        .map(|x: &Vec<u64>| bank_joltage(x, 2))
        .sum()
}

fn part2(contents: &[Vec<u64>]) -> u64 {
    contents
        .into_iter()
        .map(|x: &Vec<u64>| bank_joltage(x, 12))
        .sum()
}
