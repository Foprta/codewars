package kyu7.HighestAndLowest;

// https://www.codewars.com/kata/55f2b110f61eb01779000053/java
public class Kata {
    public static String highAndLow(String numbers) {
        String[] nums = numbers.split(" ");
        Integer lowest = Integer.MAX_VALUE;
        Integer highest = Integer.MIN_VALUE;
        for (String num : nums) {
            Integer buff = Integer.parseInt(num);
            if (lowest > buff) {
                lowest = buff;
            }
            if (highest < buff) {
                highest = buff;
            }
        }
        return Integer.toString(highest) + " " + Integer.toString(lowest);
    }
}
