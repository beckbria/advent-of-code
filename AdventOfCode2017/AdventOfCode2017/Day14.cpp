#include "Problems.h"
/*
--- Day 14: Disk Defragmentation ---

Suddenly, a scheduled job activates the system's disk defragmenter. Were the situation different, you might sit and watch it 
for a while, but today, you just don't have that kind of time. It's soaking up valuable system resources that are needed 
elsewhere, and so the only option is to help it finish its task as soon as possible.

The disk in question consists of a 128x128 grid; each square of the grid is either free or used. On this disk, the state of 
the grid is tracked by the bits in a sequence of knot hashes.

A total of 128 knot hashes are calculated, each corresponding to a single row in the grid; each hash contains 128 bits 
which correspond to individual grid squares. Each bit of a hash indicates whether that square is free (0) or used (1).

The hash inputs are a key string (your puzzle input), a dash, and a number from 0 to 127 corresponding to the row. 
For example, if your key string were flqrgnkx, then the first row would be given by the bits of the knot hash of 
flqrgnkx-0, the second row from the bits of the knot hash of flqrgnkx-1, and so on until the last row, flqrgnkx-127.

The output of a knot hash is traditionally represented by 32 hexadecimal digits; each of these digits correspond to 
4 bits, for a total of 4 * 32 = 128 bits. To convert to bits, turn each hexadecimal digit to its equivalent binary 
value, high-bit first: 0 becomes 0000, 1 becomes 0001, e becomes 1110, f becomes 1111, and so on; a hash that begins 
with a0c2017... in hexadecimal would begin with 10100000110000100000000101110000... in binary.

Continuing this process, the first 8 rows and columns for key flqrgnkx appear as follows, using # to denote used 
squares, and . to denote free ones:

##.#.#..-->
.#.#.#.#   
....#.#.   
#.#.##.#   
.##.#...   
##..#..#   
.#...#..   
##.#.##.-->
|      |   
V      V   

In this example, 8108 squares are used across the entire 128x128 grid.

Given your actual key string, how many squares are used?

*/
namespace Day14 {
std::vector<std::vector<int>> BuildGrid(const std::string& seed)
{
    std::vector<std::vector<int>> grid;
    for (int i = 0; i < 128; ++i) {
        std::stringstream rowSeed;
        rowSeed << seed << '-' << i;
        grid.emplace_back(std::move(Day10::KnotHash(rowSeed.str())));
    }
    return grid;
}

int UsedSquares(const std::string& seed)
{
    auto grid = BuildGrid(seed);
    int used = 0;
    for (const auto &row : grid) {
        for (const auto &cell : row) {
            used += Helpers::CountBits(cell);
        }
    }
    return used;
}
} // namespace Day14

void Day14Tests()
{
    const std::string input = "flqrgnkx";
    const auto used = Day14::UsedSquares(input);
    if (used != 8108) std::cerr << "Test 14A Error: Got " << used << ", Expected 8108";
}

void Day14Problems()
{
    std::cout << "Day 14:\n";
    Day14Tests();
    const std::string input = "uugsqrei";
    std::cout << Day14::UsedSquares(input) << std::endl;
}