// https://www.codewars.com/kata/54e320dcebe1e583250008fd/train/rust

const ZERO: u64 = 0;
const ONE: u64 = 1;
const TWO: u64 = ONE * 2;
const THREE: u64 = TWO * 3;
const FOUR: u64 = THREE * 4;
const FIVE: u64 = FOUR * 5;
const SIX: u64 = FIVE * 6;
const SEVEN: u64 = SIX * 7;
const EIGTH: u64 = SEVEN * 8;
const NINE: u64 = EIGTH * 9;
const A: u64 = NINE * 10;
const B: u64 = A * 11;
const C: u64 = B * 12;
const D: u64 = C * 13;
const E: u64 = D * 14;
const F: u64 = E * 15;
const G: u64 = F * 16;
const H: u64 = G * 17;
const I: u64 = H * 18;
const J: u64 = I * 19;
const K: u64 = J * 20;

const FACTORIAL_BY_INDEX: [u64; 21] = [
    ZERO, ONE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGTH, NINE, A, B, C, D, E, F, G, H, I, J, K,
];

const MULTIPLIER_BY_INDEX: [char; 21] = [
    '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I',
    'J', 'K',
];

fn dec2_fact_string(nb: u64) -> String {
    let max_idx = FACTORIAL_BY_INDEX
        .iter()
        .enumerate()
        .rev()
        .find_map(|x| {
            if x.1.le(&nb) {
                return Some(x.0);
            }

            return None;
        })
        .unwrap();

    let mut result: String = String::new();
    let mut buffer = nb;

    for idx in (1..=max_idx).rev() {
        let divider = buffer / FACTORIAL_BY_INDEX[idx];

        result.push(MULTIPLIER_BY_INDEX[divider as usize]);

        buffer -= FACTORIAL_BY_INDEX[idx] * divider;
    }

    return format!("{}0", result);
}

fn fact_string_2dec(s: String) -> u64 {
    let mut result: u64 = 0;

    for (idx, &item) in s.into_bytes().iter().rev().enumerate() {
        let mut multiplier = item as u64 - 48;

        if multiplier.gt(&9) {
            multiplier -= 7;
        }

        result += multiplier * FACTORIAL_BY_INDEX[idx];
    }

    return result;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn basics_dec2_fact_string() {
        assert_eq!(dec2_fact_string(0), "0");
        assert_eq!(dec2_fact_string(100), "40200");
        assert_eq!(dec2_fact_string(111), "42110");
        assert_eq!(dec2_fact_string(2982), "4041000");
        assert_eq!(dec2_fact_string(463), "341010");
        assert_eq!(dec2_fact_string(500000), "1331231100");
        assert_eq!(dec2_fact_string(38916800), "A7214031100");
    }

    #[test]
    fn basics_fact_string_2dec() {
        assert_eq!(fact_string_2dec(String::from("4041000")), 2982);
        assert_eq!(fact_string_2dec(String::from("341010")), 463);
        assert_eq!(fact_string_2dec(String::from("40200")), 100);
        assert_eq!(fact_string_2dec(String::from("42110")), 111);
        assert_eq!(fact_string_2dec(String::from("0")), 0);
        assert_eq!(fact_string_2dec(String::from("A7214031100")), 38916800);
        assert_eq!(fact_string_2dec(String::from("1331231100")), 500000);
    }
}
