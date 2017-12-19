#include "Problems.h"
/*
--- Day 19: A Series of Tubes ---

Somehow, a network packet got lost and ended up here. It's trying to follow a routing diagram (your puzzle input), 
but it's confused about where to go.

Its starting point is just off the top of the diagram. Lines (drawn with |, -, and +) show the path it needs to take, 
starting by going down onto the only line connected to the top of the diagram. It needs to follow this path until it 
reaches the end (located somewhere within the diagram) and stop there.

Sometimes, the lines cross over each other; in these cases, it needs to continue going the same direction, 
and only turn left or right when there's no other option. In addition, someone has left letters on the line; 
these also don't change its direction, but it can use them to keep track of where it's been. For example:

     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+

Given this diagram, the packet needs to take the following path:

Starting at the only line touching the top of the diagram, it must go down, pass through A, and continue onward to the first +.
Travel right, up, and right, passing through B in the process.
Continue down (collecting C), right, and up (collecting D).
Finally, go all the way left through E and stopping at F.

Following the path to the end, the letters it sees on its path are ABCDEF.

The little packet looks up at you, hoping you can help it find the way. What letters will it see (in the order it would see them) 
if it follows the path? (The routing diagram is very wide; make sure you view it without line wrapping.)

--- Part Two ---

The packet is curious how many steps it needs to go.

For example, using the same routing diagram from the example above...

     |          
     |  +--+    
     A  |  C    
 F---|--|-E---+ 
     |  |  |  D 
     +B-+  +--+ 

...the packet would go:

    6 steps down (including the first line at the top of the diagram).
    3 steps right.
    4 steps up.
    3 steps right.
    4 steps down.
    3 steps right.
    2 steps up.
    13 steps left (including the F it stops on).

This would result in a total of 38 steps.

How many steps does the packet need to go?
*/
namespace Day19 {

class Maze
{
public:
    Maze(std::vector<std::string> input);
    void SetDataHandler(std::function<void(char)> handler) { m_dataHandler = handler; } 
    int Solve();   // Tries to walk a solution to the maze, calling the data handler as nodes are visited

protected:
    // Each character in the input represents either a path, a line, or a letter that we should record (data).
    enum class Block {
        Empty,
        Line,
        Data
    };

    enum class Direction {
        Up = 0,
        Right = 1,
        Down = 2,
        Left = 3
    };
    Direction TurnLeft(Direction now) const { return static_cast<Direction>((static_cast<int>(now) + 3) % 4); }
    Direction TurnRight(Direction now) const { return static_cast<Direction>((static_cast<int>(now) + 1) % 4); }
    Direction ReverseDirection(Direction now) const { return static_cast<Direction>((static_cast<int>(now) + 2) % 4); }

    struct Cell {
        Cell(int r = 0, int c = 0) : row(r), col(c) {}
        int row;
        int col;

        void Move(Direction dir) {
            switch (dir) {
            case Direction::Left:
                --col;
                break;
            case Direction::Right:
                ++col;
                break;
            case Direction::Up:
                --row;
                break;
            case Direction::Down:
                ++row;
                break;
            }
        }
    };

    Cell FindStart() const;
    inline constexpr Maze::Block Classify(char c) const;
    bool OutOfBounds(const Cell& cell) const;
    inline char ReadCell(const Cell& cell) const;

    std::vector<std::string> m_maze;
    std::function<void(char)> m_dataHandler;
    Cell m_position;
    Direction m_direction;
};

Maze::Maze(std::vector<std::string> input) {
    m_maze = std::move(input);
    m_direction = Direction::Down;
    m_position = FindStart();
}

inline constexpr Maze::Block Maze::Classify(char c) const {
    switch (c) {
    case ' ':
        return Block::Empty;
    case '|':
    case '-':
    case '+':
        return Block::Line;
    default:
        return Block::Data;
    }
}

Maze::Cell Maze::FindStart() const
{
    if (m_maze.size() >= 1) {
        // The starting point is guaranteed to be in the first row
        for (int column = 0; column < m_maze[0].size(); ++column) {
            if (Classify(m_maze[0][column]) != Block::Empty) return Cell(0, column);
        }
    }
    std::cerr << "Invalid Maze\n";
    return Cell();
}

bool Maze::OutOfBounds(const Cell& cell) const
{
    return (cell.col < 0) || (cell.row < 0) || (cell.row >= m_maze.size()) || (cell.col >= m_maze[0].size());
}

inline char Maze::ReadCell(const Maze::Cell& cell) const
{
    return m_maze[cell.row][cell.col];
}

int Maze::Solve()
{
    int totalDistance = 0;
    bool solved = false;
    while (!solved) {
        ++totalDistance;

        // If our current cell has data, we should read it
        char current = ReadCell(m_position);
        if (m_dataHandler && (Classify(current) == Block::Data)) {
            m_dataHandler(current);
        }

        // First, try to continue in the current direction
        auto newPosition = m_position;
        auto newDirection = m_direction;
        newPosition.Move(m_direction);
        if (OutOfBounds(newPosition) || (Classify(ReadCell(newPosition)) == Block::Empty)) {
            // Next, try to the left
            newPosition = m_position;
            newDirection = TurnLeft(m_direction);
            newPosition.Move(newDirection);
            if (OutOfBounds(newPosition) || (Classify(ReadCell(newPosition)) == Block::Empty)) {
                // Finally, try to the right
                newPosition = m_position;
                newDirection = TurnRight(m_direction);
                newPosition.Move(newDirection);
                if (OutOfBounds(newPosition) || (Classify(ReadCell(newPosition)) == Block::Empty)) {
                    // We have nowhere else to go.  We must be at the end.
                    solved = true;
                }
            }
        }
        m_position = newPosition;
        m_direction = newDirection;
    }

    return totalDistance;
}

std::pair<std::string, int> PathText(const std::vector<std::string>& input)
{
    std::stringstream path;
    Maze maze(input);
    maze.SetDataHandler([&path](char c) {
        path << c;
    });
    auto distance = maze.Solve();
    return std::make_pair(path.str(), distance);
}

} // namespace Day19

void Day19Tests()
{
    const std::vector<std::string> input = {
        "     |          ",
        "     |  +--+    ",
        "     A  |  C    ",
        " F---|----E|--+ ",
        "     |  |  |  D ",
        "     +B-+  +--+ ",
        "                "
    };
    const auto pathText = Day19::PathText(input);
    const std::string text = "ABCDEF";
    const int distance = 38;
    if (pathText.first != text) std::cerr << "Test 19A Error: Got " << pathText.first << ", expected " << text << std::endl;
    if (pathText.second != distance) std::cerr << "Test 19B Error: Got " << pathText.second << ", expected " << distance << std::endl;
}

void Day19Problems()
{
    std::cout << "Day 19:\n";
    Day19Tests();
    const auto start = std::chrono::steady_clock::now();
    const auto input = Helpers::ReadFileLines("input_day19.txt");
    const auto pathText = Day19::PathText(input);
    const auto end = std::chrono::steady_clock::now();
    std::cout << pathText.first << std::endl << pathText.second << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl << std::endl;
}