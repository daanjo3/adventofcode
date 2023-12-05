use regex::Regex;

/**
 * Regex example from d2
 */
fn parse_regex(line: &str, color: &str) -> Vec<i32> {
    let regex_str = format!(r"(?<num_cubes>\d+) {}", color);
    let regex: Regex = Regex::new(&regex_str).unwrap();
    
    let mut nums = Vec::new();
    for cap in regex.captures_iter(line) {
        nums.push(cap.name("num_cubes").unwrap().as_str().parse::<i32>().unwrap())
    }

    nums
}

/**
 * Regex example from d4
 */
fn parse_numbers(line: &str) -> Vec<i32> {
    let number_regex: Regex = Regex::new(r"(\b(?<numbers>\d{1,2})\b)").unwrap();
    
    let mut result = Vec::new();
    for cap in number_regex.captures_iter(line) {
        result.push(cap.name("numbers").unwrap().as_str().parse::<i32>().unwrap())
    }

    result
}