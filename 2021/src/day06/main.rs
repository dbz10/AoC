use std::collections::HashMap;

use aoc_2021;
fn main() {
    let content = aoc_2021::read_content();
    let content = content
        .trim()
        .split(',')
        .map(|c| c.parse::<u32>().expect("parse error"))
        .collect::<Vec<u32>>();

    let lanternfish = LanternFishPopulation::from(&content);

    let p1 = part1(&lanternfish, 80);
    println!("Part 1: {p1}");

    let p2 = part1(&lanternfish, 256);
    println!("Part 2: {p2}");
}

#[derive(Debug, Clone)]
struct LanternFishPopulation {
    census: HashMap<u32, usize>,
}

impl LanternFishPopulation {
    fn from(arr: &[u32]) -> Self {
        let mut census = HashMap::new();
        for &v in arr {
            *census.entry(v).or_insert(0) += 1;
        }
        LanternFishPopulation { census }
    }

    fn step(&self) -> LanternFishPopulation {
        let mut census = HashMap::new();

        for age in 1..=8 {
            if let Some(&count) = self.census.get(&age) {
                census.insert(age - 1, count);
            }
        }

        if let Some(&count) = self.census.get(&0) {
            census.insert(8, count);
            *census.entry(6).or_insert(0) += count;
        }

        LanternFishPopulation { census }
    }

    fn count_lanternfish(&self) -> usize {
        self.census.values().sum()
    }
}

fn part1(l: &LanternFishPopulation, n: u32) -> usize {
    let mut li = l.clone();
    for _ in 0..n {
        li = li.step();
    }
    li.count_lanternfish()
}
