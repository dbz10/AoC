use aoc_2021::read_content;

fn main() {
    let content = read_content();

    let instructions = content
        .lines()
        .map(|line| {
            let mut parts = line.split_ascii_whitespace();
            match (parts.next(), parts.next()) {
                (Some("forward"), Some(v)) => Move::Forward(v.parse::<u32>().expect("parse error")),
                (Some("up"), Some(v)) => Move::Up(v.parse::<u32>().expect("parse error")),
                (Some("down"), Some(v)) => Move::Down(v.parse::<u32>().expect("parse error")),
                _ => unreachable!(),
            }
        })
        .collect::<Vec<Move>>();

    let part1 = part1(&instructions);
    println!("Part 1: {part1}");

    let part2 = part2(&instructions);
    println!("Part 2: {part2}");
}

enum Move {
    Forward(u32),
    Down(u32),
    Up(u32),
}

#[derive(Debug)]
struct Submarine {
    x: u32,
    y: u32,
}

impl Submarine {
    fn new() -> Self {
        Submarine { x: 0, y: 0 }
    }

    fn go(&mut self, dir: &Move) {
        match dir {
            Move::Forward(dx) => self.x += dx,
            Move::Up(dy) => self.y -= dy,
            Move::Down(dy) => self.y += dy,
        }
    }
}

struct AimedSubmarine {
    x: u32,
    y: u32,
    aim: u32,
}

impl AimedSubmarine {
    fn new() -> Self {
        AimedSubmarine { x: 0, y: 0, aim: 0 }
    }

    fn go(&mut self, dir: &Move) {
        match dir {
            Move::Forward(d) => {
                self.x += d;
                self.y += d * self.aim;
            }
            Move::Up(d) => self.aim -= d,
            Move::Down(d) => self.aim += d,
        }
    }
}

fn part1(instructions: &Vec<Move>) -> u32 {
    let mut sub = Submarine::new();
    for instruction in instructions {
        sub.go(&instruction)
    }
    return sub.x * sub.y;
}

fn part2(instructions: &Vec<Move>) -> u32 {
    let mut sub = AimedSubmarine::new();

    for instruction in instructions {
        sub.go(&instruction)
    }
    return sub.x * sub.y;
}
