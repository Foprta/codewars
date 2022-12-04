// https://www.codewars.com/kata/544675c6f971f7399a000e79/rust
pub fn string_to_number(s: &str) -> i32 {
    let mut result: i32 = 0;
    let mut modifier: i32 = 1;

    for symbol in s.chars() {
        if symbol == '-' {
            modifier = -1;
            continue;
        }

        result = result * 10 + match symbol {
            '0'..='9' => (symbol as i32) - 48,
            _ => 0
        }
    }

    result * modifier
}

#[cfg(test)]
mod tests {
    use super::string_to_number;

    #[test]
    fn returns_expected() {
        assert_eq!(string_to_number("1234"), 1234);
        assert_eq!(string_to_number("605"), 605);
        assert_eq!(string_to_number("1405"), 1405);
        assert_eq!(string_to_number("-7"), -7);
    }

    #[test]
    fn tessd() {
        print!("{}", '0' as u32);
     }
}
