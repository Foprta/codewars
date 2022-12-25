// https://www.codewars.com/kata/5264d2b162488dc400000001/train/rust
fn spin_words(words: &str) -> String {
    let words = words.split_whitespace();

    let result = words.fold(Vec::new(), |mut acc, str| {
        let count = str.chars().count();

        let str = match count >= 5 {
            true => str.chars().rev().collect::<String>(),
            false => str.to_owned(),
        };

        acc.push(str);

        acc
    });

    result.join(" ")
}

#[cfg(test)]
mod tests {
    use super::spin_words;

    #[test]
    fn examples() {
        assert_eq!(spin_words("Welcome"), "emocleW");
        assert_eq!(spin_words("Hey fellow warriors"), "Hey wollef sroirraw");
        assert_eq!(spin_words("This is a test"), "This is a test");
        assert_eq!(spin_words("This is another test"), "This is rehtona test");
        assert_eq!(spin_words("You are almost to the last test"), "You are tsomla to the last test");
        assert_eq!(spin_words("Just kidding there is still one more"), "Just gniddik ereht is llits one more");
        assert_eq!(spin_words("Seriously this is the last one"), "ylsuoireS this is the last one");
    }
}