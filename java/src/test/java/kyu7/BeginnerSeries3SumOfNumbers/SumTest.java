package kyu7.BeginnerSeries3SumOfNumbers;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

// https://www.codewars.com/kata/55f2b110f61eb01779000053/train/java
public class SumTest {
    Sum s = new Sum();

    @Test
    public void Test1() {
        assertEquals(-1, s.GetSum(0, -1));
    }

    @Test
    public void Test2() {
        assertEquals(1, s.GetSum(0, 1));
    }

    @Test
    public void Test3() {
        assertEquals(1, s.GetSum(1, 1));
    }

    @Test
    public void Test4() {
        assertEquals(-1, s.GetSum(-1, -1));
    }

    @Test
    public void Test5() {
        assertEquals(10, s.GetSum(1, 4));
    }
}
