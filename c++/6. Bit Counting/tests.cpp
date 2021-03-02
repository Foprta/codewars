#include "main.cpp"
#include <gtest/gtest.h>

TEST(CountBitsTests, Tests)
{
    ASSERT_EQ(countBits(0), 0);
    ASSERT_EQ(countBits(4), 1);
    ASSERT_EQ(countBits(7), 3);
    ASSERT_EQ(countBits(9), 2);
    ASSERT_EQ(countBits(10), 2);
    ASSERT_EQ(countBits(26), 3);
    ASSERT_EQ(countBits(77231418), 14);
    ASSERT_EQ(countBits(12525589), 11);
    ASSERT_EQ(countBits(3811), 8);
    ASSERT_EQ(countBits(392902058), 17);
    ASSERT_EQ(countBits(4239259906917024334), 30);
}

int main(int argc, char **argv)
{
    testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}