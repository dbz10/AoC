use std::collections::HashMap;

use aoc_2021;

fn main() {
    let content = aoc_2021::read_content();

    let (polymer_template, _pair_insertion_rules) = content.split_once("\n\n").unwrap();
    let mut pair_insertion_rules = HashMap::new();
    _pair_insertion_rules.lines().for_each(|line| {
        let (pair, insertion) = line.split_once(" -> ").unwrap();
        let mut chs = pair.chars();
        let l = chs.next().unwrap();
        let r = chs.next().unwrap();
        pair_insertion_rules.insert((l, r), insertion.parse::<char>().unwrap());
    });
    println!(
        "Part 1: {}",
        part1(polymer_template, &pair_insertion_rules, 10)
    );
    println!(
        "Part 2: {}",
        part1(polymer_template, &pair_insertion_rules, 40)
    );
}

fn part1(
    polymer_template: &str,
    pair_insertion_rules: &HashMap<(char, char), char>,
    steps: usize,
) -> usize {
    let polymer_pairs = genpolymer(polymer_template, pair_insertion_rules, steps);
    let mut base_counts = HashMap::new();
    for (&(l, r), count) in polymer_pairs.iter() {
        *base_counts.entry(l).or_insert(0) += count;
        *base_counts.entry(r).or_insert(0) += count;
    }

    // In the end I just tried both with/without the + 1 through trial and error, to account for whether
    // the double counting of the first and last characters cancels itself or not.
    (base_counts.values().max().unwrap() - base_counts.values().min().unwrap()) / 2 + 1
}

fn genpolymer(
    polymer_template: &str,
    pair_insertion_rules: &HashMap<(char, char), char>,
    steps: usize,
) -> HashMap<(char, char), usize> {
    let mut polymer_pairs = HashMap::new();
    polymer_template
        .chars()
        .collect::<Vec<char>>()
        .windows(2)
        .for_each(|z| *polymer_pairs.entry((z[0], z[1])).or_insert(0) += 1);

    for _ in 0..steps {
        let mut to_insert = HashMap::new();
        for (&(l, r), count) in polymer_pairs.iter() {
            let insertion = pair_insertion_rules[&(l, r)];
            *to_insert.entry((l, insertion)).or_insert(0) += count;
            *to_insert.entry((insertion, r)).or_insert(0) += count;
        }
        polymer_pairs = to_insert
    }
    polymer_pairs
}
