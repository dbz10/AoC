use aoc2025;
use std::collections::{HashMap, HashSet};
fn main() {
    let content: Vec<Vec<char>> = aoc2025::read_content()
        .lines()
        .map(|l| l.chars().collect())
        .collect();

    let mut grid = HashMap::new();
    let mut start = (0, 0);
    let lx = content.iter().count() as i32;
    let ly = content[0].len() as i32;

    for (x, line) in content.iter().enumerate() {
        for (y, &ch) in line.iter().enumerate() {
            grid.insert((x as i32, y as i32), ch);
            if ch == 'S' {
                start = (0, y as i32);
            }
        }
    }

    println!("Part 1: {}", part1(&start, &grid, lx, ly));
    let mut memo = HashMap::new();
    println!("Part 2: {}", part2(&start, &grid, lx, ly, &mut memo));
}

fn part1(start: &(i32, i32), grid: &HashMap<(i32, i32), char>, lx: i32, ly: i32) -> i32 {
    let mut splitted = 0;
    let mut tachyons = HashSet::new();
    tachyons.insert(*start);

    for x in 0..lx {
        // println!("{:?}", tachyons);
        for y in 0..ly {
            if tachyons.contains(&(x, y)) {
                match grid.get(&(x + 1, y)) {
                    Some('.') => {
                        tachyons.insert((x + 1, y));
                    }
                    Some('^') => {
                        tachyons.insert((x + 1, y - 1));
                        tachyons.insert((x + 1, y + 1));
                        splitted += 1;
                    }
                    None => (),
                    _ => (),
                }
            }
        }
    }

    splitted
}

fn part2(
    start: &(i32, i32),
    grid: &HashMap<(i32, i32), char>,
    lx: i32,
    ly: i32,
    memo: &mut HashMap<(i32, i32), usize>,
) -> usize {
    if start.0 == lx - 1 {
        return 1;
    }

    if let Some(&v) = memo.get(start) {
        return v;
    }

    let new_start = (start.0 + 1, start.1);

    match grid.get(&new_start) {
        Some('.') => {
            let paths = part2(&new_start, grid, lx, ly, memo);
            memo.insert(*start, paths);
            paths
        }
        Some('^') => {
            let new_start_left = (start.0 + 1, start.1 - 1);
            let new_start_right = (start.0 + 1, start.1 + 1);
            let paths = part2(&new_start_left, grid, lx, ly, memo)
                + part2(&new_start_right, grid, lx, ly, memo);
            memo.insert(*start, paths);
            paths
        }
        _ => 0,
    }
}
