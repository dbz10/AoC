use std::fmt::Display;

use aoc_2021;

fn main() {
    let content = aoc_2021::read_content();
    let mut content = content.split("\n\n");
    let input: Vec<u32> = content
        .next()
        .expect("expected to be able to read a line")
        .split(',')
        .map(|z| z.parse().unwrap())
        .collect();

    let bingos: Vec<Bingo> = content
        .map(|board| {
            let lines: Vec<Vec<u32>> = board
                .lines()
                .map(|l| {
                    l.split_ascii_whitespace()
                        .map(|el| el.parse::<u32>().unwrap())
                        .collect()
                })
                .collect();
            Bingo::new(lines)
        })
        .collect();

    let part1 = part1(&input, &mut bingos.clone()).expect("expected to find a solution");
    println!("Part 1: {part1}");

    let part2 = part2(&input, &mut bingos.clone()).expect("expected to find a solution");
    println!("Part 2: {part2}");
}

#[derive(Debug, Clone)]
struct Bingo {
    board: Vec<Vec<u32>>,
    marked: Vec<Vec<bool>>,
    sum_matched_vert: Vec<usize>,
    sum_matched_hor: Vec<usize>,
}

impl Bingo {
    fn new(board: Vec<Vec<u32>>) -> Self {
        let len = board.len();
        Bingo {
            board: board,
            marked: vec![vec![false; len]; len],
            sum_matched_vert: vec![0; len],
            sum_matched_hor: vec![0; len],
        }
    }

    fn play(&mut self, input: u32) {
        for (y, row) in self.board.iter().enumerate() {
            if let Some(x_found) = row.iter().position(|&r| r == input) {
                self.sum_matched_hor[y] += 1;
                self.sum_matched_vert[x_found] += 1;
                self.marked[y][x_found] = true;
                break;
            }
        }
    }

    fn check_bingo(&self) -> bool {
        let l = self.sum_matched_hor.len();
        (0..l).any(|i| self.sum_matched_hor[i] == l || self.sum_matched_vert[i] == l)
    }

    fn calculate_unmarked_sum(&self) -> u32 {
        self.board
            .iter()
            .zip(&self.marked)
            .flat_map(|(row, marked)| {
                row.iter()
                    .zip(marked)
                    .filter(|&(_, &m)| !m)
                    .map(|(&v, _)| v)
            })
            .sum()
    }
}

impl Display for Bingo {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for (y, line) in self.board.iter().enumerate() {
            for val in line.iter() {
                write!(f, "{} ", val)?;
            }
            write!(f, "{}", self.sum_matched_hor[y])?;
            write!(f, "\n")?;
        }
        writeln!(f, "{:?}", self.sum_matched_vert)?;
        Ok(())
    }
}

fn part1(input: &[u32], bingos: &mut Vec<Bingo>) -> Option<u32> {
    for &num in input {
        for i in 0..bingos.len() {
            let b = &mut bingos[i];
            b.play(num);
            if b.check_bingo() {
                return Some(b.calculate_unmarked_sum() * num);
            }
        }
    }
    None
}

fn part2(input: &[u32], bingos: &mut Vec<Bingo>) -> Option<u32> {
    for &num in input {
        let mut to_remove = Vec::new();
        let l = bingos.len();
        for i in 0..l {
            let b = &mut bingos[i];
            b.play(num);
            if b.check_bingo() {
                if l > 1 {
                    to_remove.push(i);
                } else {
                    return Some(b.calculate_unmarked_sum() * num);
                }
            }
        }

        to_remove.reverse();
        for r in to_remove {
            bingos.swap_remove(r);
        }
    }
    None
}
