use std::fs::read_to_string;
use regex::Regex;

fn parse_cube(line: &str, color: &str) -> Vec<i32> {
    let regex_str = format!(r"(?<num_cubes>\d+) {}", color);
    let game_regex: Regex = Regex::new(&regex_str).unwrap();
    
    let mut nums = Vec::new();
    for cap in game_regex.captures_iter(line) {
        nums.push(cap.name("num_cubes").unwrap().as_str().parse::<i32>().unwrap())
    }

    nums
}

fn part_one(filename: &str) {
    let mut result = 0;
    const RED_TARGET: i32 = 12;
    const GREEN_TARGET: i32 = 13;
    const BLUE_TARGET: i32 = 14;

    for (i, line) in read_to_string(filename).unwrap().lines().enumerate() {
        let red_max = *parse_cube(line, "red").iter().max().unwrap();
        let green_max = *parse_cube(line, "green").iter().max().unwrap();
        let blue_max = *parse_cube(line, "blue").iter().max().unwrap();
        if red_max <= RED_TARGET && green_max <= GREEN_TARGET && blue_max <= BLUE_TARGET {
            result += i+1
        }
    }

    println!("result: {}", result);
}

fn part_two(filename: &str) {
    let mut result = 0;

    for line in read_to_string(filename).unwrap().lines() {
        let red_max = *parse_cube(line, "red").iter().max().unwrap();
        let green_max = *parse_cube(line, "green").iter().max().unwrap();
        let blue_max = *parse_cube(line, "blue").iter().max().unwrap();
        result += red_max * green_max * blue_max;
    }

    println!("result: {}", result);
}


fn main() {
    part_one("input/puzzle-input.txt");
    part_two("input/puzzle-input.txt");
}
