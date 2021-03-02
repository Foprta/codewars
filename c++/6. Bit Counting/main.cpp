// https://www.codewars.com/kata/526571aae218b8ee490006f4/train/cpp
unsigned int countBits(unsigned long long n)
{
    unsigned int sum = 0;
    for (unsigned long long i = n; i > 0; i >>= 1)
    {
        sum += i % 2;
    }
    return sum;
}

int main()
{
    countBits(4239259906917024334);
    return 0;
}