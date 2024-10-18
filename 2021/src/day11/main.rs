use std::collections::{HashMap, HashSet};

use aoc_2021;

fn main() {
    let content = aoc_2021::read_content();
    let g = OctopusGrid::from(&content);

    println!("Part 1: {}", part1(g.clone(), 100));

    match part2(g.clone()) {
        Some(v) => println!("Part 2: {v}"),
        None => println!("No result found... yet..."),
    };
}

#[derive(Debug, Clone)]
struct OctopusGrid {
    lx: isize,
    ly: isize,
    energy_map: HashMap<(isize, isize), u32>,
    flashes: usize,
}

impl OctopusGrid {
    fn from(s: &str) -> Self {
        let lines: Vec<&str> = s.lines().collect();
        let ly = lines.len().try_into().unwrap();
        let lx = lines
            .first()
            .expect("expected at least one line")
            .len()
            .try_into()
            .unwrap();
        let mut energy_map = HashMap::new();

        for (y, line) in lines.iter().enumerate() {
            for (x, ch) in line.chars().enumerate() {
                energy_map.insert(
                    (x.try_into().unwrap(), y.try_into().unwrap()),
                    ch.to_digit(10).expect("parse error"),
                );
            }
        }

        OctopusGrid {
            lx,
            ly,
            energy_map,
            flashes: 0,
        }
    }

    fn inbounds(&self, x: isize, y: isize) -> bool {
        x >= 0 && x < self.lx.try_into().unwrap() && y >= 0 && y < self.ly.try_into().unwrap()
    }

    fn neighbors(&self, x: isize, y: isize) -> Vec<(isize, isize)> {
        let mut res = Vec::new();
        for dx in [-1, 0, 1] {
            for dy in [-1, 0, 1] {
                if self.inbounds(x + dx, y + dy) && !(dx == 0 && dy == 0) {
                    res.push((x + dx, y + dy));
                }
            }
        }

        res
    }

    fn value_at(&self, x: isize, y: isize) -> u32 {
        *self
            .energy_map
            .get(&(x, y))
            .expect(&format!("No key at {x}, {y}"))
    }

    fn step(&mut self) {
        for energy in self.energy_map.values_mut() {
            *energy += 1;
        }

        while self.energy_map.values().any(|&l| l > 9) {
            let mut energy_propagation = HashMap::new();
            let mut to_zero = HashSet::new();

            for (&(x, y), _) in self.energy_map.iter().filter(|(_, &level)| level > 9) {
                to_zero.insert((x, y));
                for (nx, ny) in self.neighbors(x, y) {
                    if self.value_at(nx, ny) > 0 {
                        *energy_propagation.entry((nx, ny)).or_insert(0) += 1;
                    }
                }
            }

            for (&(x, y), &amt) in energy_propagation.iter() {
                *self.energy_map.get_mut(&(x, y)).unwrap() += amt;
            }
            for (x, y) in to_zero {
                self.energy_map.insert((x, y), 0);
                self.flashes += 1;
            }
        }
    }
}

fn part1(mut g: OctopusGrid, n_steps: usize) -> usize {
    for _ in 0..n_steps {
        g.step();
    }
    g.flashes
}

fn part2(mut g: OctopusGrid) -> Option<usize> {
    // brute force? if not...
    for step in 1..1_000_000 {
        g.step();
        if g.energy_map.values().all(|&v| v == 0) {
            return Some(step);
        }
    }
    None
}
