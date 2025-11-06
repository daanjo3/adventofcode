use regex::Regex;
use std::fs::read_to_string;

/**
 * Regex example from d4
 */
fn parse_numbers(line: &str) -> Vec<i64> {
    let number_regex: Regex = Regex::new(r"(\b(?<numbers>\d+)\b)").unwrap();
    
    let mut result = Vec::new();
    for cap in number_regex.captures_iter(line) {
        result.push(cap.name("numbers").unwrap().as_str().parse::<i64>().unwrap())
    }

    result
}

fn parse_mapping(line: &str) -> Mapping {
    let numbers = parse_numbers(line);
    assert!(numbers.len() == 3);
    return Mapping {
        low: numbers[1],
        high: numbers[1] + numbers[2],
        change: numbers[0] - numbers[1]
    }
}

#[derive(Debug, Clone)]
struct Mapping {
    low: i64,
    high: i64, // non inclusive
    change: i64
}

/**
 * Readlines example from d1
 */
fn part_one(filename: &str) {
    let mut seeds = Vec::new();
    let mut maps = Vec::new();

    let mut mapping = Vec::new();
    for line in read_to_string(filename).unwrap().lines() {
        if line.starts_with("seeds") {
            seeds = parse_numbers(line);
            continue;
        }
        if line.is_empty() {
            if !mapping.is_empty() {
                maps.push(mapping.clone())
            }
            mapping.clear();
            continue;
        }
        if line.contains("map") {
            // println!("Parsing {}", line);
            continue;
        }
        let map = parse_mapping(line);
        // println!("{}-{} -> {}", map.low, map.high, map.change);
        mapping.push(map);
    }
    if !mapping.is_empty() {
        maps.push(mapping)
    }

    // pass seeds through each map
    let mut seeds_result = Vec::new();
    for seed_start in seeds {
        let mut seed = seed_start;
        println!("Mapping seed: {}", seed);
        
        for map in &maps {
            let mut is_mapped = false;
            for mapping in map {

                if seed >= mapping.low && seed < mapping.high {
                    println!("{} -> {}", seed, seed + mapping.change);
                    seed += mapping.change;
                    is_mapped = true;
                    break;
                }

            }
            if !is_mapped {
                println!("{} -> {}", seed, seed);
            }
        }
        
        seeds_result.push(seed);
    }

    let binding = seeds_result.clone();
    let lowest = binding.iter().min().unwrap();
    println!("Locations are:\n{:?}", seeds_result);
    println!("Lowest value is: {}", lowest);

}

// TODO
fn part_two(filename: &str) {
    let mut seeds = Vec::new();
    let mut maps = Vec::new();

    let mut mapping = Vec::new();
    for line in read_to_string(filename).unwrap().lines() {
        if line.starts_with("seeds") {
            seeds = parse_numbers(line);
            continue;
        }
        if line.is_empty() {
            if !mapping.is_empty() {
                maps.push(mapping.clone())
            }
            mapping.clear();
            continue;
        }
        if line.contains("map") {
            // println!("Parsing {}", line);
            continue;
        }
        let map = parse_mapping(line);
        // println!("{}-{} -> {}", map.low, map.high, map.change);
        mapping.push(map);
    }
    if !mapping.is_empty() {
        maps.push(mapping)
    }

    // pass seeds through each map
    let mut seeds_result = Vec::new();
    for seed_start in seeds {
        let mut seed = seed_start;
        println!("Mapping seed: {}", seed);
        
        for map in &maps {
            let mut is_mapped = false;
            for mapping in map {

                if seed >= mapping.low && seed < mapping.high {
                    println!("{} -> {}", seed, seed + mapping.change);
                    seed += mapping.change;
                    is_mapped = true;
                    break;
                }

            }
            if !is_mapped {
                println!("{} -> {}", seed, seed);
            }
        }
        
        seeds_result.push(seed);
    }

    let binding = seeds_result.clone();
    let lowest = binding.iter().min().unwrap();
    println!("Locations are:\n{:?}", seeds_result);
    println!("Lowest value is: {}", lowest);

}

fn main() {
    // part_one("input/sample-input.txt");
    part_one("input/puzzle-input.txt");
}
