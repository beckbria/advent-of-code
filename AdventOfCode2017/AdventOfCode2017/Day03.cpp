#include "Problems.h"

/*
You come across an experimental new kind of memory stored on an infinite two-dimensional grid.

Each square on the grid is allocated in a spiral pattern starting at a location marked 1 and then counting up while spiraling outward. For example, the first few squares are allocated like this:

17  16  15  14  13
18   5   4   3  12
19   6   1   2  11
20   7   8   9  10
21  22  23---> ...

While this is very space-efficient (no squares are skipped), requested data must be carried back to square 1 (the location of the only access port for this memory system) by programs that can only move up, down, left, or right. They always take the shortest path: the Manhattan Distance between the location of the data and square 1.

For example:

Data from square 1 is carried 0 steps, since it's at the access port.
Data from square 12 is carried 3 steps, such as: down, left, left.
Data from square 23 is carried only 2 steps: up twice.
Data from square 1024 must be carried 31 steps.

How many steps are required to carry the data from the square identified in your puzzle input all the way to the access port?

Your puzzle input is 347991.
*/

int SpiralDistance(int target)
{
    if (target < 1) return -1;

    // Notes: Let's call the space containing 1 (0,0).  Thus, every positive space (n,n) contains (2*n+1)^2.
    // At the least, we could find the nearest odd square and count from there

    int targetRoot = static_cast<int>(ceil(sqrt(target)));
    // We want to start from an odd square
    if (targetRoot % 2 == 0) {
        ++targetRoot;
    }

    int const cornerCoordinate = (targetRoot - 1) / 2;
    int const distanceBetweenCorners = 2 * cornerCoordinate;
    int currentX = cornerCoordinate;    // Start in the lower right corner
    int currentY = cornerCoordinate;

    int currentValue = targetRoot * targetRoot;
    int delta = currentValue - target;

    // If we're honest, a simple brute force countback SHOULD be fast enough here.
    // We can do better, though

    // First, we move left along the bottom edge
    if (delta > distanceBetweenCorners) {
        // Skip the entire bottom row
        currentX = -cornerCoordinate;
        currentValue -= distanceBetweenCorners;
        delta -= distanceBetweenCorners;

        // First, we move left along the left edge
        if (delta > distanceBetweenCorners) {
            // Skip the entire bottom row
            currentY = -cornerCoordinate;
            currentValue -= distanceBetweenCorners;
            delta -= distanceBetweenCorners;

            // Then the top edge
            if (delta > distanceBetweenCorners) {
                // Skip the entire bottom row
                currentX = cornerCoordinate;
                currentValue -= distanceBetweenCorners;
                delta -= distanceBetweenCorners;

                // Finally the right edge
                if (delta > distanceBetweenCorners) {
                    std::cerr << "SpiralDistance Error, test case " << target << std::endl;
                }
                else {
                    // Step down the right edge
                    currentX += delta;
                }
            }
            else {
                // Step right along the top edge
                currentX += delta;
            }
        }
        else {
            // Step up the left edge
            currentY -= delta;
        }
    } else {
        // Step left along the bottom edge
        currentX -= delta;
    }

    return abs(currentX) + abs(currentY);
}

/*
As a stress test on the system, the programs here clear the grid and then store the value 1 in square 1. Then, in the same allocation order as shown above, they store the sum of the values in all adjacent squares, including diagonals.

So, the first few squares' values are chosen as follows:

Square 1 starts with the value 1.
Square 2 has only one adjacent filled square (with value 1), so it also stores 1.
Square 3 has both of the above squares as neighbors and stores the sum of their values, 2.
Square 4 has all three of the aforementioned squares as neighbors and stores the sum of their values, 4.
Square 5 only has the first and fourth squares as neighbors, so it gets the value 5.

Once a square is written, its value does not change. Therefore, the first few squares would receive the following values:

147  142  133  122   59
304    5    4    2   57
330   10    1    1   54
351   11   23   25   26
362  747  806--->   ...

What is the first value written that is larger than your puzzle input?

Your puzzle input is still 347991.
*/

enum class Direction {
    Up, Left, Right, Down
};

int FirstHigherSpiral(int target)
{
    // We use a grid where the "center" is at 126,126 to give us a range from (-125,-125) through (125,125)
    int constexpr centerOffset = 126;
    int grid[(2 * centerOffset) - 1][(2 * centerOffset - 1)] = {};
    grid[centerOffset][centerOffset] = 1;

    int currentX = 1;
    int currentY = 0;
    int squareCoordinate = 1;
    Direction currentDirection = Direction::Up;
    int lastValueWritten = 1;

    while (lastValueWritten < target) {
        // Fill in the current grid cell
        int X = currentX + centerOffset;
        int Y = currentY + centerOffset;
        grid[X][Y] = grid[X-1][Y-1] + grid[X][Y-1] + grid[X+1][Y-1] +
                     grid[X-1][Y]   +                grid[X+1][Y] +
                     grid[X-1][Y+1] + grid[X][Y+1] + grid[X+1][Y+1];
        lastValueWritten = grid[X][Y];

        // Move to the next grid cell
        switch (currentDirection) {
        case Direction::Up:
            currentY--;
            if (abs(currentY) == squareCoordinate) {
                currentDirection = Direction::Left;
            }
            break;

        case Direction::Left:
            currentX--;
            if (abs(currentX) == squareCoordinate) {
                currentDirection = Direction::Down;
            }
            break;

        case Direction::Down:
            currentY++;
            if (currentY == squareCoordinate) {
                currentDirection = Direction::Right;
            }
            break;

        case Direction::Right:
            currentX++;
            // When we reach the end of our square, we don't change direction - we go right for an additional space
            // to start the next square
            if (currentX == (squareCoordinate + 1)) {
                squareCoordinate++;
                currentDirection = Direction::Up;
            }
            break;
        }
    }

    return lastValueWritten;
}

void Day3Tests()
{
    struct {
        int input;
        int answer;
    } testCaseA[] = { {1,0}, {12,3}, {23,2}, {1024,31} };

    for (auto &t : testCaseA) {
        int result = SpiralDistance(t.input);
        if (result != t.answer) {
            std::cerr << "Test 3A failed: " << t.input << " => " << result << " (expected " << t.answer << ")" << std::endl;
        }
    }

    struct {
        int input;
        int answer;
    } testCaseB[] = { { 4,5 },{ 20,23 },{ 100,122 },{ 200,304 },{500,747} };

    for (auto &t : testCaseA) {
        int result = FirstHigherSpiral(t.input);
        if (result != t.answer) {
            std::cerr << "Test 3B failed: " << t.input << " => " << result << " (expected " << t.answer << ")" << std::endl;
        }
    }
}

void Day3()
{
    int const target = 347991;
    std::cout << "Day 3:\n";
    std::cout << SpiralDistance(target) << std::endl;
    std::cout << FirstHigherSpiral(target) << std::endl << std::endl;
}