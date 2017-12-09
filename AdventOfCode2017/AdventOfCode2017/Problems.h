#pragma once
#include <iostream>
#include <fstream>
#include <sstream>
#include <iterator>
#include <string>
#include <vector>
#include <cmath>
#include <set>
#include <algorithm>
#include <sstream>
#include <utility>
#include <map>
#include <memory>
#include <cstdlib>
#include <queue>
#include <stack>

#define ARRAYSIZE(a) (sizeof(a) / sizeof(a[0]))

// Helper Functions
std::vector<std::string> ReadFileLines(const std::string& fileName);
std::vector<std::string> Tokenize(const std::string& line);

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