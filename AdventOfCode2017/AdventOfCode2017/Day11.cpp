#include "Problems.h"
/*
--- Day 11: Hex Ed ---

Crossing the bridge, you've barely reached the other side of the stream when a program comes up to you, clearly in distress.
"It's my child process," she says, "he's gotten lost in an infinite grid!"

Fortunately for her, you have plenty of experience with infinite grids.

Unfortunately for you, it's a hex grid.

The hexagons ("hexes") in this grid are aligned such that adjacent hexes can be found to the north, northeast,
southeast, south, southwest, and northwest:

  \ n  /
nw +--+ ne
  /    \
-+      +-
  \    /
sw +--+ se
  / s  \

You have the path the child process took. Starting where he started, you need to determine the fewest number of steps
required to reach him. (A "step" means to move from the hex you are in to any adjacent hex.)

For example:

    ne,ne,ne is 3 steps away.
    ne,ne,sw,sw is 0 steps away (back where you started).
    ne,ne,s,s is 2 steps away (se,se).
    se,sw,se,sw,sw is 3 steps away (s,s,sw).

--- Part Two ---

How many steps away is the furthest he ever got from his starting position?
*/

std::map<std::string, int> Count(const std::vector<std::string>& directions)
{
    std::map<std::string, int> wordCount;
    for (auto& dir : directions) {
        if (wordCount.count(dir) == 0) {
            wordCount[dir] = 1;
        } else {
            wordCount[dir]++;
        }
    }
    return wordCount;
}

// Returns true if it changed anything
bool NegateOpposites(int& North, int& South) {
    if ((North > 0) && (South > 0)) {
        const int min = std::min(North, South);
        North -= min;
        South -= min;
        return true;
    }
    return false;
}

// Returns true if anything changed.  Pairs of the first two variables are removed and added to the third (thus, SW+SE becomes S)
bool CombineDirections(int& SouthWest, int& SouthEast, int& South)
{
    if ((SouthWest > 0) && (SouthEast > 0))
    {
        const int min = std::min(SouthWest, SouthEast);
        SouthWest -= min;
        SouthEast -= min;
        South += min;
        return true;
    }
    return false;
}

// This function reduces a count of hex directions to their minimum
void OptimizeHexDirections(int& N, int& NW, int& NE, int& S, int& SW, int& SE)
{
    bool performedOptimization;
    do {
        performedOptimization = false;

        // First, negate opposite directions
        performedOptimization |= NegateOpposites(NW, SE);
        performedOptimization |= NegateOpposites(N, S);
        performedOptimization |= NegateOpposites(NE, SW);

        // Next, combine west/east into a single direction
        performedOptimization |= CombineDirections(SW, SE, S);
        performedOptimization |= CombineDirections(S, NE, SE);
        performedOptimization |= CombineDirections(SE, N, NE);
        performedOptimization |= CombineDirections(NE, NW, N);
        performedOptimization |= CombineDirections(N, SW, NW);
        performedOptimization |= CombineDirections(NW, S, SW);
    } while (performedOptimization);
}

int StepsAway(const std::vector<std::string>& directions)
{
    auto stepCount = Count(directions);
    int NW = stepCount["nw"];
    int NE = stepCount["ne"];
    int N = stepCount["n"];
    int SW = stepCount["sw"];
    int SE = stepCount["se"];
    int S = stepCount["s"];

    OptimizeHexDirections(N, NW, NE, S, SW, SE);
    return N + NW + NE + S + SW + SE;
}

int MaxStepsEverAway(const std::vector<std::string>& directions)
{
    int N = 0, NW = 0, NE = 0, S = 0, SW = 0, SE = 0;
    int furthestAway = 0;
    for (auto &dir : directions) {
        if (dir == "s") ++S;
        else if (dir == "sw") ++SW;
        else if (dir == "se") ++SE;
        else if (dir == "n") ++N;
        else if (dir == "nw") ++NW;
        else if (dir == "ne") ++NE;

        OptimizeHexDirections(N, NW, NE, S, SW, SE);
        const int currentDistance = N + NW + NE + S + SW + SE;
        furthestAway = std::max(furthestAway, currentDistance);
    }
    return furthestAway;
}

void Day11Tests()
{
    const struct {
        std::vector<std::string> input;
        int answer;
    } testCasesA[] = {
        { {"ne","ne","ne"}, 3 },
        { { "ne","ne","sw","sw" }, 0 },
        { { "ne","ne","s","s" }, 2 },
        { { "se","sw","se","sw","sw" }, 3 },
    };
    for (auto &test : testCasesA) {
        int result = StepsAway(test.input);
        if (result != test.answer) std::cerr << "Test 11A Failed: Got " << result << ", Expected " << test.answer << std::endl;
    }
}

void Day11()
{
    auto source = ReadFileLines("input_day11.txt");
    auto input = Tokenize(source[0], ',');
    std::cout << "Day 11:\n";
    Day11Tests();
    std::cout << StepsAway(input) << std::endl << MaxStepsEverAway(input) << std::endl << std::endl;
}