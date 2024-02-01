use std::collections::HashMap;

// https://www.codewars.com/kata/550498447451fbbd7600041c/train/rust
fn comp(a: Vec<i64>, b: Vec<i64>) -> bool {
    let mut map = HashMap::<i64, i8>::new();
    for ele in b.iter() {
        match map.get(ele) {
            Some(count) => {
                map.insert(*ele, count + 1);
            }
            None => {
                map.insert(*ele, 1);
            }
        }
    }
    
    for ele in a.iter() {
        let square = ele * ele;
        match map.get(&square) {
            Some(count) => {
                if count.le(&0) {
                    return false;
                } else {
                    map.insert(square, count - 1);
                }
            }
            None => {
                return false;
            }
        }
    }

    return map.values().all(|&v| v == 0)
}

#[test]
fn tests_comp() {
    let a1 = vec![121, 144, 19, 161, 19, 144, 19, 11];
    let a2 = vec![
        11 * 11,
        121 * 121,
        144 * 144,
        19 * 19,
        161 * 161,
        19 * 19,
        144 * 144,
        19 * 19,
    ];
    assert_eq!(comp(a1, a2), true);
    let a1 = vec![121, 144, 19, 161, 19, 144, 19, 11];
    let a2 = vec![
        11 * 21,
        121 * 121,
        144 * 144,
        19 * 19,
        161 * 161,
        19 * 19,
        144 * 144,
        19 * 19,
    ];
    assert_eq!(comp(a1, a2), false);
}
