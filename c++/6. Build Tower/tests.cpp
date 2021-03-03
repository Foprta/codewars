#include "main.cpp"
#include <gtest/gtest.h>

TEST(TowerBuilderTest, Tests)
{
    std::vector<std::string> expected = {"*"};
    ASSERT_EQ(towerBuilder(1), expected);
    expected = {" * ", "***"};
    ASSERT_EQ(towerBuilder(2), expected);
    expected = {"  *  ", " *** ", "*****"};
    ASSERT_EQ(towerBuilder(3), expected);
}

int main(int argc, char **argv)
{
    testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}