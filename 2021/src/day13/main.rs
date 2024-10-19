use std::collections::HashSet;

use aoc_2021;

fn main() {
    let content = aoc_2021::read_content();
    let (points_, instructions_) = content.split_once("\n\n").unwrap();

    let mut points = HashSet::new();
    for p in points_.lines() {
        let (x, y) = p.split_once(',').unwrap();
        points.insert((x.parse::<usize>().unwrap(), y.parse::<usize>().unwrap()));
    }
    let instructions: Vec<(&str, usize)> = instructions_
        .lines()
        .map(|line| {
            line.split_ascii_whitespace()
                .last()
                .unwrap()
                .split_once('=')
                .unwrap()
        })
        .map(|(dir, val)| (dir, val.parse::<usize>().unwrap()))
        .collect();

    println!("Part 1: {}", part1(&points, &instructions));
    part2(&points, &instructions);
}

fn part1(points: &HashSet<(usize, usize)>, instructions: &[(&str, usize)]) -> usize {
    foldem(points, &instructions[..1]).len()
}

fn part2(points: &HashSet<(usize, usize)>, instructions: &[(&str, usize)]) {
    let folded = foldem(points, instructions);
    let xmax = folded.iter().map(|&(x, y)| x).max().unwrap();
    let ymax = folded.iter().map(|&(x, y)| y).max().unwrap();

    for y in 0..=ymax {
        let line: String = (0..=xmax)
            .map(|x| if folded.contains(&(x, y)) { "#" } else { " " })
            .collect();
        println!("{line}");
    }
}
fn foldem(
    points: &HashSet<(usize, usize)>,
    instructions: &[(&str, usize)],
) -> HashSet<(usize, usize)> {
    let mut pout = points.clone();
    for &instruction in instructions {
        let mut pnew = HashSet::new();
        match instruction {
            ("x", fold) => {
                for (x0, y0) in pout {
                    if x0 > fold {
                        pnew.insert((2 * fold - x0, y0));
                    } else {
                        pnew.insert((x0, y0));
                    }
                }
            }
            ("y", fold) => {
                for (x0, y0) in pout {
                    if y0 > fold {
                        pnew.insert((x0, 2 * fold - y0));
                    } else {
                        pnew.insert((x0, y0));
                    }
                }
            }
            _ => unreachable!(),
        }
        pout = pnew.clone();
    }
    pout
}
