#include <vector>
#include <string>

// https://www.codewars.com/kata/526571aae218b8ee490006f4/train/cpp
std::vector<std::string> towerBuilder(int nFloors)
{
    const char empty = ' ';
    std::vector<std::string> tree = std::vector<std::string>();
    for (int i = 0; i < nFloors; i++)
    {
        std::string level = std::string("");
        for (int j = 0; j < nFloors * 2 - 1; j++)
        {
            level.append("*");
        }
        level.insert(0, nFloors - 1 - i, empty);
        level.insert(nFloors + i, nFloors - 1 - i, empty);
        tree.push_back(level.substr(0, nFloors * 2 - 1));
    }
    return tree;
}
