// https://www.codewars.com/kata/54d4c8b08776e4ad92000835/train/rust

use std::cmp::Ordering;

fn is_perfect_power(n: u64) -> Option<(u64, u32)> {
    let mut power = (n as f64).log2().floor() as u32;
    let mut base: u64 = 2;

    loop {
        if power.lt(&2) {
            return None;
        };

        match base.pow(power).cmp(&n) {
            Ordering::Greater => power -= 1,
            Ordering::Less => base += 1,
            Ordering::Equal => return Some((base, power)),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::is_perfect_power;

    #[test]
    fn basic_examples() {
        assert_eq!(is_perfect_power(4), Some((2, 2)), "4 = 2^2");
        assert_eq!(is_perfect_power(9), Some((3, 2)), "9 = 3^2");
        assert_eq!(is_perfect_power(2), None, "2 is not a perfect power");
        assert_eq!(is_perfect_power(5), None, "5 is not a perfect power");
    }

    #[test]
    fn first_perfect_powers() {
        let pp = &[
            4, 8, 9, 16, 25, 27, 32, 36, 49, 64, 81, 100, 121, 125, 128, 144, 169, 196, 216, 225,
            243, 256, 289, 324, 343, 361, 400, 441, 484,
        ];
        for p in pp {
            if is_perfect_power(*p) == None {
                assert!(false, "{} wasn't recognized as a perfect power", p)
            } else {
                assert!(true)
            }
        }
    }
}
