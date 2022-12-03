// https://www.codewars.com/kata/55a2d7ebe362935a210000b2/train/rust
pub fn find_smallest_int(arr: &[i32]) -> i32 {
    arr.iter().fold(arr[0], |acc, x| if x < &acc { *x } else { acc })
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test1() {
        assert_eq!(find_smallest_int(&[34, 15, 88, 2]), 2);
    }

    #[test]
    fn test2() {
        assert_eq!(find_smallest_int(&[34, -345, -1, 100]), -345);
    }
}
