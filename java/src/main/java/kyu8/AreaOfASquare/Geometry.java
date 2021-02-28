package kyu8.AreaOfASquare;

import java.text.DecimalFormat;

public class Geometry {
    public static double squareArea(double A) {
        double radius = A / (0.5 * Math.PI);
        DecimalFormat df = new DecimalFormat("#.00");
        return Double.parseDouble(df.format(radius * radius));
    }
}