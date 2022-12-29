// https://www.codewars.com/kata/530e15517bc88ac656000716/train/rust
fn rot13(message: &str) -> String {
    message
        .chars()
        .into_iter()
        .map(|ch| match ch {
            'a' => 'n',
            'b' => 'o',
            'c' => 'p',
            'd' => 'q',
            'e' => 'r',
            'f' => 's',
            'g' => 't',
            'h' => 'u',
            'i' => 'v',
            'j' => 'w',
            'k' => 'x',
            'l' => 'y',
            'm' => 'z',
            'n' => 'a',
            'o' => 'b',
            'p' => 'c',
            'q' => 'd',
            'r' => 'e',
            's' => 'f',
            't' => 'g',
            'u' => 'h',
            'v' => 'i',
            'w' => 'j',
            'x' => 'k',
            'y' => 'l',
            'z' => 'm',
            'A' => 'N',
            'B' => 'O',
            'C' => 'P',
            'D' => 'Q',
            'E' => 'R',
            'F' => 'S',
            'G' => 'T',
            'H' => 'U',
            'I' => 'V',
            'J' => 'W',
            'K' => 'X',
            'L' => 'Y',
            'M' => 'Z',
            'N' => 'A',
            'O' => 'B',
            'P' => 'C',
            'Q' => 'D',
            'R' => 'E',
            'S' => 'F',
            'T' => 'G',
            'U' => 'H',
            'V' => 'I',
            'W' => 'J',
            'X' => 'K',
            'Y' => 'L',
            'Z' => 'M',
            _ => ch,
        })
        .collect()
}

#[cfg(test)]
mod tests {
    use super::rot13;

    const ERR_MSG: &str = "\nYour result (left) did not match the expected output (right)";

    fn dotest(s: &str, expected: &str) {
        assert_eq!(rot13(s), expected, "{ERR_MSG} with message = \"{s}\"")
    }

    #[test]
    fn sample_tests() {
        dotest("§", "§");
        dotest("test", "grfg");
        dotest("Test", "Grfg");
    }
}
