use std::cmp::max;

use aoc2025;

fn main() {
    let content = aoc2025::read_content();
    let (ranges, ingredients) = content.split_once("\n\n").unwrap();

    let mut ranges: Vec<(usize, usize)> = ranges
        .lines()
        .map(|s| {
            let (lo, hi) = s.split_once("-").unwrap();
            (lo.parse().unwrap(), hi.parse().unwrap())
        })
        .collect();

    let ingredients: Vec<usize> = ingredients.lines().map(|s| s.parse().unwrap()).collect();

    println!("Part 1: {}", part1(&ranges, &ingredients));
    println!("Part 2: {}", part2(&mut ranges));
}

fn part1(ranges: &[(usize, usize)], ingredients: &[usize]) -> usize {
    ingredients
        .iter()
        .filter(|&i| ranges.iter().any(|(lo, hi)| (lo <= i) && (i <= hi)))
        .count()
}

fn part2(ranges: &mut [(usize, usize)]) -> i64 {
    ranges.sort_by_key(|(lo, _)| *lo);

    let mut r_condensed = vec![];

    for &(lo, hi) in ranges.iter() {
        if r_condensed.is_empty() {
            r_condensed.push((lo, hi));
            continue;
        }

        let latest = r_condensed.pop().unwrap();
        if lo <= latest.1 {
            r_condensed.push((latest.0, max(hi, latest.1)));
        } else {
            r_condensed.push(latest);
            r_condensed.push((lo, hi));
        }
    }

    r_condensed
        .into_iter()
        .map(|(lo, hi)| (hi - lo + 1) as i64)
        .sum()
}
