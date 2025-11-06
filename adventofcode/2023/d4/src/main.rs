use std::{fs::read_to_string, collections::HashMap};
use regex::Regex;

/**
 * Regex example from d2
 */
fn parse_numbers(line: &str) -> Vec<i32> {
    let number_regex: Regex = Regex::new(r"(\b(?<numbers>\d{1,2})\b)").unwrap();
    
    let mut result = Vec::new();
    for cap in number_regex.captures_iter(line) {
        result.push(cap.name("numbers").unwrap().as_str().parse::<i32>().unwrap())
    }

    result
}

fn get_num_win(line: &str) -> usize {
    let number_str = line.split(':').collect::<Vec<_>>()[1];
    let split_line = number_str.split('|').collect::<Vec<_>>();

    let winning = parse_numbers(split_line[0]);
    let entry = parse_numbers(split_line[1]);

    let num_win = entry.iter().filter(|n| winning.contains(n)).collect::<Vec<_>>().len();
    
    num_win
}
/**
 * Readlines example from d1
 */
fn part_one(filename: &str) {
    let mut result = 0;

    for (i, line) in read_to_string(filename).unwrap().lines().enumerate() {
        let num_win = get_num_win(line);
        if num_win > 0 {
            let score = 2i32.pow((num_win - 1).try_into().unwrap());
            result += score
            // println!("game: {}\nwinning: {:?}\nentry: {:?}\nscore: {}", i+1, winning, entry, score);
        }
    }

    println!("Total score is: {}", result)
}

fn part_two(filename: &str) {
    let mut num_card_map: HashMap<i32, i32> = HashMap::new();

    // yes big dirty
    for (i, _line) in read_to_string(filename).unwrap().lines().enumerate() {
        num_card_map.insert((i+1).try_into().unwrap(), 1);
    }

    for (i, line) in read_to_string(filename).unwrap().lines().enumerate() {
        let cur_card: i32 = (i+1).try_into().unwrap();
        let num_card = *num_card_map.get(&cur_card).unwrap();
        let num_win: i32 = get_num_win(line).try_into().unwrap();
        // println!("game {}: amount winning {} * num cards {}", i+1, num_win, num_card);

        for new_card in cur_card+1..cur_card+1+num_win {
            // println!("Adding {} cards for game {}", num_card, new_card);
            *num_card_map.get_mut(&new_card).unwrap() += num_card;
        }
    }

    // for (i, val) in num_card_map.values().enumerate() {
        // println!("game {}: cards = {}", i, val)
    // }
    let total_cards = num_card_map.values().fold(0, |acc,  &e| acc + e);

    println!("Total score is: {}", total_cards)
}

fn main() {
    // part_one("input/sample-input.txt");
    // part_one("input/puzzle-input.txt");
    part_two("input/sample-input.txt");
    part_two("input/puzzle-input.txt");
}
