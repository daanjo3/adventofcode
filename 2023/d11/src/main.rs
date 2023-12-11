use std::fs::read_to_string;

fn part_one(filename: &str) {
    let input = read_to_string(filename).unwrap();
    let lines = input.lines();
    let mut map: Vec<Vec<char>> = Vec::new();

    for line in lines {
        let mut row = Vec::new();
        for char in line.chars() {
            row.push(char)
        }
        map.push(row);
    }

    // expand rows
    for row in map {
        if row.iter().all(|c| *c == '.') {

        }
    }
    // expand columns

    // find galaxy locations

    // calculate pair paths

    // profit

}

fn main() {
    
}
