use aoc2025;

fn main() {
    let content = aoc2025::read_content();

    let instructions: Vec<Instruction> = content
        .split_ascii_whitespace()
        .map(|s| Instruction::from_string(s))
        .collect();

    println!("Part 1: {}", part1(&instructions));
    println!("Part 2: {}", part2(&instructions));
}

fn part1(instructions: &Vec<Instruction>) -> i32 {
    let mut position: i32 = 50;
    let mut counter = 0;

    for i in instructions {
        match i {
            Instruction::Left(amount) => position -= amount,
            Instruction::Right(amount) => position += amount,
        }
        position = (position + 100) % 100;
        if position == 0 {
            counter += 1
        }
    }
    counter
}

fn part2(instructions: &Vec<Instruction>) -> i32 {
    let mut position: i32 = 50;
    let mut counter = 0;
    for i in instructions {
        match i {
            Instruction::Left(amount) => {
                for _ in 0..*amount {
                    position -= 1;
                    if position == 0 {
                        counter += 1;
                    }
                    if position < 0 {
                        position += 100;
                    }
                }
            }
            Instruction::Right(amount) => {
                for _ in 0..*amount {
                    position += 1;
                    if position == 100 {
                        counter += 1;
                        position = 0;
                    }
                }
            }
        }
    }

    counter
}

#[derive(Debug)]
enum Instruction {
    Left(i32),
    Right(i32),
}

impl Instruction {
    fn from_string(s: &str) -> Self {
        let chars: Vec<char> = s.chars().collect();
        let direction = chars[0];
        let amount: String = chars[1..].into_iter().collect();
        let amount = amount.parse().unwrap();

        match direction {
            'L' => Instruction::Left(amount),
            'R' => Instruction::Right(amount),
            _ => panic!(),
        }
    }
}
