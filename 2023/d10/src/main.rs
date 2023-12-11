use std::{fs::read_to_string, collections::HashMap, slice::Iter, fmt::Display};

#[derive(PartialEq, Eq, Hash, Clone)]
struct Point {
    x: i32,
    y: i32
}
impl Point {
    fn from(x: i32, y: i32) -> Point {
        return Point {
            x,
            y
        }
    }
}
impl Display for Point {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.write_fmt(format_args!("({},{})", self.x, self.y))
    }
}

#[derive(PartialEq, Eq)]
struct Cell {
    point: Point,
    is_start: bool,
    opening: [Direction; 2]
}
impl Cell {
    fn get_neighbor<'a>(&'a self, map: &'a HashMap<Point, Cell>, d: &Direction) -> Option<&Cell> {
        match d {
            Direction::North => map.get(&Point { x: self.point.x, y: self.point.y-1 }),
            Direction::East => map.get(&Point::from(self.point.x+1, self.point.y)),
            Direction::South => map.get(&Point::from(self.point.x, self.point.y+1)),
            Direction::West => map.get(&Point::from(self.point.x-1, self.point.y)),
            _ => panic!("Direction none should never be provided to this function.")
        }
    }
    fn get_poss_neighbor<'a>(&'a self, map: &'a HashMap<Point, Cell>, filter_dir: bool) -> Vec<&Cell> {
        Direction::iter()
            .filter(|d| {
                if filter_dir {
                    return self.opening.contains(d)
                }
                true
            })
            .map(|d| (self.get_neighbor(map, d), d))
            .filter(|(o, _)| o.is_some())
            .map(|(c, d)| (c.unwrap(), d))
            .filter(|(c, d)| c.opening.contains(&d.opposite()) || c.is_start)
            .map(|(c, _)| c)
            .collect()
    }
    // fn get_next_neighbor<'a>(&'a self, map: &'a HashMap<Point, Cell>, origin: Point) -> Option<&&Cell> {
    //     let neighbors = self.get_poss_neighbor(map);
    //     let neighbor = neighbors.iter().find(|n| n.point != origin);
    //     neighbor
    // }
}

#[derive(Clone, Copy, PartialEq, Eq)]
enum Direction {
    None,
    North,
    East,
    South,
    West
}
impl Direction {
    pub fn iter() -> Iter<'static, Direction> {
        static DIRECTIONS: [Direction; 4] = [Direction::North, Direction::South, Direction::East, Direction::West];
        DIRECTIONS.iter()
    }

    fn opposite(&self) -> Direction {
        match self {
            Direction::None => Direction::None,
            Direction::North => Direction::South,
            Direction::East => Direction::West,
            Direction::South => Direction::North,
            Direction::West => Direction::East, 
        }
    }
}

fn part_one(filename: &str) -> i32 {
    let input = read_to_string(filename).unwrap();
    let lines = input.lines();
    let mut map: HashMap<Point, Cell> = HashMap::new();

    // Populate grid
    for (y, line) in lines.enumerate() {
        for (x, c) in line.chars().enumerate() {
            let point = Point { x: i32::try_from(x).unwrap(), y: i32::try_from(y).unwrap() };
            match c {
                '.' => map.insert(point.clone(), Cell { point: point.clone(), is_start: false, opening: [Direction::None, Direction::None] }),
                'F' => map.insert(point.clone(), Cell { point: point.clone(), is_start: false, opening: [Direction::South, Direction::East] }),
                '7' => map.insert(point.clone(), Cell { point: point.clone(), is_start: false, opening: [Direction::South, Direction::West] }),
                'J' => map.insert(point.clone(), Cell { point: point.clone(), is_start: false, opening: [Direction::North, Direction::West] }),
                'L' => map.insert(point.clone(), Cell { point: point.clone(), is_start: false, opening: [Direction::East, Direction::North] }),
                '-' => map.insert(point.clone(), Cell { point: point.clone(), is_start: false, opening: [Direction::West, Direction::East] }),
                '|' => map.insert(point.clone(), Cell { point: point.clone(), is_start: false, opening: [Direction::North, Direction::South] }),
                'S' => map.insert(point.clone(), Cell { point: point.clone(), is_start: true, opening: [Direction::None, Direction::None] }),
                _ => panic!("Should never be reached")
            };
        }
    }
    // Walk through map
    let start = map.values().find(|c| c.is_start).unwrap();
    let mut current = start;
    let mut previous = None;
    let mut count: i32 = 0;
    loop {
        println!("Current cell: {}", current.point);
        let neighbors = current.get_poss_neighbor(&map, previous.is_some());
        print!("Neighbors: ");
        neighbors.iter().map(|n| n.point.clone()).for_each(|p| print!("{}, ", p));
        print!("\n");
        match previous.as_ref() {
            None => {
                previous = Some(&current.point);
                current = neighbors.get(0).unwrap();
            }
            Some(&p) => {
                previous = Some(&current.point);
                current = neighbors.iter().find(|n| n.point != *p).unwrap()
            }
        }
        count += 1;
        if start == current {
            break
        }
    }

    println!("Path length = {}, target = {}", count, count / 2);
    count / 2

}

fn main() {
    part_one("input/sample-input.txt");
    part_one("input/puzzle-input.txt");
}
