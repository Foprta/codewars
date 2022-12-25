// https://www.codewars.com/kata/52761ee4cffbc69732000738/train/rust

const GOOD_POWERS: [u32; 6] = [1, 2, 3, 3, 4, 10];
const EVIL_POWERS: [u32; 7] = [1, 2, 2, 2, 3, 5, 10];

fn good_vs_evil(good: &str, evil: &str) -> String {
    let good_power: u32 = good
        .split_whitespace()
        .enumerate()
        .fold(0, |acc, (idx, el)| {
            acc + el.parse::<u32>().expect("Only numbers are allowed") * GOOD_POWERS[idx]
        });

    let evil_power: u32 = evil
        .split_whitespace()
        .enumerate()
        .fold(0, |acc, (idx, el)| {
            acc + el.parse::<u32>().expect("Only numbers are allowed") * EVIL_POWERS[idx]
        });

    match good_power.cmp(&evil_power) {
        std::cmp::Ordering::Greater => "Battle Result: Good triumphs over Evil".to_string(),
        std::cmp::Ordering::Less => "Battle Result: Evil eradicates all trace of Good".to_string(),
        std::cmp::Ordering::Equal => "Battle Result: No victor on this battle field".to_string(),
    }
}

#[test]
fn returns_expected() {
    assert_eq!(
        good_vs_evil("0 0 0 0 0 10", "0 0 0 0 0 0 0"),
        "Battle Result: Good triumphs over Evil"
    );
    assert_eq!(
        good_vs_evil("0 0 0 0 0 0", "0 0 0 0 0 0 10"),
        "Battle Result: Evil eradicates all trace of Good"
    );
    assert_eq!(
        good_vs_evil("0 0 0 0 0 10", "0 0 0 0 0 0 10"),
        "Battle Result: No victor on this battle field"
    );
}
