#include "Problems.h"

/*
You receive a signal directly from the CPU.Because of your recent assistance with jump instructions, it would like you to compute the result of a series of unusual register instructions.

Each instruction consists of several parts : the register to modify, whether to increase or decrease that register's value, the amount by which to increase or decrease it, and a condition. If the condition fails, skip the instruction without modifying the register. The registers all start at 0. The instructions look like this:

b inc 5 if a > 1
a inc 1 if b < 5
    c dec - 10 if a >= 1
    c inc - 20 if c == 10

    These instructions would be processed as follows :

Because a starts at 0, it is not greater than 1, and so b is not modified.
a is increased by 1 (to 1) because b is less than 5 (it is 0).
c is decreased by - 10 (to 10) because a is now greater than or equal to 1 (it is 1).
c is increased by - 20 (to - 10) because c is equal to 10.

After this process, the largest value in any register is 1.

You might also encounter <= (less than or equal to) or != (not equal to).However, the CPU doesn't have the bandwidth to tell you what all the registers are named, and leaves that to you to determine.

What is the largest value in any register after completing the instructions in your puzzle input ?
*/

constexpr char MAX_AT_ANY_TIME[] = "__MaxSeenAtAnyTime";

std::map<std::string, int> ComputeRegisters(const std::vector<std::string>& commands)
{
    enum TokenName {
        TARGET = 0, // The register we're acting on
        ACTION,     // What we're doing to that register (inc/dec)
        AMOUNT,     // The amount we're adjusting the register by
        IF_TOKEN,   // Must be "if"
        DEPENDENT,  // The register that our action depends on
        COMPARISON, // The comparison operator
        REFERENCE,  // The value we're comparing the dependent to
    };

    std::map<std::string, int> registers;
    int maxSeenAtAnyTime = INT_MIN;

    for (auto &line : commands) {
        auto tokens = Tokenize(line);
        if (tokens.size() != 7) std::cerr << "Unexpected token count: " << line << std::endl;

        // Ensure that we have registers for the names
        if (registers.find(tokens[TARGET]) == registers.end()) registers[tokens[TARGET]] = 0;
        if (registers.find(tokens[DEPENDENT]) == registers.end()) registers[tokens[DEPENDENT]] = 0;

        if (tokens[IF_TOKEN] != "if") std::cerr << "Unexpected if token: " << line << std::endl;
        int reference = atoi(tokens[REFERENCE].c_str());

        bool doAction = false;
        if (tokens[COMPARISON] == ">=") {
            doAction = registers[tokens[DEPENDENT]] >= reference;
        }
        else if (tokens[COMPARISON] == "<=") {
            doAction = registers[tokens[DEPENDENT]] <= reference;
        }
        else if (tokens[COMPARISON] == "<") {
            doAction = registers[tokens[DEPENDENT]] < reference;
        }
        else if (tokens[COMPARISON] == ">") {
            doAction = registers[tokens[DEPENDENT]] > reference;
        }
        else if (tokens[COMPARISON] == "==") {
            doAction = registers[tokens[DEPENDENT]] == reference;
        }
        else if (tokens[COMPARISON] == "!=") {
            doAction = registers[tokens[DEPENDENT]] != reference;
        }
        else {
            std::cerr << "Unexpected comparison operator: " << line << std::endl;
        }

        if (doAction) {
            int delta = atoi(tokens[AMOUNT].c_str());

            if (tokens[ACTION] == "inc") {
                registers[tokens[TARGET]] += delta;
            } else if (tokens[ACTION] == "dec") {
                registers[tokens[TARGET]] -= delta;
            } else {
                std::cerr << "Unexpected action operator: " << line << std::endl;
            }

            maxSeenAtAnyTime = std::max(maxSeenAtAnyTime, registers[tokens[TARGET]]);
        }
    }

    registers[MAX_AT_ANY_TIME] = maxSeenAtAnyTime;

    return registers;
}

std::pair<int, int> MaxRegister(const std::vector<std::string>& commands)
{
    auto registers = ComputeRegisters(commands);
    int maxValueSeen = INT_MIN;
    for (auto& reg : registers) {
        if (reg.first != MAX_AT_ANY_TIME) {
            maxValueSeen = std::max(reg.second, maxValueSeen);
        }
    }
    return std::make_pair(maxValueSeen, registers[MAX_AT_ANY_TIME]);
}

void Day8Tests()
{
    std::vector<std::string> testInput = {
        "b inc 5 if a > 1",
        "a inc 1 if b < 5",
        "c dec -10 if a >= 1",
        "c inc -20 if c == 10"
    };

    auto maxValue = MaxRegister(testInput);
    if (maxValue.first != 1) std::cerr << "Test 8A Error: Got " << maxValue.first << ", Expected 1" << std::endl;
    if (maxValue.second != 10) std::cerr << "Test 8B Error: Got " << maxValue.second << ", Expected 10" << std::endl;
}

void Day8()
{
    std::cout << "Day 8:\n";
    Day8Tests();
    auto input = ReadFileLines("input_day8.txt");
    auto maxRegister = MaxRegister(input);
    std::cout << maxRegister.first << std::endl << maxRegister.second << std::endl << std::endl;
}