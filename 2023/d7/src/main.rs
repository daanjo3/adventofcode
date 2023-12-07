use std::collections::HashMap;

#[derive(PartialEq, Hash, Eq)]
struct Card(char);
impl Card {
    fn value(&self) -> i32 {
        if (self.0.is_numeric()) {
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

struct Hand {
    cards: Vec<Card>,
    bid: i32,
}

// get N of card
enum Score {
    FiveOak,
    FourOak,
    FullHouse,
    ThreeOak,
    TwoPair,
    OnePair,
    HighCard
}

fn get_score(cards: Vec<Card>) {
    let mut count_map: HashMap<&Card, usize> = cards.iter().fold(HashMap::new(), |m, test| {
        let count = cards.iter().filter(|c| c.0 == test.0).count();
        m.insert(test, count).unwrap();
        return m
    });
}

fn is_fiveoak(cards: Vec<Card>) -> bool {
    return cards.iter().all(|c| c.0 == cards[0].0);
}

fn is_fouroak(cards: Vec<Card>) -> bool {
    return cards.iter().any(|test| {
        return cards.iter().filter(|c| test.0 == c.0).count() == 4
    })
}

fn is_fullhouse() {

}

fn is_threeoak(cards: Vec<Card>) -> bool {
    return cards.iter().any(|test| {
        return cards.iter().filter(|c| test.0 == c.0).count() == 3
    })
}

fn is_twopair() {

}

fn is_onepair(cards: Vec<Card>) -> bool {
    return cards.iter().any(|test| {
        return cards.iter().filter(|c| test.0 == c.0).count() == 2
    })
}

fn parse_hand(line: &str) -> Hand {
    let mut parts = line.split_ascii_whitespace();
    let cards: Vec<Card> = parts.next().unwrap().chars().map(|c| Card(c)).collect();
    let bid: i32 = parts.next().unwrap().parse().unwrap();
    return Hand { cards, bid };
}

fn part_one(filename: &str) {
    let input = std::fs::read_to_string(filename).unwrap();
    let hands: Vec<Hand> = input.lines().map(|l| parse_hand(l)).collect();

}

fn main() {
    part_one("input/sample-input.txt");
}
