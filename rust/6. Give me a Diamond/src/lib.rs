// https://www.codewars.com/kata/5503013e34137eeeaa001648/train/rust
fn print(n: i32) -> Option<String> {
    if n <= 0 || n % 2 == 0 {
        return None;
    }

    let mut result = String::new();

    let mut stars: i32 = 1;
    let mut spaces: i32 = n / 2;

    for _ in 0..n {
        let space = " ".repeat(spaces.abs() as usize);
        let star = "*".repeat(stars as usize);

        result.push_str(&space);
        result.push_str(&star);
        result.push('\n');

        if spaces <= 0 {
            stars -= 2;
        } else {
            stars += 2;
        }

        spaces -= 1;
    }

    Some(result)
}

#[test]
fn basic_test() {
    assert_eq!(print(3), Some(" *\n***\n *\n".to_string()));
    assert_eq!(print(5), Some("  *\n ***\n*****\n ***\n  *\n".to_string()));
    assert_eq!(print(-3), None);
    assert_eq!(print(2), None);
    assert_eq!(print(0), None);
    assert_eq!(print(1), Some("*\n".to_string()));
}
