#include "Problems.h"
/*
A new system policy has been put in place that requires all accounts to use a passphrase instead of simply 
a password. A passphrase consists of a series of words (lowercase letters) separated by spaces.

To ensure security, a valid passphrase must contain no duplicate words.

For example:

aa bb cc dd ee is valid.
aa bb cc dd aa is not valid - the word aa appears more than once.
aa bb cc dd aaa is valid - aa and aaa count as different words.

The system's full passphrase list is available as your puzzle input. How many passphrases are valid?

--- Part Two ---

For added security, yet another system policy has been put in place. Now, a valid passphrase must contain no two words that are anagrams of each other - that is, a passphrase is invalid if any word's letters can be rearranged to form any other word in the passphrase.

For example:

abcde fghij is a valid passphrase.
abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
iiii oiii ooii oooi oooo is valid.
oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.

Under this new system policy, how many passphrases are valid?
*/
namespace Day4 {
bool AllWordsUnique(const std::vector<std::string>& words)
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
    auto words = Helpers::Tokenize(line);
    return AllWordsUnique(words);
}

bool IsValidAnagramPassphrase(const std::string& line)
{
    auto words = Helpers::Tokenize(line);
    // Sort the characters in each word to detect anagrams
    for (auto& word : words) {
        std::sort(word.begin(), word.end());
    }
    return AllWordsUnique(words);
}
} // namespace Day4

void Day4Tests()
{
    const struct test {
        std::string input;
        bool answer;
    } testCaseA[] = { {"aa bb cc dd ee", true}, { "aa bb cc dd aa", false }, { "aa bb cc dd aaa", true } };

    for (auto &t : testCaseA) {
        bool result = Day4::IsValidPassphrase(t.input);
        if (result != t.answer) {
            std::cerr << "Test 4A failed: " << t.input << " => " << result << " (expected " << t.answer << ")" << std::endl;
        }
    }
}

void Day4Problems()
{
    std::cout << "Day 4:\n";
    Day4Tests();

    const auto input = Helpers::ReadFileLines("input_day4.txt");
    int validPassphrases = 0;
    int anagramPassphrases = 0;
    for (auto &line : input) {
        if (Day4::IsValidPassphrase(line)) ++validPassphrases;
        if (Day4::IsValidAnagramPassphrase(line)) ++anagramPassphrases;
    }
    std::cout << validPassphrases << std::endl;
    std::cout << anagramPassphrases << std::endl << std::endl;
}