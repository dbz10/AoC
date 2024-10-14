use std::{collections::HashMap, ops::RangeInclusive};

use aoc_2021;

fn main() {
    let content = aoc_2021::read_content();
    let vents = parse_vents(&content);

    let part1 = part1(&vents);
    println!("Part 1: {part1}");

    let part2 = part2(&vents);
    println!("Part 2: {part2}");
}

#[derive(Debug, Clone, PartialEq, Eq, Hash)]
struct Point {
    x: usize,
    y: usize,
}

impl Point {
    fn are_hline(&self, other: &Point) -> bool {
        self.y == other.y
    }

    fn are_vline(&self, other: &Point) -> bool {
        self.x == other.x
    }

    fn are_diag(&self, other: &Point) -> bool {
        self.x.abs_diff(other.x) == self.y.abs_diff(other.y)
    }
}

#[derive(Debug, Clone)]
struct Vent {
    start: Point,
    end: Point,
}

impl Vent {
    fn from(s: &str) -> Vent {
        let mut split = s.split(" -> ");
        let mut start = split.next().unwrap().split(',');
        let mut end = split.next().unwrap().split(',');

        Vent {
            start: Point {
                x: start.next().unwrap().parse().expect("parse error"),
                y: start.next().unwrap().parse().expect("parse error"),
            },
            end: Point {
                x: end.next().unwrap().parse().expect("parse error"),
                y: end.next().unwrap().parse().expect("parse error"),
            },
        }
    }

    fn is_hv(&self) -> bool {
        self.start.are_hline(&self.end) || self.start.are_vline(&self.end)
    }

    fn is_hvd(&self) -> bool {
        self.is_hv() || self.start.are_diag(&self.end)
    }
}

fn parse_vents(s: &String) -> Vec<Vent> {
    s.lines().map(|l| Vent::from(l)).collect()
}

fn part1(vents: &[Vent]) -> u32 {
    let vents: Vec<&Vent> = vents.iter().filter(|&v| v.is_hv()).collect();
    count_intersections(&vents)
}

fn part2(vents: &[Vent]) -> u32 {
    let vents: Vec<&Vent> = vents.iter().filter(|&v| v.is_hvd()).collect();
    count_intersections(&vents)
}

fn range_inclusive_increasing<T: PartialOrd>(a: T, b: T) -> RangeInclusive<T> {
    if a <= b {
        a..=b
    } else {
        b..=a
    }
}

fn count_intersections(vents: &[&Vent]) -> u32 {
    let mut floor: HashMap<Point, u32> = HashMap::new();

    for &v in vents {
        if v.start == v.end {
            *floor.entry(v.start.clone()).or_insert(0) += 1;
        } else if v.start.are_hline(&v.end) {
            range_inclusive_increasing(v.start.x, v.end.x).for_each(|el| {
                *floor
                    .entry(Point {
                        x: el,
                        y: v.start.y,
                    })
                    .or_insert(0) += 1;
            });
        } else if v.start.are_vline(&v.end) {
            range_inclusive_increasing(v.start.y, v.end.y).for_each(|el| {
                *floor
                    .entry(Point {
                        x: v.start.x,
                        y: el,
                    })
                    .or_insert(0) += 1;
            });
        } else if v.start.are_diag(&v.end) {
            match [v.start.x <= v.end.x, v.start.y <= v.end.y] {
                [true, true] => {
                    let xr = v.start.x..=v.end.x;
                    let yr = v.start.y..=v.end.y;
                    xr.zip(yr)
                        .for_each(|(x, y)| *floor.entry(Point { x: x, y: y }).or_insert(0) += 1);
                }
                [true, false] => {
                    let xr = v.start.x..=v.end.x;
                    let yr = (v.end.y..=v.start.y).rev();
                    xr.zip(yr)
                        .for_each(|(x, y)| *floor.entry(Point { x: x, y: y }).or_insert(0) += 1);
                }
                [false, true] => {
                    let xr = (v.end.x..=v.start.x).rev();
                    let yr = v.start.y..=v.end.y;
                    xr.zip(yr)
                        .for_each(|(x, y)| *floor.entry(Point { x: x, y: y }).or_insert(0) += 1);
                }
                [false, false] => {
                    let xr = v.end.x..=v.start.x;
                    let yr = v.end.y..=v.start.y;
                    xr.zip(yr)
                        .for_each(|(x, y)| *floor.entry(Point { x: x, y: y }).or_insert(0) += 1);
                }
            }
        }
    }

    floor
        .into_values()
        .filter(|&count_vents| count_vents > 1)
        .count()
        .try_into()
        .expect("parse to u32 error")
}
