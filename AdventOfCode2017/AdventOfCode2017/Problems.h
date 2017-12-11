#pragma once
#include <algorithm>
#include <array>
#include <cmath>
#include <cstdlib>
#include <fstream>
#include <iostream>
#include <iterator>
#include <map>
#include <memory>
#include <queue>
#include <set>
#include <sstream>
#include <stack>
#include <string>
#include <utility>
#include <vector>

#define ARRAYSIZE(a) (sizeof(a) / sizeof(a[0]))

// Helper Functions
std::vector<std::string> ReadFileLines(const std::string& fileName);
std::vector<std::string> Tokenize(const std::string& line, char delimiter = ' ', bool splitWhitespace = true);

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

struct IntDefaultToZero {
    int val = 0;
};

// Daily Functions
void Day1();
void Day2();
void Day3();
void Day4();
void Day5();
void Day6();
void Day7();
void Day8();
void Day9();
void Day10();
void Day11();