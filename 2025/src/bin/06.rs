use std::{
    iter::{zip, Iterator},
    str::FromStr,
};

use aoc2025;

fn main() {
    let content = aoc2025::read_content();
    let mut lines: Vec<&str> = content.lines().collect();

    let ops: Vec<Op> = lines
        .pop()
        .unwrap()
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    let inputs: Vec<Vec<i64>> = lines
        .iter()
        .map(|&line| {
            line.split_whitespace()
                .map(|el| el.parse().unwrap())
                .collect()
        })
        .collect();

    println!("Part 1: {}", part1(&inputs, &ops));
    println!("Part 2: {}", part2());
}

#[derive(Debug)]
enum Op {
    Add,
    Mult,
}
#[derive(Debug)]
struct ParseOpError;

impl FromStr for Op {
    type Err = ParseOpError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "+" => Ok(Op::Add),
            "*" => Ok(Op::Mult),
            _ => Err(ParseOpError),
        }
    }
}

fn part1(inputs: &[Vec<i64>], ops: &[Op]) -> i64 {
    let mut result = vec![];
    for op in ops.iter() {
        match op {
            Op::Add => result.push(0),
            Op::Mult => result.push(1),
        }
    }

    for (_, input_line) in inputs.iter().enumerate() {
        for (i, (op, el)) in zip(ops, input_line).enumerate() {
            match op {
                Op::Add => result[i] += el,
                Op::Mult => result[i] *= el,
            }
        }
    }
    result.iter().sum()
}

fn part2() -> i64 {
    // we need an entirely different approach here :/
    let content = aoc2025::read_content();
    let mut lines: Vec<&str> = content.lines().collect();

    let op_line_iterator: Vec<char> = lines.pop().unwrap().chars().rev().collect();
    let input_line_iterators: Vec<Vec<char>> =
        lines.iter().map(|&s| s.chars().rev().collect()).collect();

    let n_input_lines = input_line_iterators.len();

    let mut total = 0;
    // we now need to iterate backwards over all lines in sync. until we reach an op, we build up
    // an accumulator of numbers. when we reach an op, we combine them with the op and add the result to the total.

    let mut num_accumulator = vec![];
    let mut skip_next = false;
    for (i, op_element) in op_line_iterator.iter().enumerate() {
        if skip_next {
            skip_next = false;
            continue;
        }
        let parsed_num = (0..n_input_lines)
            .map(|z| input_line_iterators[z][i])
            .collect::<String>()
            .trim()
            .parse()
            .unwrap();

        num_accumulator.push(parsed_num);
        if let Ok(op) = op_element.to_string().parse::<Op>() {
            total += combine(&num_accumulator, &op);
            num_accumulator.clear();
            skip_next = true;
            continue;
        }
    }
    total
}

fn combine(nums: &[i64], op: &Op) -> i64 {
    match op {
        Op::Add => nums.iter().sum(),
        Op::Mult => nums.iter().fold(1, |acc, x| acc * x),
    }
}
