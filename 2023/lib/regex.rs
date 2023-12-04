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