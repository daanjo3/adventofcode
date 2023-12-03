use std::fs::read_to_string;

/**
 * Readlines example from d1
 */
fn read_lines(filename: &str) -> Vec<u32> {
    let mut result = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        ...
    }

    result
}
