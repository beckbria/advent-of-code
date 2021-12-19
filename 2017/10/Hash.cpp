#include "Hash.h"

namespace Day10
{
    std::vector<unsigned int> InitializeHash(unsigned int listSize)
    {
        std::vector<unsigned int> hash;
        hash.resize(listSize);
        // Initialize the list to 0....listSize-1
        for (unsigned int i = 0; i < listSize; ++i)
        {
            hash[i] = i;
        }
        return hash;
    }

    void ReverseSegment(std::vector<unsigned int> &hash, int startPosition, int length)
    {
        int start = startPosition % hash.size();
        int end = startPosition + (length - 1);
        while (end > start)
        {
            std::swap(hash[start % hash.size()], hash[end % hash.size()]);
            ++start;
            --end;
        }
    }

    void ComputeHash(const std::vector<unsigned int> &lengths, std::vector<unsigned int> &hash, int &skip, int &currentPosition)
    {
        for (auto len : lengths)
        {
            ReverseSegment(hash, currentPosition, len);
            currentPosition = (currentPosition + len + skip) % hash.size();
            ++skip;
        }
    }

    std::vector<unsigned int> KnotHash(std::string plainText)
    {
        // For readability, we're going to explicitly convert the input string to a vector of integers.  Yes, it's true
        // that a std::string is array-like behind the scenes, and if we were dealing with megabyte-scale input avoiding
        // the copy would be nice, but here's it's a trivial amount of work/memory, so just copy rather than genericizing
        // the function further
        std::vector<unsigned int> lengths(plainText.begin(), plainText.end());

        // Per problem instructions, append a few values to the end
        for (auto i : {17, 31, 73, 47, 23})
            lengths.push_back(i);
        auto sparseHash = InitializeHash(256);

        // Do 64 rounds of the hash function to obtain a permutation of the initial list
        int skip = 0, currentPosition = 0;
        for (unsigned int i = 0; i < 64; ++i)
        {
            ComputeHash(lengths, sparseHash, skip, currentPosition);
        }

        // Compute the Dense Hash by XORing each 16 characters together
        std::vector<unsigned int> denseHash;
        unsigned int current = sparseHash[0];
        for (size_t i = 1; i < sparseHash.size(); ++i)
        {
            if (i % 16 == 0)
            {
                // We've filled in a block
                denseHash.push_back(current);
                current = sparseHash[i];
            }
            else
            {
                current ^= sparseHash[i];
            }
        }
        // Add the last element
        denseHash.push_back(current);
        return denseHash;
    }
}