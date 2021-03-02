package kyu7.BeginnerSeries3SumOfNumbers;

// https://www.codewars.com/kata/55f2b110f61eb01779000053/train/java
public class Sum {
    public int GetSum(int a, int b) {
        Integer sum = 0;
        if (b < a) {
            Integer buffer = a;
            a = b;
            b = buffer;
        }
        for (Integer i = a; i <= b; i++) {
            sum += i;
        }
        return sum;
    }
}
