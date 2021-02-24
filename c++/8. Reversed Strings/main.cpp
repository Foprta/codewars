#include <string>
using namespace std;

// https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/cpp
string reverseString(string str)
{
    string result = "";

    for (int i = 0; i < str.length(); i++)
    {
        result = str[i] + result;
    }

    return result;
}