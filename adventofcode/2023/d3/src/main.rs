use std::{collections::HashMap, fs::read_to_string};
use regex::Regex;


struct Part {
    start: usize,
    end: usize,
    coord: Point
}

#[derive(PartialEq, Eq, Hash, Copy, Clone)]
struct Point {
    x: usize,
    y: usize
}

/**
 * Regex example from d2
 */
fn parse_parts(line: &str) -> Vec<regex::Match<'_>> {
    let regex_str = format!(r"(?<part_id>\d+)");
    let game_regex: Regex = Regex::new(&regex_str).unwrap();
    
    let mut parts = Vec::new();
    for cap in game_regex.captures_iter(line) {
        println!("{:?}", cap.name("part_id").unwrap());
        parts.push(cap.name("part_id").unwrap())
    }

    parts
}

fn min_x(x: usize) -> usize {
    return if x == 0 { 0 } else { (x-1).try_into().unwrap() };
}

fn max_x(x: usize, max_len: usize) -> usize {
    return if x == max_len-1 { x } else { x+1 };
}

/**
 * Readlines example from d1
 */
fn part_one(filename: &str) {
    let mut valid_map: HashMap<usize, Vec<Part>> = HashMap::new();
    for (y, line) in read_to_string(filename).unwrap().lines().enumerate() {
        
        for (x, elem) in line.chars().enumerate() {
            if elem == '.' || elem.is_numeric() {
                continue;
            }
            // encountered a symbol
            if y > 0 {
                if !valid_map.contains_key(&(y-1)) {
                    valid_map.insert(y-1, Vec::new());
                }
                valid_map.get_mut(&(y-1)).unwrap().push(Part {
                    start: min_x(x),
                    end: max_x(x, line.len()),
                    coord: Point { x, y }
                });
            }
            if !valid_map.contains_key(&y) {
                valid_map.insert(y, Vec::new());
            }
            valid_map.get_mut(&(y)).unwrap().push(Part {
                start: min_x(x),
                end: max_x(x, line.len()),
                coord: Point { x, y }
            });
            // This might create an entry for a line that doesn't exist but that's alright.
            if !valid_map.contains_key(&(y+1)) {
                valid_map.insert(y+1, Vec::new());
            }
            valid_map.get_mut(&(y+1)).unwrap().push(Part {
                start: min_x(x),
                end: max_x(x, line.len()),
                coord: Point { x, y }
            });
        }
    }

    let mut result: i32 = 0;
    for (y, line) in read_to_string(filename).unwrap().lines().enumerate() {
        let valid_parts = valid_map.get(&y).unwrap();
        for part in parse_parts(line).iter() {
            for valid_part in valid_parts {
                if is_adjacent(part, valid_part) {
                    println!("Part {} is valid", part.as_str());
                    result += part.as_str().parse::<i32>().unwrap();
                    break;
                }
            }
        }
    }

    println!("Result is {}", result);
    
}

fn part_two(filename: &str) {
    // Identify all gears are create a map of them
    let mut gear_reach_map: HashMap<usize, Vec<Part>> = HashMap::new();
    for (y, line) in read_to_string(filename).unwrap().lines().enumerate() {
        
        for (x, elem) in line.chars().enumerate() {
            if !gear_reach_map.contains_key(&y) {
                gear_reach_map.insert(y, Vec::new());
            }
            if elem != '*' {
                continue;
            }
            // encountered a symbol
            if y > 0 {
                if !gear_reach_map.contains_key(&(y-1)) {
                    gear_reach_map.insert(y-1, Vec::new());
                }
                gear_reach_map.get_mut(&(y-1)).unwrap().push(Part {
                    start: min_x(x),
                    end: max_x(x, line.len()),
                    coord: Point { x, y }
                });
            }
            gear_reach_map.get_mut(&(y)).unwrap().push(Part {
                start: min_x(x),
                end: max_x(x, line.len()),
                coord: Point { x, y }
            });
            // This might create an entry for a line that doesn't exist but that's alright.
            if !gear_reach_map.contains_key(&(y+1)) {
                gear_reach_map.insert(y+1, Vec::new());
            }
            gear_reach_map.get_mut(&(y+1)).unwrap().push(Part {
                start: min_x(x),
                end: max_x(x, line.len()),
                coord: Point { x, y }
            });
        }
    }

    // Check which parts are connected to which gears
    let mut gear_map: HashMap<Point, Vec<regex::Match<'_>>> = HashMap::new();
    let binding = read_to_string(filename).unwrap();
    for (y, line) in binding.lines().enumerate() {
        println!("Line: {}", y);
        let valid_parts = gear_reach_map.get(&y).unwrap();
        
        for part in parse_parts(line).iter() {

            for valid_part in valid_parts {
                if is_adjacent(part, valid_part) {
                    println!("Part {} is adjacent to gear ({}, {})", part.as_str(), valid_part.coord.x, valid_part.coord.y);
                    if !gear_map.contains_key(&valid_part.coord) {
                        gear_map.insert(valid_part.coord, Vec::new());
                    }
                    gear_map.get_mut(&valid_part.coord).unwrap().push(*part);
                }
            }
        }
    }

    // Check for which "gears" we can calculate the ratio
    let mut result: i32 = 0;
    for gear in gear_map.values() {
        if gear.len() != 2 {
            continue;
        }
        let p_val1 = gear[0].as_str().parse::<i32>().unwrap();
        let p_val2 = gear[1].as_str().parse::<i32>().unwrap();
        result += p_val1 * p_val2;
    }

    println!("Result is {}", result);
    
}

fn is_adjacent(part: &regex::Match<'_>, valid_part: &Part) -> bool {
    let part_len = part.end()-1 - part.start();
    let valid_len: usize = 3;
    let min_start = *[valid_part.start, part.start()].iter().min().unwrap();
    let max_end = *[valid_part.end, part.end()-1].iter().max().unwrap();
    
    min_start + part_len + valid_len > max_end
}

fn main() {
    // part_one("input/sample-input.txt");
    part_one("input/puzzle-input.txt");
    // part_two("input/sample-input.txt");
    part_two("input/puzzle-input.txt");
}
