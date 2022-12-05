// https://www.codewars.com/kata/5bb904724c47249b10000131/train/rust
pub fn points(games: &[String]) -> u32 {
    games.iter().fold(0, |res, item| {
        let left_number: u32 = item.chars().nth(0).unwrap_or('0').to_digit(10).unwrap_or(0);
        let right_number: u32 = item.chars().nth(2).unwrap_or('0').to_digit(10).unwrap_or(0);

        if left_number > right_number {
            return res + 3;
        } else if left_number == right_number {
            return res + 1;
        }

        return res;
    })
}

#[cfg(test)]
mod tests {
    use super::points;

    const ERR_MSG: &str = "\nYour result (left) did not match the expected output (right)";

    fn do_fixed_test(e: &[&str], expected: u32) {
        let a = &e.iter().map(|x| x.to_string()).collect::<Vec<_>>();
        assert_eq!(points(a), expected, "{ERR_MSG} with games = {a:?}")
    }

#[test]
fn fixed_tests() {
        do_fixed_test(&["1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"], 30);
        do_fixed_test(&["1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"], 10);
        do_fixed_test(&["0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"], 0);
        do_fixed_test(&["1:0", "2:0", "3:0", "4:0", "2:1", "1:3", "1:4", "2:3", "2:4", "3:4"], 15);
        do_fixed_test(&["1:0", "2:0", "3:0", "4:4", "2:2", "3:3", "1:4", "2:3", "2:4", "3:4"], 12);
    }
}
