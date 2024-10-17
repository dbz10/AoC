use std::collections::HashMap;

use aoc_2021;
const POSITIVES: [char; 4] = ['(', '[', '{', '<'];
// const NEGATIVES: [char; 4] = ['(', '[', '{', '<'];

fn main() {
    let content = aoc_2021::read_content();

    println!("Part 1: {}", part1(&content));
    println!("Part 2: {}", part2(&content));
}

fn part1(content: &str) -> u32 {
    let point_values: HashMap<char, u32> =
        HashMap::from([(')', 3), (']', 57), ('}', 1197), ('>', 25137)]);

    let remap = HashMap::from([(')', '('), (']', '['), ('}', '{'), ('>', '<')]);

    content
        .lines()
        .filter_map(|l| checkline(l, &remap))
        .map(|ch| *point_values.get(&ch).unwrap())
        .sum()
}

fn part2(content: &str) -> usize {
    let point_values: HashMap<char, usize> =
        HashMap::from([(')', 1), (']', 2), ('}', 3), ('>', 4)]);
    let bmap = HashMap::from([(')', '('), (']', '['), ('}', '{'), ('>', '<')]);
    let fmap = HashMap::from([('(', ')'), ('[', ']'), ('{', '}'), ('<', '>')]);

    let mut scores: Vec<usize> = content
        .lines()
        .filter_map(|l| finishline(l, &bmap, &fmap, &point_values))
        .collect();

    scores.sort();
    scores[scores.len() / 2]
}

fn checkline(line: &str, remap: &HashMap<char, char>) -> Option<char> {
    let mut stack = Vec::new();
    for char in line.chars() {
        if POSITIVES.contains(&char) {
            stack.push(char);
        } else {
            if stack.len() == 0 {
                return Some(char);
            }
            let test = remap[&char];
            if *stack.last().unwrap() != test {
                return Some(char);
            }
            stack.pop();
        }
    }
    None
}

fn finishline(
    line: &str,
    bmap: &HashMap<char, char>,
    fmap: &HashMap<char, char>,
    point_values: &HashMap<char, usize>,
) -> Option<usize> {
    let mut stack = Vec::new();
    for char in line.chars() {
        if POSITIVES.contains(&char) {
            stack.push(char);
        } else {
            if stack.len() == 0 {
                return None;
            }
            let test = bmap[&char];
            if *stack.last().unwrap() != test {
                return None;
            }
            stack.pop();
        }
    }

    let mut score = 0;
    while let Some(el) = stack.pop() {
        let completion = fmap[&el];
        score = 5 * score + point_values[&completion];
    }
    Some(score)
}
