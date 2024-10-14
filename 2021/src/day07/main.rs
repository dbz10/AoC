use aoc_2021;

fn main() {
    let content = aoc_2021::read_content();
    let positions: Vec<u32> = content
        .trim()
        .split(',')
        .map(|el| el.parse::<u32>().expect("parse error"))
        .collect();

    let p1 = fuel_consumption(&positions, |l, r| l.abs_diff(r));
    println!("Part 1: {p1}");

    let p2 = fuel_consumption(&positions, |l, r| {
        let d = l.abs_diff(r);
        d * (d + 1) / 2
    });
    println!("Part 2: {p2}");
}

fn fuel_consumption<F>(positions: &[u32], cost_fn: F) -> u32
where
    F: Fn(u32, u32) -> u32,
{
    let start = *positions
        .iter()
        .min()
        .expect("couldn't get a minimum value");
    let end = *positions.iter().max().expect("coulnd't get a max value");

    (start..=end)
        .map(|candidate_position| {
            positions
                .iter()
                .map(|&p| cost_fn(p, candidate_position))
                .sum()
        })
        .min()
        .expect("couldn't get min fuel consumption")
}
