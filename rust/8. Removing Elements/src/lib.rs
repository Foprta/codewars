// https://www.codewars.com/kata/5769b3802ae6f8e4890009d2/train/rust
pub fn remove_every_other(arr: &[u8]) -> Vec<u8> {
    arr.into_iter().cloned().step_by(2).collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sample_test() {
        assert_eq!(remove_every_other(&[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]), &[1, 3, 5, 7, 9]);
    }
}
