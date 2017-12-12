#include "Problems.h"
/*
--- Day 9: Stream Processing ---

A large stream blocks your path. According to the locals, it's not safe to cross the stream at the moment because it's full of garbage. You look down at the stream; rather than water, you discover that it's a stream of characters.

You sit for a while and record part of the stream (your puzzle input). The characters represent groups - sequences that begin with { and end with }. Within a group, there are zero or more other things, separated by commas: either another group or garbage. Since groups can contain other groups, a } only closes the most-recently-opened unclosed group - that is, they are nestable. Your puzzle input represents a single, large group which itself contains many smaller ones.

Sometimes, instead of a group, you will find garbage. Garbage begins with < and ends with >. Between those angle brackets, almost any character can appear, including { and }. Within garbage, < has no special meaning.

In a futile attempt to clean up the garbage, some program has canceled some of the characters within it using !: inside garbage, any character that comes after ! should be ignored, including <, >, and even another !.

You don't see any characters that deviate from these rules. Outside garbage, you only find well-formed groups, and garbage always terminates according to the rules above.

Here are some self-contained pieces of garbage:

<>, empty garbage.
<random characters>, garbage containing random characters.
<<<<>, because the extra < are ignored.
<{!>}>, because the first > is canceled.
<!!>, because the second ! is canceled, allowing the > to terminate the garbage.
<!!!>>, because the second ! and the first > are canceled.
<{o"i!a,<{i<a>, which ends at the first >.

Here are some examples of whole streams and the number of groups they contain:

{}, 1 group.
{{{}}}, 3 groups.
{{},{}}, also 3 groups.
{{{},{},{{}}}}, 6 groups.
{<{},{},{{}}>}, 1 group (which itself contains garbage).
{<a>,<a>,<a>,<a>}, 1 group.
{{<a>},{<a>},{<a>},{<a>}}, 5 groups.
{{<!>},{<!>},{<!>},{<a>}}, 2 groups (since all but the last > are canceled).

Your goal is to find the total score for all groups in your input. Each group is assigned a score which is one more than the score of the group that immediately contains it. (The outermost group gets a score of 1.)

{}, score of 1.
{{{}}}, score of 1 + 2 + 3 = 6.
{{},{}}, score of 1 + 2 + 2 = 5.
{{{},{},{{}}}}, score of 1 + 2 + 3 + 3 + 3 + 4 = 16.
{<a>,<a>,<a>,<a>}, score of 1.
{{<ab>},{<ab>},{<ab>},{<ab>}}, score of 1 + 2 + 2 + 2 + 2 = 9.
{{<!!>},{<!!>},{<!!>},{<!!>}}, score of 1 + 2 + 2 + 2 + 2 = 9.
{{<a!>},{<a!>},{<a!>},{<ab>}}, score of 1 + 2 = 3.

What is the total score for all groups in your input?

--- Part Two ---

Now, you're ready to remove the garbage.

To prove you've removed it, you need to count all of the characters within the garbage. The leading and trailing < and > don't count, nor do any canceled characters or the ! doing the canceling.

<>, 0 characters.
<random characters>, 17 characters.
<<<<>, 3 characters.
<{!>}>, 2 characters.
<!!>, 0 characters.
<!!!>>, 0 characters.
<{o"i!a,<{i<a>, 10 characters.

How many non-canceled characters are within the garbage in your puzzle input?
*/
namespace Day9 {
struct Group
{
    Group(int start) : beginPosition(start) {}
    int beginPosition;
    int endPosition = -1;
    int garbageCharacters = 0;

    int Score(int parentScore = 0) const {
        int currentScore = parentScore + 1;
        int totalScore = currentScore;
        for (auto &c : children) totalScore += c->Score(currentScore);
        return totalScore;
    }

    int TotalGarbageCharacters() const {
        int totalGarbage = garbageCharacters;
        for (auto &c : children) totalGarbage += c->TotalGarbageCharacters();
        return totalGarbage;
    }

    std::vector<std::shared_ptr<Group>> children;
};

// Understand the groups in a source text.  This assumes a well-formatted input since that's guaranteed in the problem description
std::shared_ptr<Group> Parse(const std::string& input)
{
    static constexpr char GroupBegin = '{';
    static constexpr char GroupEnd = '}';
    static constexpr char GarbageBegin = '<';
    static constexpr char GarbageEnd = '>';
    static constexpr char NegateNext = '!';

    std::stack<std::shared_ptr<Group>> groups;
    std::shared_ptr<Group> firstGroup;
    
    bool inGarbage = false;
    bool negateNextInput = false;
    for (size_t pos = 0; pos < input.size(); ++pos) {
        // If the previous character was ! we should ignore the next one regardless of what it is
        if (negateNextInput) {
            negateNextInput = false;
            continue;
        }

        const char current = input[pos];

        if (current == NegateNext) {
            negateNextInput = true;
            continue;
        }

        // If we're processing garbage, then we ignore any other characters except the end of garbage
        if (inGarbage) {
            if (current == GarbageEnd) {
                inGarbage = false;
            }
            else {
                groups.top()->garbageCharacters++;
            }
            continue;
        }

        // Otherwise, we're processing groups normally
        switch (current) {
        case GarbageBegin:
            inGarbage = true;
            break;

        case GroupEnd:
            groups.top()->endPosition = pos;
            groups.pop();
            break;

        case GroupBegin:
            auto newGroup = std::make_shared<Group>(pos);
            if (!groups.empty()) {
                // This group is a child of the group preceding it on the stack, unless it's the very first one
                groups.top()->children.push_back(newGroup);
            }
            else {
                // We should keep a reference to the very first group around, or else we risk losing it when we parse
                // its GroupEnd token and pop it off the stack
                firstGroup = newGroup;
            }
            groups.push(newGroup);
            break;
        }
    }

    return firstGroup;
}

std::pair<int, int> Score(const std::string& input) 
{
    auto group = Parse(input);
    return std::make_pair(group->Score(), group->TotalGarbageCharacters());
}
} // namespace Day9

void Day9Tests() 
{
    const struct {
        std::string input;
        int score;
    } testCases[] = {
        { "{}", 1 },
        { "{{{}}}", 6 },
        { "{{},{}}", 5 },
        { "{{{},{},{{}}}}", 16 },
        { "{<a>,<a>,<a>,<a>}", 1 },
        { "{{<ab>},{<ab>},{<ab>},{<ab>}}", 9 },
        { "{{<!!>},{<!!>},{<!!>},{<!!>}}", 9 },
        { "{{<a!>},{<a!>},{<a!>},{<ab>}}", 3 }
    };

    for (auto& test : testCases) {
        auto result = Day9::Score(test.input);
        if (result.first != test.score) {
            std::cerr << "Test 9A Failed: Got " << result.first << ", expected " << test.score << std::endl;
        }
    }

    const struct {
        std::string input;
        int score;
    } testCasesB[] = {
        { "{<>}", 0 },
        { "{<random characters>}", 17 },
        { "{<<<<>}", 3 },
        { "{<{!>}>}", 2 },
        { "{<!!>}", 0 },
        { "{<!!!>>}", 0 },
        { "{<{o\"i!a,<{i<a>}", 10},
    };

    for (auto& test : testCasesB) {
        auto result = Day9::Score(test.input);
        if (result.second != test.score) {
            std::cerr << "Test 9B Failed: Got " << result.second << ", expected " << test.score << std::endl;
        }
    }
}

void Day9Problems()
{
    Day9Tests();

    auto input = ReadFileLines("input_day9.txt");
    std::cout << "Day 9:\n";
    if (input.size() != 1) std::cerr << "Day 9: Malformed input" << std::endl;
    auto score = Day9::Score(input[0]);
    std::cout << score.first << std::endl << score.second << std::endl << std::endl;
}