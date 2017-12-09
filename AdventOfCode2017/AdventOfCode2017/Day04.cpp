#include "Problems.h"

/*
A new system policy has been put in place that requires all accounts to use a passphrase instead of simply a password. A passphrase consists of a series of words (lowercase letters) separated by spaces.

To ensure security, a valid passphrase must contain no duplicate words.

For example:

aa bb cc dd ee is valid.
aa bb cc dd aa is not valid - the word aa appears more than once.
aa bb cc dd aaa is valid - aa and aaa count as different words.

The system's full passphrase list is available as your puzzle input. How many passphrases are valid?
*/

bool AllWordsUnique(std::vector<std::string>& words)
{
    std::set<std::string> seen;
    for (auto &word : words) {
        if (seen.count(word) == 1) return false;
        seen.emplace(word);
    }

    return true;
}

bool IsValidPassphrase(const std::string& line)
{
    // This is where C++ starts to fall apart.  Oh for a language with a standard string.split function.......
    auto words = Tokenize(line);
    return AllWordsUnique(words);
}

bool IsValidAnagramPassphrase(const std::string& line)
{
    auto words = Tokenize(line);
    // Sort the characters in each word to detect anagrams
    for (auto& word : words)
    {
        std::sort(word.begin(), word.end());
    }
    return AllWordsUnique(words);
}

void Day4Tests()
{
    struct test {
        std::string input;
        bool answer;
    } testCaseA[] = { {"aa bb cc dd ee", true}, { "aa bb cc dd aa", false }, { "aa bb cc dd aaa", true } };

    for (auto &t : testCaseA) {
        bool result = IsValidPassphrase(t.input);
        if (result != t.answer) {
            std::cerr << "Test 4A failed: " << t.input << " => " << result << " (expected " << t.answer << ")" << std::endl;
        }
    }
}

void Day4()
{
    Day4Tests();

    auto input = ReadFileLines("input_day4.txt");

    int validPassphrases = 0;
    int anagramPassphrases = 0;
    for (auto &line : input)
    {
        if (IsValidPassphrase(line)) ++validPassphrases;
        if (IsValidAnagramPassphrase(line)) ++anagramPassphrases;
    }
    std::cout << "Day 4:\n";
    std::cout << validPassphrases << std::endl;
    std::cout << anagramPassphrases << std::endl << std::endl;
}