use std::fs::read_to_string;

struct Code {
    index: u32,
    number: u32
}

const NUMBER_STR: &'static [(u32, &'static str)] = &[
    (1, "one"),
    (2, "two"),
    (3, "three"),
    (4, "four"),
    (5, "five"),
    (6, "six"),
    (7, "seven"),
    (8, "eight"),
    (9, "nine")
];

fn extract_digits(line: &str) -> Vec<Code> {
    let mut result = Vec::new();
    
    for (i, c) in line.chars().enumerate() {
        if c.is_numeric() {
            result.push(Code { 
                index: u32::try_from(i).unwrap(),
                number: c.to_digit(10).unwrap() 
            })
        }
    }

    result
}

fn parse_numbers(line: &str) -> Vec<Code> {
    let mut result = Vec::new();

    for (num, numstr) in NUMBER_STR {
        let v: Vec<_> = line.match_indices(numstr).collect();
        for (i, _) in v.into_iter() {
            result.push(Code {
                index: u32::try_from(i).unwrap() + 1,
                number: *num
            })
        }
    }

    result
}

fn read_lines(filename: &str) -> Vec<u32> {
    let mut result = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        let mut nums = extract_digits(&line.to_string());
        let mut num2 = parse_numbers(line);
        nums.append(&mut num2);
        nums.sort_by_key(|code| code.index);
        let comb = nums[0].number.to_string() + &nums.last().unwrap().number.to_string();
        result.push(comb.parse::<u32>().unwrap());
    }

    result
}

fn main() {
    let calvalues = read_lines("input/puzzle-input.txt");
    let sum = calvalues.iter().sum::<u32>();
    println!("sum: {}", sum);
}
