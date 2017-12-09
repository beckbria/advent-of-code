#include "Problems.h"

std::vector<std::string> Tokenize(const std::string& line)
{
    std::istringstream ss(line);
    std::istream_iterator<std::string> begin(ss), end;

    //putting all the tokens in the vector
    std::vector<std::string> tokens(begin, end);
    return tokens;
}

std::vector<std::string> ReadFileLines(const std::string& fileName)
{
    // Read the Input
    std::ifstream inputFile(fileName);
    std::vector<std::string> input;
    for (std::string line; std::getline(inputFile, line); )
    {
        input.emplace_back(std::move(line));
    }
    inputFile.close();
    return input;
}