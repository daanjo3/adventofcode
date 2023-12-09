use std::{collections::HashMap, cmp::Ordering, path::Display, fmt};

#[derive(PartialEq, Hash, Eq, Debug)]
struct Card(char);
impl Card {
    fn value(&self) -> i32 {
        if self.0.is_numeric() {
            return self.0.to_digit(10).unwrap().try_into().unwrap();
        }
        match self.0 {
            'A' => 14,
            'K' => 13,
            'Q' => 12,
            'J' => 11,
            'T' => 10,
            _ => panic!(),
        }
    }
}

#[derive(Debug)]
struct Hand {
    cards: Vec<Card>,
    score: Score,
    rank: i32,
    bid: i32,
}
impl fmt::Display for Hand {
    // This trait requires `fmt` with this exact signature.
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.cards.iter().map(|c| c.0).collect::<String>())
    }
}

#[derive(Debug, PartialEq, PartialOrd, Ord, Eq)]
enum Score {
    FiveOak,
    FourOak,
    FullHouse,
    ThreeOak,
    TwoPair,
    OnePair,
    HighCard(i32)
}
impl Score {
    fn is_equal(&self, score: &Score) -> bool {
        match self {
            Score::HighCard(_) => {
                match score {
                    Score::HighCard(_) => true,
                    _ => false
                }
            }
            _ => {
                self == score
            }
        }
    }
}

fn compare(hand_a: &Hand, hand_b: &Hand) -> Ordering {
    if hand_a.score.is_equal(&hand_b.score) {
        for i in 0..5 {
            if hand_a.cards[i] == hand_b.cards[i] {
                continue;
            }
            return hand_b.cards[i].value().cmp(&hand_a.cards[i].value());
        }
        return Ordering::Equal;
    }
    return hand_a.score.cmp(&hand_b.score);
}

fn calculate_score(cards: &Vec<Card>) -> Score {
    let count_map: HashMap<&Card, usize> = cards.iter().fold(HashMap::new(), |mut m, test| {
        let count = cards.iter().filter(|c| c.0 == test.0).count();
        m.insert(test, count);
        return m
    });
    let max = count_map.values().max().unwrap();
    match max {
        5 => Score::FiveOak,
        4 => Score::FourOak,
        3 => {
            if count_map.values().any(|c| *c == 2) {
                return Score::FullHouse
            }
            return Score::ThreeOak
        }
        2 => {
            if count_map.values().filter(|c| **c == 2).count() == 2 {
                return Score::TwoPair
            }
            return Score::OnePair
        }
        1 => Score::HighCard(cards.iter().map(|c| c.value()).max().unwrap()),
        _ => panic!("Condition should never be reached")
    }
}

fn parse_hand(line: &str) -> Hand {
    let mut parts = line.split_ascii_whitespace();
    let cards: Vec<Card> = parts.next().unwrap().chars().map(|c| Card(c)).collect();
    let bid: i32 = parts.next().unwrap().parse().unwrap();
    let score = calculate_score(&cards);
    return Hand { cards, score, rank: 0, bid };
}

fn part_one(filename: &str) {
    let input = std::fs::read_to_string(filename).unwrap();
    let mut hands: Vec<Hand> = input.lines().map(|l| parse_hand(l)).collect();

    hands.sort_by(|a, b| compare(a, b));
    hands.reverse();
    hands.iter_mut().enumerate().for_each(|(i, hand)| hand.rank = (i+1).try_into().unwrap());
    let total: i32 = hands.iter().fold(0, |acc,  h| acc + (h.rank * h.bid));
    hands.iter().for_each(|h| println!("{:?}", h));
    println!("Total score = {}", total);
}

fn main() {
    // part_one("input/sample-input.txt");
    part_one("input/puzzle-input.txt");
}
