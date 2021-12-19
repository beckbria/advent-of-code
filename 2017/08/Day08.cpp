#include "2017/lib/Helpers.h"
/*
--- Day 8: I Heard You Like Registers ---
You receive a signal directly from the CPU. Because of your recent assistance with jump instructions, 
it would like you to compute the result of a series of unusual register instructions.

Each instruction consists of several parts : the register to modify, whether to increase or decrease 
that register's value, the amount by which to increase or decrease it, and a condition. If the 
condition fails, skip the instruction without modifying the register. The registers all start at 0. 
The instructions look like this:

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

You might also encounter <= (less than or equal to) or != (not equal to).  However, the CPU doesn't have the 
bandwidth to tell you what all the registers are named, and leaves that to you to determine.

What is the largest value in any register after completing the instructions in your puzzle input ?

--- Part Two ---

To be safe, the CPU also needs to know the highest value held in any register during this process so that it 
can decide how much memory to allocate to these operations. For example, in the above instructions, the 
highest value ever held was 10 (in register c after the third instruction was evaluated).
*/
namespace Day8
{
    std::pair<std::map<std::string, int>, int> ComputeRegisters(const std::vector<std::string> &commands)
    {
        enum TokenName
        {
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

        for (auto &line : commands)
        {
            auto tokens = Helpers::Tokenize(line);
            if (tokens.size() != 7)
                std::cerr << "Unexpected token count: " << line << std::endl;

            // Ensure that we have registers for the names
            if (registers.find(tokens[TARGET]) == registers.end())
                registers[tokens[TARGET]] = 0;
            if (registers.find(tokens[DEPENDENT]) == registers.end())
                registers[tokens[DEPENDENT]] = 0;

            if (tokens[IF_TOKEN] != "if")
                std::cerr << "Unexpected if token: " << line << std::endl;
            int reference = std::stoi(tokens[REFERENCE]);

            bool doAction = false;
            if (tokens[COMPARISON] == ">=")
            {
                doAction = registers[tokens[DEPENDENT]] >= reference;
            }
            else if (tokens[COMPARISON] == "<=")
            {
                doAction = registers[tokens[DEPENDENT]] <= reference;
            }
            else if (tokens[COMPARISON] == "<")
            {
                doAction = registers[tokens[DEPENDENT]] < reference;
            }
            else if (tokens[COMPARISON] == ">")
            {
                doAction = registers[tokens[DEPENDENT]] > reference;
            }
            else if (tokens[COMPARISON] == "==")
            {
                doAction = registers[tokens[DEPENDENT]] == reference;
            }
            else if (tokens[COMPARISON] == "!=")
            {
                doAction = registers[tokens[DEPENDENT]] != reference;
            }
            else
            {
                std::cerr << "Unexpected comparison operator: " << line << std::endl;
            }

            if (doAction)
            {
                int delta = std::stoi(tokens[AMOUNT]);

                if (tokens[ACTION] == "inc")
                {
                    registers[tokens[TARGET]] += delta;
                }
                else if (tokens[ACTION] == "dec")
                {
                    registers[tokens[TARGET]] -= delta;
                }
                else
                {
                    std::cerr << "Unexpected action operator: " << line << std::endl;
                }
                maxSeenAtAnyTime = std::max(maxSeenAtAnyTime, registers[tokens[TARGET]]);
            }
        }

        return std::make_pair(registers, maxSeenAtAnyTime);
    }

    std::pair<int, int> MaxRegister(const std::vector<std::string> &commands)
    {
        auto registers = ComputeRegisters(commands);
        int maxValueSeen = INT_MIN;
        for (auto &reg : registers.first)
        {
            maxValueSeen = std::max(reg.second, maxValueSeen);
        }
        return std::make_pair(maxValueSeen, registers.second);
    }
} // namespace Day8

void Day8Tests()
{
    const std::vector<std::string> testInput = {
        "b inc 5 if a > 1",
        "a inc 1 if b < 5",
        "c dec -10 if a >= 1",
        "c inc -20 if c == 10"};

    const auto maxValue = Day8::MaxRegister(testInput);
    if (maxValue.first != 1)
        std::cerr << "Test 8A Error: Got " << maxValue.first << ", Expected 1" << std::endl;
    if (maxValue.second != 10)
        std::cerr << "Test 8B Error: Got " << maxValue.second << ", Expected 10" << std::endl;
}

void Day8Problems()
{
    std::cout << "Day 8:\n";
    Day8Tests();
    const auto start = std::chrono::steady_clock::now();
    const auto input = Helpers::ReadFileLines("2017/08/input_day8.txt");
    const auto maxRegister = Day8::MaxRegister(input);
    const auto end = std::chrono::steady_clock::now();
    std::cout << maxRegister.first << std::endl
              << maxRegister.second << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl
              << std::endl;
}

int main()
{
    Day8Problems();
    return 0;
}