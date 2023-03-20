// https://www.codewars.com/kata/526989a41034285187000de4/train/rust
const MULTIPLICATORS: [u32; 4] = [256 * 256 * 256, 256 * 256, 256, 1];

fn ips_between(start: &str, end: &str) -> u32 {
    let start: Vec<u8> = start.split(".").map(|x| x.parse::<u8>().unwrap()).collect();
    let end: Vec<u8> = end.split(".").map(|x| x.parse::<u8>().unwrap()).collect();

    start
        .iter()
        .zip(end)
        .enumerate()
        .fold(0, |acc, (idx, (s, e))| {
            if s > &e {
                acc - ((s - e) as u32 * MULTIPLICATORS[idx])
            } else if s < &e {
                acc + ((e - s) as u32 * MULTIPLICATORS[idx])
            } else {
                acc
            }
        })
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn basic() {
        assert_eq!(ips_between("10.0.0.0", "10.0.0.50"), 50);
        assert_eq!(ips_between("20.0.0.10", "20.0.1.0"), 246);
        assert_eq!(ips_between("20.0.0.0", "20.0.1.0"), 256);
        assert_eq!(ips_between("20.0.0.20", "20.0.2.20"), 512);
        assert_eq!(ips_between("20.0.0.0", "20.1.0.0"), 65536);
    }
}
