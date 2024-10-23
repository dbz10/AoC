use std::{
    collections::{BinaryHeap, HashMap},
    usize,
};

use aoc_2021;

fn main() {
    let content = aoc_2021::read_content();
    let map = ValueGrid::parse(content.clone());
    println!("Part 1: {}", shortest_path(map));
    let map = ValueGrid::parse_fivex(content.clone());
    println!("Part 1: {}", shortest_path(map));
}

#[derive(Debug, Clone, PartialEq, Eq, PartialOrd, Hash, Copy)]
struct Point {
    x: isize,
    y: isize,
}

impl Point {
    fn new(z: &(isize, isize)) -> Self {
        Point { x: z.0, y: z.1 }
    }
}

struct ValueGrid {
    lx: usize,
    ly: usize,
    v: HashMap<Point, usize>,
}

impl ValueGrid {
    fn parse(s: String) -> Self {
        let mut v = HashMap::new();
        let data: Vec<Vec<char>> = s.lines().map(|l| l.chars().collect()).collect();
        let lx = data[0].len().try_into().unwrap();
        let ly = data.len();

        for (y, line) in data.iter().enumerate() {
            for (x, ch) in line.iter().enumerate() {
                let x = x.try_into().unwrap();
                let y = y.try_into().unwrap();
                v.insert(Point { x, y }, ch.to_digit(10).unwrap().try_into().unwrap());
            }
        }
        ValueGrid { lx, ly, v }
    }

    fn parse_fivex(s: String) -> Self {
        let mut v = HashMap::new();
        let data: Vec<Vec<char>> = s.lines().map(|l| l.chars().collect()).collect();
        let lx_base = data[0].len();
        let ly_base = data.len();

        for (y, line) in data.iter().enumerate() {
            for (x, ch) in line.iter().enumerate() {
                for vx in 0..5 {
                    for vy in 0..5 {
                        let cost: usize = ch.to_digit(10).unwrap().try_into().unwrap();
                        let cost = (cost + vx + vy - 1) % 9 + 1;
                        let x: isize = (x + vx * lx_base).try_into().unwrap();
                        let y: isize = (y + vy * ly_base).try_into().unwrap();
                        v.insert(Point { x, y }, cost);
                    }
                }
            }
        }
        ValueGrid {
            lx: lx_base * 5,
            ly: ly_base * 5,
            v,
        }
    }

    fn neighbors(&self, p: &Point) -> Vec<Point> {
        let mut n: Vec<Point> = vec![
            (p.x + 1, p.y),
            (p.x - 1, p.y),
            (p.x, p.y + 1),
            (p.x, p.y - 1),
        ]
        .iter()
        .map(Point::new)
        .collect();

        n.retain(|p| {
            p.x >= 0
                && p.x < self.lx.try_into().unwrap()
                && p.y >= 0
                && p.y < self.ly.try_into().unwrap()
        });
        n
    }

    fn cost(&self, p: &Point) -> Option<usize> {
        self.v.get(p).copied()
    }
}

#[derive(Debug, Clone, PartialEq, Eq, Copy)]
struct State {
    point: Point,
    cost: usize,
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        other.cost.cmp(&self.cost) // reversing here to turn behavior under BinaryHeap from max heap into min heap
    }
}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}

fn shortest_path(map: ValueGrid) -> usize {
    let mut front = BinaryHeap::new(); // A priority queue
    let mut dist = HashMap::new();

    let start = Point { x: 0, y: 0 };
    let end = Point {
        x: (map.lx - 1).try_into().unwrap(),
        y: (map.ly - 1).try_into().unwrap(),
    };

    front.push(State {
        point: start,
        cost: 0,
    });
    dist.insert(start, 0);

    while let Some(State { point, cost }) = front.pop() {
        if point == end {
            return cost;
        }

        if cost > *dist.get(&point).unwrap_or(&usize::MAX) {
            // Then we already found a "shorter" path to this point
            continue;
        }

        for neighbor in map.neighbors(&point) {
            let next = State {
                point: neighbor,
                cost: cost + map.cost(&neighbor).unwrap(),
            };
            if next.cost < *dist.get(&next.point).unwrap_or(&usize::MAX) {
                front.push(next);
                dist.insert(next.point, next.cost);
            }
        }
    }
    panic!("Never found a way to the end?")
}
