#pragma once
#include <algorithm>
#include <array>
#include <chrono>
#include <cmath>
#include <cstdlib>
#include <fstream>
#include <functional>
#include <iostream>
#include <iterator>
#include <list>
#include <map>
#include <memory>
#include <mutex>
#include <queue>
#include <set>
#include <sstream>
#include <stack>
#include <string>
#include <unordered_map>
#include <utility>
#include <vector>

#include "Day10.h"

#define ARRAYSIZE(a) (sizeof(a) / sizeof(a[0]))

// Helper Functions

namespace Helpers
{
std::vector<std::string> ReadFileLines(const std::string& fileName);
std::vector<std::string> Tokenize(const std::string& line, char delimiter = ' ', bool splitWhitespace = true);
// Removes all instances of the specified character at the end of a string.  Used for removing commas, etc.
void RemoveTrailingCharacter(std::string& toBeModified, char toBeRemoved);
std::string ByteArrayToHex(const std::vector<unsigned int>& bytes);

inline int CountBits(int i)
{
    int bits = 0;
    while (i != 0) {
        i &= (i - 1);
        ++bits;
    }
    return bits;
}

template<class T>
std::vector<T> ReadFile(const std::string& fileName)
{
    // Read the Input
    std::ifstream inputFile(fileName);
    std::vector<T> input;
    for (T item; inputFile >> item; )
    {
        input.emplace_back(std::move(item));
    }
    inputFile.close();
    return input;
}
} // namespace Helpers

struct IntDefaultToZero {
    int val = 0;
};

// Daily Functions
void Day1Problems();
void Day2Problems();
void Day3Problems();
void Day4Problems();
void Day5Problems();
void Day6Problems();
void Day7Problems();
void Day8Problems();
void Day9Problems();
void Day10Problems();
void Day11Problems();
void Day12Problems();
void Day13Problems();
void Day14Problems();
void Day15Problems();
void Day16Problems();
void Day17Problems();
void Day18Problems();
void Day19Problems();
void Day20Problems();