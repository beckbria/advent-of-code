#include "Helpers.h"

namespace Helpers
{
    inline bool IsWhitespace(char c)
    {
        switch (c)
        {
        case ' ':
        case '\t':
        case '\r':
        case '\n':
            return true;
        }
        return false;
    }

    std::vector<std::string> Tokenize(const std::string &line, char delimiter /* = ' '*/, bool splitWhitespace /* = true*/)
    {
        std::vector<std::string> tokens;

        bool parsingToken = false;
        int start = 0;
        for (size_t pos = 0; pos < line.size(); ++pos)
        {
            const bool partOfToken = (line[pos] != delimiter) && !(splitWhitespace && IsWhitespace(line[pos]));
            if (partOfToken && !parsingToken)
            {
                // This is the beginning of a word
                start = pos;
                parsingToken = true;
            }
            else if (!partOfToken && parsingToken)
            {
                // The end of the word was the previous character
                tokens.emplace_back(std::move(line.substr(start, pos - start)));
                parsingToken = false;
            }
        }

        // If we hit the end of the string while parsing a token, we should include it
        if (parsingToken)
        {
            tokens.emplace_back(std::move(line.substr(start)));
        }

        return tokens;
    }

    std::vector<std::string> ReadFileLines(const std::string &fileName)
    {
        // Read the Input
        std::ifstream inputFile(fileName);
        std::vector<std::string> input;
        for (std::string line; std::getline(inputFile, line);)
        {
            input.emplace_back(std::move(line));
        }
        inputFile.close();
        return input;
    }

    void RemoveTrailingCharacter(std::string &toBeModified, char toBeRemoved)
    {
        int removalCount = 0;
        for (int pos = toBeModified.size() - 1; (pos >= 0) && (toBeModified[pos] == toBeRemoved); --pos)
        {
            ++removalCount;
        }
        toBeModified.resize(toBeModified.size() - removalCount);
    }

    // Assumes well-formed input
    inline char HexDigit(int i)
    {
        if (i < 10)
        {
            return ('0' + i);
        }
        else
        {
            return 'a' + (i - 10);
        }
    }

    std::string ByteArrayToHex(const std::vector<unsigned int> &bytes)
    {
        std::string output(bytes.size() * 2, ' ');
        for (size_t i = 0; i < bytes.size(); ++i)
        {
            auto current = bytes[i];
            output[2 * i] = HexDigit(current / 16);
            output[(2 * i) + 1] = HexDigit(current % 16);
        }
        return output;
    }

    void HelpersTests()
    {
        // Next available Helper Test #: 2

        std::vector<char> removeIndexInput = {'a', 'b', 'c', 'd', 'e', 'f', 'g'};
        std::vector<unsigned int> toRemove = {3, 0, 0, 2, 3};
        const std::vector<char> expectedRemoveIndexOutput = {'b', 'e', 'f', 'g'};
        RemoveIndexes(removeIndexInput, toRemove);
        for (unsigned int i = 0; i < removeIndexInput.size(); ++i)
        {
            if (removeIndexInput[i] != expectedRemoveIndexOutput[i])
            {
                std::cerr << "Helpers Test 1 Error: Got ";
                for (auto c : removeIndexInput)
                    std::cerr << c;
                std::cerr << ", Expected ";
                for (auto c : expectedRemoveIndexOutput)
                    std::cerr << c;
                std::cerr << std::endl;
                break;
            }
        }
    }

} // namespace Helpers