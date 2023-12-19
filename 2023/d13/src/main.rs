use std::{collections::HashMap, fs::read_to_string};
use itertools::Itertools;

fn main() {
    let filename = "input/sample-input.txt";
    let input = read_to_string(filename).unwrap();
    let arrays = input.split("\n\n");

    for array in arrays {
        println!("Array");
        let mut col_map: HashMap<usize, String> = HashMap::new();

        let row_vals: Vec<i32> = array.lines()
            .map(|r| r
                .replace(".", "0")
                .replace("#", "1"))
            .map(|r| {
                // Also process the row to correctly update the column map
                r.chars().enumerate().for_each(|(col_i, c)| {
                    if !col_map.contains_key(&col_i) {
                        col_map.insert(col_i, String::new());
                    }
                    col_map.get_mut(&col_i).unwrap().push(c);
                });
                // Parse the row string as binary
                return i32::from_str_radix(&r, 2).unwrap();
            })
            .collect();

        let col_vals: Vec<i32> = col_map
            .keys()
            .sorted()
            .map(|k| col_map.get(k).unwrap())
            .map(|v| i32::from_str_radix(&v, 2).unwrap())
            .collect::<Vec<i32>>();

        // Print values
        println!("Rows");
        row_vals.iter().for_each(|v| println!("row value: {}", v));
        println!("Columns");
        col_vals.iter().for_each(|v| println!("row value: {}", v));
    }
}
