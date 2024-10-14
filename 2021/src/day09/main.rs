use std::{collections::HashMap, str::FromStr};

use aoc_2021;

fn main() {
    let content = aoc_2021::read_content();
    let g = content.parse::<Grid>().expect("parse error");

    println!("(5,4) neighbors: {:?}", g.neighbors(5, 4));
    println!("(5,4) neighbors: {:?}", g.is_low_point(5, 4));

    println!("Part 1: {}", part1(&g));

    println!("Part 2: {}", part2(&g));
}

#[derive(Debug)]
struct Grid {
    lx: usize,
    ly: usize,
    height_map: HashMap<(usize, usize), u32>,
}

impl Grid {
    fn neighbors(&self, x: usize, y: usize) -> Vec<(usize, usize)> {
        let mut res = Vec::new();
        if x > 0 {
            res.push((x - 1, y))
        }
        if x < self.lx - 1 {
            res.push((x + 1, y))
        }
        if y > 0 {
            res.push((x, y - 1))
        }
        if y < self.ly - 1 {
            res.push((x, y + 1))
        }
        res
    }

    fn value_at(&self, x: usize, y: usize) -> &u32 {
        self.height_map.get(&(x, y)).expect("invalid location")
    }

    fn is_low_point(&self, x: usize, y: usize) -> bool {
        self.neighbors(x, y)
            .iter()
            .all(|&(nx, ny)| self.value_at(x, y) < self.value_at(nx, ny))
    }

    fn get_low_points(&self) -> Vec<(usize, usize)> {
        self.height_map
            .keys()
            .filter(|&&(x, y)| self.is_low_point(x, y))
            .cloned()
            .collect()
    }

    fn lowest_neighbor(&self, x: usize, y: usize) -> (usize, usize) {
        self.neighbors(x, y)
            .iter()
            .min_by_key(|(x, y)| self.value_at(*x, *y))
            .cloned()
            .expect("couldn't find a minimum neighbor")
    }
}

#[derive(Debug)]
struct FromStringErr(String);

impl FromStr for Grid {
    type Err = FromStringErr;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let lines: Vec<&str> = s.lines().collect();
        let ly = lines.len();
        let lx = lines.first().expect("expected at least one line").len();
        let mut height_map = HashMap::new();

        for (y, line) in lines.iter().enumerate() {
            for (x, ch) in line.chars().enumerate() {
                height_map.insert((x, y), ch.to_digit(10).expect("parse error"));
            }
        }

        Ok(Grid { lx, ly, height_map })
    }
}

fn part1(g: &Grid) -> u32 {
    g.get_low_points()
        .iter()
        .map(|&(x, y)| g.value_at(x, y) + 1)
        .sum()
}

fn part2(g: &Grid) -> usize {
    let mut basin_assignments: HashMap<(usize, usize), (usize, usize)> = HashMap::new();

    for &(x, y) in g
        .height_map
        .keys()
        .filter(|(x, y)| *g.value_at(*x, *y) != 9)
    {
        derive_cluster(&mut basin_assignments, (x, y), &g);
    }

    let mut basin_sizes: HashMap<(usize, usize), usize> = HashMap::new();
    for &v in basin_assignments.values() {
        *basin_sizes.entry(v).or_insert(0) += 1;
    }

    let mut sizes: Vec<usize> = basin_sizes.values().cloned().collect();
    sizes.sort_unstable();
    sizes.reverse();

    sizes[0] * sizes[1] * sizes[2]
}

fn derive_cluster(
    existing_clusters: &mut HashMap<(usize, usize), (usize, usize)>,
    point: (usize, usize),
    g: &Grid,
) -> (usize, usize) {
    if g.is_low_point(point.0, point.1) {
        existing_clusters.insert(point, point);
        return point;
    }
    match existing_clusters.get(&point) {
        Some(&cluster) => cluster,
        None => {
            let c = derive_cluster(existing_clusters, g.lowest_neighbor(point.0, point.1), g);
            existing_clusters.insert(point, c);
            return c;
        }
    }
}
