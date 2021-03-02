package kyu7.HighestAndLowest;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class KataTest {
    @Test
    public void test1() {
        assertEquals("42 -9", Kata.highAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"));
        assertEquals("9 -5", Kata.highAndLow("1 9 3 4 -5"));
    }
}