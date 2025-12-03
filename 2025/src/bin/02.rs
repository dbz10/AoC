use aoc2025;

fn main() {
    let content = aoc2025::read_content();

    let ranges: Vec<(i64, i64)> = content
        .trim()
        .split(',')
        .map(|x| {
            let (left, right) = x.split_once('-').unwrap();
            (left.parse().unwrap(), right.parse().unwrap())
        })
        .collect();

    println!("Part 1: {}", part1(&ranges));
    println!("Part 2: {}", part2(&ranges));
}

fn is_valid_p1(n: i64) -> bool {
    let s = n.to_string();
    let len = s.len();

    if len % 2 == 1 {
        return false;
    }

    let (left, right) = s.split_at(len / 2);
    left == right
}

fn is_valid_p2_inner(s: &str, chunksize: usize) -> bool {
    if s.len() % chunksize != 0 {
        return true;
    }

    let n_chunks = s.len() / chunksize;
    let pattern = &s[..chunksize];

    for i in 0..n_chunks {
        let segment = &s[i * chunksize..(i + 1) * chunksize];
        if segment != pattern {
            return true;
        }
    }
    false
}

fn is_valid_p2(n: i64) -> bool {
    let s = n.to_string();
    let len = s.len();

    for chunksize in 1..=(len / 2) {
        if !is_valid_p2_inner(&s, chunksize) {
            return true;
        }
    }
    false
}

fn part1(ranges: &[(i64, i64)]) -> i64 {
    ranges
        .iter()
        .flat_map(|&(start, end)| start..=end)
        .filter(|&n| is_valid_p1(n))
        .sum()
}

fn part2(ranges: &[(i64, i64)]) -> i64 {
    ranges
        .iter()
        .flat_map(|&(start, end)| start..=end)
        .filter(|&n| is_valid_p2(n))
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert!(is_valid_p2_inner("1111", 3));
        assert!(!is_valid_p2_inner("1111", 2));
        assert!(!is_valid_p2_inner("1111", 1));

        assert!(is_valid_p2_inner("123123", 2));
        assert!(!is_valid_p2_inner("123123", 3));
    }
}
