#pragma once
#include "Hash.h"

namespace Day10
{
    std::vector<unsigned int> InitializeHash(unsigned int listSize);
    void ReverseSegment(std::vector<unsigned int> &hash, int startPosition, int length);
    void ComputeHash(const std::vector<unsigned int> &lengths, std::vector<unsigned int> &hash, int &skip, int &currentPosition);
}