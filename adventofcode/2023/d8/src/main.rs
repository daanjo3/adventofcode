use std::{fs::read_to_string, collections::HashMap};
use num::{integer::{gcd_lcm, self}, FromPrimitive};

struct MapPart {
    left: String,
    right: String
}
impl MapPart {
    fn direction(&self, instr: char) -> &str {
        if instr == 'L' {
            return self.left.as_str();
        } else {
            return self.right.as_str();
        }
    }
}

fn parse_map(input: String) -> (String, HashMap<String, MapPart>) {
    let mut data = input.split("\n\n");
    let instructions = data.next().unwrap();

    let mut map: HashMap<String, MapPart> = HashMap::new();
    for line in data.next().unwrap().lines() {
        let mut parts = line.split(" = ");
        let path_in = parts.next().unwrap();
        let mut path_out = parts.next().unwrap()
            .trim_matches(|c| c == '(' || c == ')')
            .split(", ");
        let left = path_out.next().unwrap().to_string();
        let right = path_out.next().unwrap().to_string();
        map.insert(path_in.to_string(), MapPart { 
            left, 
            right
        });
    }
    (instructions.to_string(), map)
}

fn calculate_distance(instructions: &String, map: &HashMap<String, MapPart>, start: &str, is_dest: fn (&String) -> bool) -> i32 {
    let mut location: String = start.to_string();
    let mut instr_i = 0;
    let instr_i_max = instructions.len() - 1;
    let mut num_steps = 0;

    while !is_dest(&location) {
        let instr = instructions.chars().nth(instr_i).unwrap();

        let map_part = map.get(&location).unwrap();
        location = map_part.direction(instr).to_string();
        num_steps += 1;

        if instr_i == instr_i_max {
            instr_i = 0;
        } else {
            instr_i += 1;
        }
    }
    num_steps
}

fn at_destination_simple(location: &String) -> bool {
    return location == "AAA";
}

fn part_one(filename: &str) -> i32 {
    let input = read_to_string(filename).unwrap();
    let (instructions, map) = parse_map(input);

    let distance = calculate_distance(&instructions, &map, "AAA", at_destination_simple);
    println!("Reached destination in {} steps.", distance);
    return distance;
}

fn at_destination(location: &String) -> bool {
    return location.ends_with('Z')
}

fn part_two(filename: &str) -> i64 {
    // Wanted to extract this to separate func but had to fight
    // with borrow checker.
    let input = read_to_string(filename).unwrap();
    let (instructions, map) = parse_map(input);

    let locations: Vec<String> = map.keys().filter(|k| k.ends_with('A')).map(|k| k.clone()).collect();
    let distances: Vec<i32> = locations.iter()
        .map(|l| calculate_distance(&instructions, &map, l, at_destination))
        .collect();

    println!("distances: {:?}", distances);
    let min_dist = distances.iter().fold(1, |acc, d| {
        let (_, lcm) = gcd_lcm(acc, FromPrimitive::from_i32(*d).unwrap());
        return lcm;
    });
    println!("min distance is {}", min_dist);
    min_dist
}

fn main() {
    // assert_eq!(part_one("input/sample-input1.txt"), 2);
    // assert_eq!(part_one("input/sample-input2.txt"), 6);
    // part_one("input/puzzle-input.txt");
    // assert_eq!(part_two("input/sample-input3.txt"), 6);
    part_two("input/puzzle-input.txt");
}
