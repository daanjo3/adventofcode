fn calculate_distance(time: i64, max_time: i64) -> i64 {
    let rem = max_time - time;
    return time * rem;
}

fn race(max_time: i64, winning_distance: i64) -> i32 {
    let mut margin = 0;
    for i in 1..max_time+1 {
        let distance = calculate_distance(i, max_time);
        if distance > winning_distance {
            margin += 1;
        }
    }
    println!("simulated race [time={}, distance={}]: margin={}", max_time, winning_distance, margin);
    margin
}

fn part_one() {
    // Sample input
    // const TIMES: [i32; 3] = [7, 15, 30];
    // const DISTANCES: [i32; 3] = [9, 40, 200];

    // Puzzle input
    const TIMES: [i32; 4] = [45, 98, 83, 73];
    const DISTANCES: [i32; 4] = [295, 1734, 1278, 1210];

    let mut total_margin = 1;
    for i in 0..4 {
        total_margin *= race(TIMES[i].try_into().unwrap(), DISTANCES[i].try_into().unwrap());
    }
    println!("Margin of error: {}", total_margin);
}

fn part_two() {
    // Sample input
    // let margin = race(71530, 940200);

    // Puzzle input
    let margin = race(45988373, 295173412781210);
    println!("Margin of error: {}", margin);
}

fn main() {
    // part_one()
    part_two()
}
