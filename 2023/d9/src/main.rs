use std::fs::read_to_string;

enum Direction {
    Forward,
    Backward
}

fn part_one(filename: &str, direction: Direction) {
    let input = read_to_string(filename).unwrap();
    let lines = input.lines();

    let mut sum = 0;
    for line in lines {
        let mut sequences = Vec::new();
        let seq: Vec<i32> = line
            .split_ascii_whitespace()
            .map(|part| part.parse::<i32>().unwrap())
            .collect();
        sequences.push(seq);
        
        while !sequences.last().unwrap().iter().all(|n| *n == 0) {
            let mut new_seq: Vec<i32> = Vec::new();
            let last = sequences.last().unwrap();
            let mut prev = last[0];
            for i in 1..last.len() {
                let cur = last[i];
                let diff = cur - prev;
                new_seq.push(diff);
                prev = cur;
            }
            sequences.push(new_seq);
        }

        sequences.reverse();

        match direction {
            Direction::Forward => {
                let prediction = sequences.iter().map(|s| s.last().unwrap()).fold(0, |acc, n| acc + n);
                sum += prediction
            },
            Direction::Backward => {
                let mut prediction = 0;
                for sequence in sequences {
                    prediction = sequence.first().unwrap() - prediction;
                }
                sum += prediction
            }
        }
    }
    println!("Sum of predications is: {}", sum);
}

fn main() {
    part_one("input/sample-input.txt", Direction::Forward);
    part_one("input/puzzle-input.txt", Direction::Forward);
    part_one("input/sample-input.txt", Direction::Backward);
    part_one("input/puzzle-input.txt", Direction::Backward);
}
