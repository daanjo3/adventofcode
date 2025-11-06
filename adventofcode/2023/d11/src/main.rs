use std::{fs::read_to_string, fmt::Display};

#[derive(Clone, Copy)]
struct Point {
    x: i64,
    y: i64
}
impl Display for Point {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.write_fmt(format_args!("({}, {})", self.x, self.y))
    }
}

fn part_two(filename: &str) {
    let input = read_to_string(filename).unwrap();
    let lines = input.lines();
    let mut map: Vec<Vec<char>> = Vec::new();

    for line in lines {
        let mut row = Vec::new();
        for char in line.chars() {
            row.push(char)
        }
        map.push(row);
    }

    // record expanded rows
    let mut y_expanded: Vec<i64> = Vec::new();
    for (i, row) in map.iter().enumerate() {
        if row.iter().all(|c| *c == '.') {
            y_expanded.push(i.try_into().unwrap());
        }
    }
    // record expanded columns
    let mut x_expanded: Vec<i64> = Vec::new();
    for col_i  in 0..map.get(0).unwrap().len() {
        let should_expand = map.iter()
            .map(|r| r.get(col_i).unwrap())
            .all(|c| *c == '.');
        if should_expand {
            x_expanded.push(col_i.try_into().unwrap());
        }
    }

    // find galaxy locations
    let mut galaxies: Vec<Point> = Vec::new();
    for (y, row) in map.iter().enumerate() {
        for (x, c) in row.iter().enumerate() {
            if *c == '#' {
                galaxies.push(Point { 
                    x: x.try_into().unwrap(), 
                    y: y.try_into().unwrap()
                })
            }
        }
    }

    // calculate pair paths
    let mut sum = 0;
    for ai in 0..galaxies.len() {
        let a = galaxies[ai];
        for bi in ai+1..galaxies.len() {
            let b = galaxies[bi];
            let distance = calculate_distance(a, b, &x_expanded, &y_expanded);
            sum += distance;
        }
    }

    println!("Distance sum of galaxy pairs paths is {}", sum);
}

fn calculate_distance(a: Point, b: Point, x_expanded: &Vec<i64>, y_expanded: &Vec<i64>) -> i64 {
    let x_passed: i64 = x_expanded.iter()
        .filter(|x| {
            if b.x > a.x {
                return a.x < **x && **x < b.x;
            }
            return b.x < **x && **x < a.x;
        })
        .count().try_into().unwrap();
    let y_passed: i64 = y_expanded.iter()
        .filter(|y| {
            if b.y > a.y {
                return a.y < **y && **y < b.y;
            }
            return b.y < **y && **y < a.y;
        })
        .count().try_into().unwrap();
    let x_dist = if a.x > b.x { a.x - b.x } else { b.x - a.x };
    let y_dist = if a.y > b.y { a.y - b.y } else { b.y - a.y };
    return x_dist + (x_passed * (1000000-1)) + y_dist + (y_passed * (1000000-1));
}

fn main() {
    // part_two("input/sample-input.txt");
    part_two("input/puzzle-input.txt");
}
