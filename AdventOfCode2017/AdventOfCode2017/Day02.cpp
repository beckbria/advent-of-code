#include "Problems.h"

int checksum(std::vector<std::vector<int>>& spreadsheet)
{
    int checksum = 0;

    for (auto &v : spreadsheet)
    {
        int smallest = INT_MAX;
        int largest = INT_MIN;

        for (auto &i : v)
        {
            smallest = std::min(smallest, i);
            largest = std::max(largest, i);
        }

        checksum += (largest - smallest);
    }

    return checksum;
}

int divisorChecksum(std::vector<std::vector<int>>& spreadsheet)
{
    int checksum = 0;

    for (auto &v : spreadsheet)
    {
        bool found = false;
        for (int i = 1; (i < v.size()) && !found; ++i)
        {
            for (int j = 0; (j < i) && !found; ++j)
            {
                bool const iLarger = (v[i] > v[j]);
                int larger = iLarger ? v[i] : v[j];
                int smaller = iLarger ? v[j] : v[i];

                if (larger % smaller == 0)
                {
                    checksum += larger / smaller;
                    found = true;
                }
            }
        }
    }

    return checksum;
}

void Day2Tests()
{
    std::vector<std::vector<int>> testA = { { 5,1,9,5 },{ 7,5,3 },{ 2,4,6,8 } };
    int A = checksum(testA);
    if (A != 18) std::cerr << "TestA failed: " << A << " (expected 18)" << std::endl;

    std::vector<std::vector<int>> testB = { { 5,9,2,8 }, { 9,4,7,3 }, { 3,8,6,5 } };
    int B = divisorChecksum(testB);
    if (B != 9) std::cerr << "TestA failed: " << B << " (expected 9)" << std::endl;
}

void Day2()
{
    Day2Tests();

    std::vector<std::vector<int>> spreadsheet = { {116, 1470, 2610, 179, 2161, 2690, 831, 1824, 2361, 1050, 2201, 118, 145, 2275, 2625, 2333},
    { 976,220,1129,553,422,950,332,204,1247,1092,1091,159,174,182,984,713 },
    { 84,78,773,62,808,83,1125,1110,1184,145,1277,982,338,1182,75,679 },
    { 3413,3809,3525,2176,141,1045,2342,2183,157,3960,3084,2643,119,108,3366,2131 },
    { 1312,205,343,616,300,1098,870,1008,1140,1178,90,146,980,202,190,774 },
    { 4368,3905,3175,4532,3806,1579,4080,259,2542,221,4395,4464,208,3734,234,4225 },
    { 741,993,1184,285,1062,372,111,118,63,843,325,132,854,105,956,961 },
    { 85,79,84,2483,858,2209,2268,90,2233,1230,2533,322,338,68,2085,1267 },
    { 2688,2022,112,130,1185,103,1847,3059,911,107,2066,1788,2687,2633,415,1353 },
    { 76,169,141,58,161,66,65,225,60,152,62,64,156,199,80,56 },
    { 220,884,1890,597,3312,593,4259,222,113,2244,3798,4757,216,1127,4400,178 },
    { 653,369,216,132,276,102,265,889,987,236,239,807,1076,932,84,864 },
    { 799,739,75,1537,82,228,69,1397,1396,1203,1587,63,313,1718,1375,469 },
    { 1176,112,1407,136,1482,1534,1384,1202,604,851,190,284,1226,113,114,687 },
    { 73,1620,81,1137,812,75,1326,1355,1545,1666,1356,1681,1732,85,128,902},
    { 571,547,160,237,256,30,496,592,385,576,183,692,192,387,647,233 } };

    std::cout << "Day 2:\n" << checksum(spreadsheet) << std::endl << divisorChecksum(spreadsheet) << std::endl << std::endl;
}