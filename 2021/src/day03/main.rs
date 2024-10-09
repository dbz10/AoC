use std::collections::HashMap;

use aoc_2021;

fn main() {
    let content = aoc_2021::read_content();
    let content: Vec<Vec<u32>> = content
        .lines()
        .map(|line| {
            line.chars()
                .map(|ch| ch.to_digit(10).expect("parse error"))
                .collect()
        })
        .collect();

    let part1 = part1(&content);
    println!("Part 1: {part1}");

    let part2 = part2(&content);
    println!("Part 2: {part2}");
}

fn part1(content: &Vec<Vec<u32>>) -> u32 {
    let val_length = content[0].len();
    let n = content.len();
    let mut counts = HashMap::new();

    for line in content.iter() {
        for (i, &v) in line.iter().enumerate() {
            if v == 1 {
                *counts.entry(i).or_insert(0) += 1;
            }
        }
    }

    let mut epsilon_binary = Vec::new();
    let mut gamma_binary = Vec::new();

    for i in 0..val_length {
        if counts[&i] > n / 2 {
            epsilon_binary.push('0');
            gamma_binary.push('1');
        } else {
            epsilon_binary.push('1');
            gamma_binary.push('0');
        }
    }

    let epsilon_rate =
        u32::from_str_radix(&epsilon_binary.iter().collect::<String>(), 2).expect("parse error");
    let gamma_rate =
        u32::from_str_radix(&gamma_binary.iter().collect::<String>(), 2).expect("parse error");

    epsilon_rate * gamma_rate
}

fn part2(content: &Vec<Vec<u32>>) -> u32 {
    let val_length = content[0].len();

    let mut c_epsilon = content.clone();
    for i in 0..val_length {
        let counts = build_counts(&c_epsilon);

        let n = c_epsilon.len();

        let target_value = if counts[&i] >= n - counts[&i] { 1 } else { 0 };
        if c_epsilon.len() == 1 {
            break;
        }
        c_epsilon = filter(c_epsilon, i, target_value);
    }
    let c_epsilon_str = &c_epsilon
        .first()
        .expect("should be one element")
        .iter()
        .map(|&z| z.to_string())
        .collect::<Vec<String>>()
        .join("");

    let mut c_gamma = content.clone();
    for i in 0..val_length {
        let counts = build_counts(&c_gamma);

        let n = c_gamma.len();

        let target_value = if counts[&i] >= n - counts[&i] { 0 } else { 1 };

        if c_gamma.len() == 1 {
            break;
        }
        c_gamma = filter(c_gamma, i, target_value);
    }
    let c_gamma_str = &c_gamma
        .first()
        .expect("should be one element")
        .iter()
        .map(|&z| z.to_string())
        .collect::<Vec<String>>()
        .join("");

    let epsilon_rate = u32::from_str_radix(c_epsilon_str, 2).expect("parse error");
    let gamma_rate = u32::from_str_radix(c_gamma_str, 2).expect("parse error");
    epsilon_rate * gamma_rate
}

fn filter<'a>(data: Vec<Vec<u32>>, position: usize, target_value: u32) -> Vec<Vec<u32>> {
    data.iter()
        .filter(|&line| line[position] == target_value)
        .map(|x| x.clone())
        .collect()
}

fn build_counts(data: &Vec<Vec<u32>>) -> HashMap<usize, usize> {
    let mut counts = HashMap::new();
    for i in 0..data[0].len() {
        counts.insert(i, 0);
    }

    for line in data.iter() {
        for (i, &v) in line.iter().enumerate() {
            if v == 1 {
                *counts.get_mut(&i).unwrap() += 1;
            }
        }
    }
    counts
}
