// https://www.codewars.com/kata/634d0723075de3f97a9eb604/train/rust
pub fn encode(s: String) -> String {
    let mut result = String::new();

    let mut shift = 0;

    if s.len() == 0 {
        return s;
    };

    loop {
        result.push(s.chars().nth(shift).unwrap());

        if shift >= s.chars().count() - shift - 1 {
            break;
        }

        result.push(s.chars().nth(s.chars().count() - shift - 1).unwrap());

        if shift + 1 >= s.chars().count() - shift - 1 {
            break;
        }

        shift = shift + 1;
    }

    result
}

pub fn decode(s: String) -> String {
    let mut left = String::new();
    let mut right = String::new();

    s.chars().enumerate().for_each(|(idx, ch)| {
        if idx % 2 == 0 {
            left.push(ch);
        } else {
            right.insert(0, ch)
        }
    });

    left + &right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn encode_basic_tests() {
        assert_eq!("ceod".to_string(), encode("code".to_string()));
        assert_eq!("csordaew".to_string(), encode("codewars".to_string()));
        assert_eq!("wehti".to_string(), encode("white".to_string()));
        assert_eq!("Atsrse".to_string(), encode("Assert".to_string()));
        assert_eq!(
            "H!edlllroo w".to_string(),
            encode("Hello world!".to_string())
        );
        assert_eq!(
            "Y.oaut ahka vsei hcth oesteanl stnoa rt".to_string(),
            encode("You have chosen to translate this kata.".to_string())
        );
        assert_eq!(" ".to_string(), encode(" ".to_string()));
        assert_eq!("".to_string(), encode("".to_string()));
    }

    #[test]
    fn decode_basic_tests() {
        assert_eq!("codewars".to_string(), decode("csordaew".to_string()));
        assert_eq!("white".to_string(), decode("wehti".to_string()));
        assert_eq!("Assert".to_string(), decode("Atsrse".to_string()));
        assert_eq!(
            "Hello world!".to_string(),
            decode("H!edlllroo w".to_string())
        );
        assert_eq!(
            "You have chosen to translate this kata.".to_string(),
            decode("Y.oaut ahka vsei hcth oesteanl stnoa rt".to_string())
        );
        assert_eq!(" ".to_string(), decode(" ".to_string()));
        assert_eq!("".to_string(), decode("".to_string()));
    }
}
