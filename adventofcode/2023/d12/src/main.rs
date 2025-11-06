use std::fs::read_to_string;

use itertools::{Itertools, repeat_n};

fn part_one(filename: &str) {
    let input = read_to_string(filename).unwrap();
    let lines = input.lines();

    let mut sum = 0;
    let mut progress = 1;
    let total = 1000;
    for line in lines {
        println!("[{}/{}]: {}", progress, total, line);
        let mut parts = line.split_ascii_whitespace();
        let mask = parts.next().unwrap();
        let springarr_str = parts.next().unwrap();
        let springarrs: Vec<usize> = springarr_str.split(",").map(|c| c.parse::<usize>().unwrap()).collect();
        sum += brute_force(mask, springarrs);
        progress += 1;
    }

    println!("The total sum is {}", sum)
}

fn brute_force(mask: &str, springarrs: Vec<usize>) -> i32 {
    // println!("Brute forcing pattern {}", mask);
    let options = vec!['#', '.'];
    let size = mask.len();
    let sum_springs = springarrs.iter().sum();
    let mut count = 0;
    
    for arr in repeat_n(options, size).multi_cartesian_product() {
        // 1. Check if num '#' is equal to sized combined
        let num_spring = arr.iter().filter(|c| **c == '#').count();
        let arr_str: String = arr.iter().join("");
        // println!("Attempting pattern {}, arrangment {:?}", arr_str, springarrs);

        if num_spring != sum_springs {
            // println!("{} failed: not enough spring", arr_str);
            continue;
        }
        // 2. Check if string fits in the mask
        let mut is_match = true;
        for (i, char) in mask.chars().enumerate() {
            if char == '?' {
                // println!("Continuing because of '?'");
                continue;
            }
            if arr[i] != char {
                is_match = false;
                // println!("Mismatch {} and {}", arr[i], char);
                break;
            }
        }
        if !is_match {
            // println!("{} failed: does not match mask", arr_str);
            continue;
        }
        // 3. Check if string is valid (gaps between sizes)
        is_match = arr_str.split('.')
            .map(|g| g.len())
            .filter(|gl| gl > &0)
            .eq(springarrs.clone());
        
        if !is_match {
            // println!("{} failed: no valid arrangement", arr_str);
            continue;
        }
        count += 1
    }
    count
}

fn main() {
    part_one("input/puzzle-input.txt")
}
