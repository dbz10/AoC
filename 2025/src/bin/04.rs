use std::collections::HashMap;

use aoc2025;

fn main() {
    let content = aoc2025::read_content();

    let lx = content.lines().count();
    let ly = content.lines().next().unwrap().len();

    let mut positions: HashMap<(i32, i32), i32> = HashMap::new();
    for (i, line) in content.lines().enumerate() {
        for (j, ch) in line.chars().enumerate() {
            if ch == '@' {
                positions.insert((i as i32, j as i32), 1);
            }
        }
    }

    println!("Part 1: {}", part1(&positions, lx, ly).len());
    println!("Part 2: {}", part2(positions, lx, ly));
}

fn part1(positions: &HashMap<(i32, i32), i32>, lx: usize, ly: usize) -> Vec<(i32, i32)> {
    let mut removable = Vec::new();
    for x in 0..lx {
        for y in 0..ly {
            let x = x as i32;
            let y = y as i32;
            if !positions.contains_key(&(x, y)) {
                continue;
            } else {
                let mut neighbors = -1;
                for dx in -1..=1 {
                    for dy in -1..=1 {
                        if positions.contains_key(&(x + dx, y + dy)) {
                            neighbors += 1;
                        }
                    }
                }
                if neighbors < 4 {
                    // println!("{x}, {y}");
                    removable.push((x, y));
                }
            }
        }
    }
    removable
}

fn part2(mut positions: HashMap<(i32, i32), i32>, lx: usize, ly: usize) -> usize {
    let mut count_removed = 1;
    let mut total_removed = 0;
    while count_removed > 0 {
        let removable = part1(&positions, lx, ly);
        count_removed = removable.len();
        total_removed += count_removed;
        for (x, y) in removable.into_iter() {
            positions.remove(&(x, y));
        }
    }

    total_removed
}
