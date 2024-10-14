use std::collections::{HashMap, HashSet};

use aoc_2021;

fn main() {
    let content = aoc_2021::read_content();
    let content: Vec<&str> = content.lines().collect();

    let p1 = part1(&content);
    println!("Part 1: {p1}");

    let p2 = part2(&content);
    println!("Part 2: {p2}");
}

fn part1(lines: &[&str]) -> u32 {
    let mut counter = 0;
    for &line in lines {
        let (_, output) = line
            .split_once('|')
            .expect("expected to split on delimiter");

        counter += output
            .trim()
            .split_ascii_whitespace()
            .filter(|&s| s.len() == 2 || s.len() == 3 || s.len() == 4 || s.len() == 7)
            .count()
    }
    counter.try_into().expect("parse error")
}

fn part2(lines: &[&str]) -> u32 {
    let mut counter = 0;

    for &line in lines {
        let (l, r) = line
            .trim()
            .split_once('|')
            .expect("expected to split on delimiter");

        let l = l
            .split_ascii_whitespace()
            .map(|z| {
                let mut v: Vec<char> = z.chars().collect();
                v.sort_unstable();
                v.into_iter().collect()
            })
            .collect::<Vec<String>>();

        let r = r
            .split_ascii_whitespace()
            .map(|z| {
                let mut v: Vec<char> = z.chars().collect();
                v.sort_unstable();
                v.into_iter().collect()
            })
            .collect::<Vec<String>>();

        let mapping = derive_mapping(&l);

        counter += r.iter().fold(0, |acc, pattern| {
            10 * acc
                + mapping
                    .get(pattern)
                    .expect(&format!("coulnd't find key {pattern}"))
        });
    }

    counter
}

fn derive_mapping(patterns: &[String]) -> HashMap<String, u32> {
    let mut mapping = HashMap::new();
    // Approximate logic....
    // The character appearing in the string of length 3 but not in length 2 maps to `a`
    let one = patterns
        .iter()
        .find(|p| p.len() == 2)
        .expect("expected to find one of length 2");
    let four = patterns
        .iter()
        .find(|p| p.len() == 4)
        .expect("expected to find one of length 4");

    for pattern in patterns.iter() {
        match pattern.len() {
            2 => mapping.insert(pattern.to_string(), 1),
            4 => mapping.insert(pattern.to_string(), 4),
            3 => mapping.insert(pattern.to_string(), 7),
            7 => mapping.insert(pattern.to_string(), 8),
            len => {
                match [
                    len,
                    letter_count(&pattern)
                        .intersection(&letter_count(&one))
                        .count(),
                    letter_count(&pattern)
                        .intersection(&letter_count(&four))
                        .count(),
                ] {
                    [6, 2, 4] => mapping.insert(pattern.to_string(), 9),
                    [6, 2, 3] => mapping.insert(pattern.to_string(), 0),
                    [6, 1, _] => mapping.insert(pattern.to_string(), 6),
                    [5, 2, _] => mapping.insert(pattern.to_string(), 3),
                    [5, 1, 2] => mapping.insert(pattern.to_string(), 2),
                    [5, 1, 3] => mapping.insert(pattern.to_string(), 5),
                    _ => unreachable!(),
                }
            }
        };
    }

    mapping
}

fn letter_count(s: &str) -> HashSet<char> {
    let mut h = HashSet::new();
    s.chars().for_each(|c| {
        h.insert(c);
    });
    h
}
