#pragma once
#include <algorithm>
#include <array>
#include <chrono>
#include <climits>
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
#include <unordered_set>
#include <utility>
#include <vector>

#define ARRAYSIZE(a) (sizeof(a) / sizeof(a[0]))

// Helper Functions

namespace Helpers
{
    void HelpersTests();
    std::vector<std::string> ReadFileLines(const std::string &fileName);
    std::vector<std::string> Tokenize(const std::string &line, char delimiter = ' ', bool splitWhitespace = true);
    // Removes all instances of the specified character at the end of a string.  Used for removing commas, etc.
    void RemoveTrailingCharacter(std::string &toBeModified, char toBeRemoved);
    std::string ByteArrayToHex(const std::vector<unsigned int> &bytes);

    template <typename T>
    constexpr bool IsSingleBitSet(T value) { return (value > 0) && ((value & (value - 1)) == 0); }

    inline int CountBits(int i)
    {
        int bits = 0;
        while (i != 0)
        {
            i &= (i - 1);
            ++bits;
        }
        return bits;
    }

    template <class T>
    std::vector<T> ReadFile(const std::string &fileName)
    {
        // Read the Input
        std::ifstream inputFile(fileName);
        std::vector<T> input;
        for (T item; inputFile >> item;)
        {
            input.emplace_back(std::move(item));
        }
        inputFile.close();
        return input;
    }

    template <class T>
    void RemoveIndexes(std::vector<T> &content, std::vector<unsigned int> indexesToErase)
    {
        std::sort(indexesToErase.begin(), indexesToErase.end());
        // Copy the elements in a single pass to avoid repeated copies as we erase from the array
        unsigned int writePosition = 0;
        unsigned int readPosition = 0;
        auto eraseIndex = indexesToErase.begin();
        while (readPosition < content.size())
        {
            if ((eraseIndex != indexesToErase.end()) && (readPosition == *eraseIndex))
            {
                // We should skip copying this element because it's scheduled for erasing
                ++readPosition;
                const auto erasing = *eraseIndex;
                // Check for any duplicate values that were passed to us
                while ((eraseIndex != indexesToErase.end()) && (*eraseIndex == erasing))
                    ++eraseIndex;
            }
            else
            {
                content[writePosition++] = content[readPosition++];
            }
        }
        // Discard the unused elements
        content.resize(writePosition);
    }

} // namespace Helpers

struct IntDefaultToZero
{
    int val = 0;
};
