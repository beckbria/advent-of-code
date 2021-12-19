#include "2017/lib/Helpers.h"
/*
-- Day 16: Permutation Promenade ---

You come upon a very unusual sight; a group of programs here appear to be dancing.

There are sixteen programs in total, named a through p. They start by standing in a line: a stands in position 0, b stands in position 1, and so on until p, which stands in position 15.

The programs' dance consists of a sequence of dance moves:

    Spin, written sX, makes X programs move from the end to the front, but maintain their order otherwise. (For example, s3 on abcde produces cdeab).
    Exchange, written xA/B, makes the programs at positions A and B swap places.
    Partner, written pA/B, makes the programs named A and B swap places.

For example, with only five programs standing in a line (abcde), they could do the following dance:

    s1, a spin of size 1: eabcd.
    x3/4, swapping the last two programs: eabdc.
    pe/b, swapping programs e and b: baedc.

After finishing their dance, the programs end up in order baedc.

You watch the dance for a while and record their dance moves (your puzzle input). In what order are the programs standing after their dance?

--- Part Two ---

Now that you're starting to get a feel for the dance moves, you turn your attention to the dance as a whole.

Keeping the positions they ended up in from their previous dance, the programs perform it again and again: including the first dance, a total of one billion (1000000000) times.

In the example above, their second dance would begin with the order baedc, and use the same dance moves:

s1, a spin of size 1: cbaed.
x3/4, swapping the last two programs: cbade.
pe/b, swapping programs e and b: ceadb.

In what order are the programs standing after their billion dances?
*/
namespace Day16
{

    // Spin, written sX, makes X programs move from the end to the front, but maintain their order otherwise. (For example, s3 on abcde produces cdeab).
    void Spin(std::string &programs, std::string &action)
    {
        const int shift = std::stoi(action) % programs.size();
        std::rotate(programs.begin(), programs.end() - shift, programs.end());
    }

    // Exchange, written xA/B, makes the programs at positions A and B swap places.
    void Exchange(std::string &programs, std::string &action)
    {
        const auto positions = Helpers::Tokenize(action, '/');
        const int first = std::stoi(positions[0]);
        const int second = std::stoi(positions[1]);
        std::swap(programs[first], programs[second]);
    }

    // Partner, written pA/B, makes the programs named A and B swap places.
    void Partner(std::string &programs, std::string &action)
    {
        const auto toSwap = Helpers::Tokenize(action, '/');
        const char first = toSwap[0][0];
        const char second = toSwap[1][0];

        // Swap the two characters
        for (char &c : programs)
        {
            if (c == first)
            {
                c = second;
            }
            else if (c == second)
            {
                c = first;
            }
        }
    }

    void RunCommands(std::string &programs, const std::vector<std::string> &commands, unsigned int iterations = 1)
    {
        // An entry stored in this table means that we saw a certain state at the end of iteration X.  The initial state give to us is iteration -1.
        std::unordered_map<std::string, int> knownProgramStates;
        knownProgramStates[programs] = -1;
        bool shortcutApplied = false; // Have we detected a cycle and jumped ahead?  If so, stop looking for cycles

        for (unsigned int i = 0; i < iterations; ++i)
        {
            // Run the program
            for (const auto &command : commands)
            {
                // Split the command into the type and the actual instructions
                char commandType = command[0];
                std::string actions = command.substr(1);

                switch (commandType)
                {
                case 's':
                    Spin(programs, actions);
                    break;

                case 'x':
                    Exchange(programs, actions);
                    break;

                case 'p':
                    Partner(programs, actions);
                    break;
                }
            }

            // Store our output.  If we've ever seen it before, we have a cycle and can shortcut
            if (!shortcutApplied)
            {
                if (knownProgramStates.count(programs) > 0)
                {
                    int cycleLength = i - knownProgramStates[programs];
                    while (i < iterations)
                        i += cycleLength;
                    // We've gone too far, so backtrack one cycle.
                    i -= cycleLength;
                    shortcutApplied = true;
                }
                else
                {
                    knownProgramStates[programs] = i;
                }
            }
        }
    }

} // namespace Day16

void Day16Tests()
{
    std::string dance = "abcde";
    std::vector<std::string> commands = {"s1", "x3/4", "pe/b"};
    Day16::RunCommands(dance, commands);
    const std::string expected = "baedc";
    if (dance != expected)
        std::cerr << "Test 16A Error: Got " << dance << ", expected " << expected << std::endl;

    std::string doubleDance = "abcde";
    Day16::RunCommands(doubleDance, commands, 2);
    const std::string expectedTwoTimes = "ceadb";
    if (doubleDance != expectedTwoTimes)
        std::cerr << "Test 16B Error: Got " << doubleDance << ", expected " << expectedTwoTimes << std::endl;
}

void Day16Problems()
{
    std::cout << "Day 16:\n";
    Day16Tests();
    const auto start = std::chrono::steady_clock::now();
    auto input = Helpers::ReadFile<std::string>("2017/16/input_day16.txt");
    auto commands = Helpers::Tokenize(input[0], ',', false);
    std::string programs = "abcdefghijklmnop";
    std::string billionRuns = programs;
    Day16::RunCommands(programs, commands);
    Day16::RunCommands(billionRuns, commands, 1000000000);
    const auto end = std::chrono::steady_clock::now();
    std::cout << programs << std::endl;
    std::cout << billionRuns << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl
              << std::endl;
}

int main()
{
    Day16Problems();
    return 0;
}