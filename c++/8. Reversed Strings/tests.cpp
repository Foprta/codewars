#include "main.cpp"
#include <gtest/gtest.h>

TEST(ReverseStringTests, Tests)
{
    ASSERT_EQ("apoj", reverseString("jopa"));
    ASSERT_EQ("world", reverseString("dlrow"));
    ASSERT_EQ("", reverseString(""));
}

int main(int argc, char **argv)
{
    testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}