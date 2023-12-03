use regex::Regex;

/**
 * Regex example from d2
 */
fn parse_cube(line: &str, color: &str) -> Vec<i32> {
    let regex_str = format!(r"(?<num_cubes>\d+) {}", color);
    let game_regex: Regex = Regex::new(&regex_str).unwrap();
    
    let mut nums = Vec::new();
    for cap in game_regex.captures_iter(line) {
        nums.push(cap.name("num_cubes").unwrap().as_str().parse::<i32>().unwrap())
    }

    nums
}