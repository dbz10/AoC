use std::collections::HashMap;

use aoc_2021;

fn main() {
    let content = aoc_2021::read_content();
    let mut caves: HashMap<&str, Cave> = HashMap::new();

    // Instantiate caves
    content
        .lines()
        .map(|line| line.split_once('-').unwrap())
        .for_each(|(l, r)| {
            caves.entry(l).or_insert(Cave::new(l)).add_connection(r);
            caves.entry(r).or_insert(Cave::new(r)).add_connection(l);
        });

    println!("Part 1: {}", part1(caves.clone()));
    println!("Part 2: {}", part2(caves.clone()));
}

fn part1(caves: HashMap<&str, Cave>) -> u32 {
    go1(caves, &"start")
}

fn part2(caves: HashMap<&str, Cave>) -> u32 {
    go2(caves, &"start", false)
}

fn go1(mut caves: HashMap<&str, Cave>, cur: &str) -> u32 {
    if cur == "end" {
        return 1;
    }
    caves.get_mut(cur).unwrap().visit();
    let current_cave = &caves[cur];
    let nexts = current_cave
        .connects_to
        .iter()
        .filter(|connection| caves[&connection[..]].can_visit());

    nexts.map(|next| go1(caves.clone(), &next[..])).sum()
}

fn go2(mut caves: HashMap<&str, Cave>, cur: &str, double_small_consumed: bool) -> u32 {
    if cur == "end" {
        return 1;
    }
    caves.get_mut(cur).unwrap().visit();
    let current_cave = &caves[cur];
    current_cave
        .connects_to
        .iter()
        .map(|next| {
            let (can_visit, consume_double_small) =
                caves[&next[..]].can_visit_part2(double_small_consumed);

            if !can_visit {
                return 0;
            }

            go2(
                caves.clone(),
                &next[..],
                double_small_consumed || consume_double_small,
            )
        })
        .sum()
}

#[derive(Debug, Clone)]
enum CaveType {
    Small,
    Big,
}

#[derive(Debug, Clone)]
struct Cave {
    name: String,
    cave_type: CaveType,
    connects_to: Vec<String>,
    visit_count: u32,
}

impl Cave {
    fn new(name: &str) -> Self {
        let cave_type = if name.to_ascii_lowercase() == name {
            CaveType::Small
        } else {
            CaveType::Big
        };

        Cave {
            name: name.to_string(),
            cave_type,
            connects_to: vec![],
            visit_count: 0,
        }
    }

    fn add_connection(&mut self, other: &str) {
        self.connects_to.push(other.to_string());
    }

    fn visit(&mut self) {
        self.visit_count += 1;
    }

    fn can_visit(&self) -> bool {
        match self.cave_type {
            CaveType::Big => true,
            CaveType::Small => self.visit_count == 0,
        }
    }

    fn can_visit_part2(&self, double_small_consumed: bool) -> (bool, bool) {
        // Returns (can_visit, consume_double_small)
        match self.cave_type {
            CaveType::Big => (true, false),
            CaveType::Small => match &self.name[..] {
                "start" => (self.visit_count == 0, false),
                "end" => (self.visit_count == 0, false),
                _ => match (self.visit_count, double_small_consumed) {
                    (0, _) => (true, false),
                    (1, false) => (true, true),
                    _ => (false, false),
                },
            },
        }
    }
}
