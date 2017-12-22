#include "Problems.h"
/*
--- Day 22: Sporifica Virus ---

Diagnostics indicate that the local grid computing cluster has been contaminated with the Sporifica Virus. 
The grid computing cluster is a seemingly-infinite two-dimensional grid of compute nodes. Each node is either 
clean or infected by the virus.

To prevent overloading the nodes (which would render them useless to the virus) or detection by system 
administrators, exactly one virus carrier moves through the network, infecting or cleaning nodes as it 
moves. The virus carrier is always located on a single node in the network (the current node) and keeps 
track of the direction it is facing.

To avoid detection, the virus carrier works in bursts; in each burst, it wakes up, does some work, and 
goes back to sleep. The following steps are all executed in order one time each burst:

    - If the current node is infected, it turns to its right. Otherwise, it turns to its left. 
      (Turning is done in-place; the current node does not change.)
    - If the current node is clean, it becomes infected. Otherwise, it becomes cleaned. 
      (This is done after the node is considered for the purposes of changing direction.)
    - The virus carrier moves forward one node in the direction it is facing.

Diagnostics have also provided a map of the node infection status (your puzzle input). 
Clean nodes are shown as .; infected nodes are shown as #. This map only shows the 
center of the grid; there are many more nodes beyond those shown, but none of them 
are currently infected.

The virus carrier begins in the middle of the map facing up.

For example, suppose you are given a map like this:

..#
#..
...

Then, the middle of the infinite grid looks like this, with the virus carrier's position marked with [ ]:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . # . . .
. . . #[.]. . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .

The virus carrier is on a clean node, so it turns left, infects the node, and moves left:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . # . . .
. . .[#]# . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .

The virus carrier is on an infected node, so it turns right, cleans the node, and moves up:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . .[.]. # . . .
. . . . # . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .

Four times in a row, the virus carrier finds a clean, infects it, turns left, and moves forward, ending in the same place and still facing up:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . #[#]. # . . .
. . # # # . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .

Now on the same node as before, it sees an infection, which causes it to turn right, clean the node, and move forward:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . # .[.]# . . .
. . # # # . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .

After the above actions, a total of 7 bursts of activity had taken place. Of them, 5 bursts of activity caused an infection.

After a total of 70, the grid looks like this, with the virus carrier facing up:

. . . . . # # . .
. . . . # . . # .
. . . # . . . . #
. . # . #[.]. . #
. . # . # . . # .
. . . . . # # . .
. . . . . . . . .
. . . . . . . . .

By this time, 41 bursts of activity caused an infection (though most of those nodes have since been cleaned).

After a total of 10000 bursts of activity, 5587 bursts will have caused an infection.

Given your actual map, after 10000 bursts of activity, how many bursts cause a node to become infected? (Do not count nodes that begin infected.)

*/
namespace Day22 {

struct Point {
    Point(int64_t X = 0, int64_t Y = 0) : x(X), y(Y) {}
    int64_t x;
    int64_t y;
};

bool operator==(const Point& left, const Point& right);

class PointHash {
public:
    std::size_t operator() (const Point& p) const
    {
        // This isn't a particularly good hash.  A much better choice would be something like boost::hash_combine.
        // I'm avoiding doing that here because I don't want to pull in a massive set of dependencies and licensing
        // constraints.
        return std::hash<uint64_t>()(p.x) + std::hash<uint64_t>()(p.y);
    }
};

class VirusCarrier {
public:
    Point Position() const { return m_position; }
    void RotateLeft() { m_direction = static_cast<Direction>((static_cast<int>(m_direction) + 3) % 4); }
    void RotateRight() { m_direction = static_cast<Direction>((static_cast<int>(m_direction) + 1) % 4); }
    void MoveForward();

private:
    enum class Direction {
        Up = 0,
        Right = 1,
        Down = 2,
        Left = 3
    };

    Point m_position;   // The carrier always starts at (0,0)
    Direction m_direction = Direction::Up;
};

// Plan: Store the infected nodes in a hash table.  Don't store the entire grid.  This lets us scale indefinitely
class VirusGrid {
public:
    VirusGrid(const std::vector<std::string>& input);
    void AdvanceToTurn(unsigned int turn);
    uint64_t TurnsThatInfectedNode() { return m_infectedTurns; }
    friend std::ostream& operator<<(std::ostream& out, const VirusGrid& grid);

private:
    void AdvanceTurn();
    inline bool IsInfected(const Point& position) const;

    std::unordered_set<Point, PointHash> m_infected;
    VirusCarrier m_carrier;
    uint64_t m_infectedTurns = 0;
    uint64_t m_currentTurn = 0;
    static constexpr char InfectedNode = '#';
    static constexpr char CleanNode = '.';
};

bool operator==(const Point& left, const Point& right)
{
    return (left.x == right.x) && (left.y == right.y);
}

void VirusCarrier::MoveForward()
{
    switch (m_direction) {
    case Direction::Up:
        m_position.y--;
        break;
    case Direction::Down:
        m_position.y++;
        break;
    case Direction::Left:
        m_position.x--;
        break;
    case Direction::Right:
        m_position.x++;
    }
}

VirusGrid::VirusGrid(const std::vector<std::string>& input)
{
    if (input.size() > 0) {
        Point middle((input[0].size() / 2), (input.size() / 2));    // The +1 to get to the center of odd shapes is implied by 0 indexing
        for (int y = 0; y < (int)input.size(); ++y) {
            for (int x = 0; x < (int)input[0].size(); ++x) {
                if (input[y][x] == InfectedNode) {
                    m_infected.emplace(x - middle.x, y - middle.y);
                }
            }
        }
    }
}

void VirusGrid::AdvanceToTurn(unsigned int turn)
{
    while (m_currentTurn < static_cast<uint64_t>(turn)) {
        AdvanceTurn();
    }
}

void VirusGrid::AdvanceTurn()
{
    auto currentPosition = m_carrier.Position();
    if (IsInfected(currentPosition)) {
        m_infected.erase(currentPosition);  // Clean the current node
        m_carrier.RotateRight();
    } else {
        m_infected.insert(currentPosition); // Infect the current node
        ++m_infectedTurns;
        m_carrier.RotateLeft();
    }
    m_carrier.MoveForward();
    ++m_currentTurn;
}

inline bool VirusGrid::IsInfected(const Point& position) const
{
    return m_infected.count(position) > 0;
}

std::ostream& operator<<(std::ostream& out, const VirusGrid& grid)
{
    if (grid.m_infected.size() > 0) {
        // Find the corners
        int64_t minX = INT64_MAX, minY = INT64_MAX, maxX = INT64_MIN, maxY = INT64_MIN;
        for (const auto &point : grid.m_infected) {
            minX = std::min(minX, point.x);
            minY = std::min(minY, point.y);
            maxX = std::max(maxX, point.x);
            maxY = std::max(maxY, point.y);
        }
        // inflate the grid
        --minX;
        --minY;
        ++maxX;
        ++maxY;

        // Print the Grid
        Point carrierPos = grid.m_carrier.Position();
        Point leftOfCarrier = carrierPos;
        leftOfCarrier.x--;
        Point pt;
        for (int64_t y = minY; y <= maxY; ++y) {
            pt.y = y;
            for (int64_t x = minX; x <= maxX; ++x) {
                pt.x = x;
                out << (grid.IsInfected(pt) ? grid.InfectedNode : grid.CleanNode);
                if (pt == carrierPos) {
                    out << ']';
                } else if (pt == leftOfCarrier) {
                    out << '[';
                } else {
                    out << ' ';
                }
            }
            out << std::endl;
        }
    }
    return out;
}

} // namespace Day22

constexpr bool g_PrintGrid = false;

void Day22Tests()
{
    const std::vector<std::string> input = { "..#", "#..", "..." };
    Day22::VirusGrid grid(input);
    if (g_PrintGrid) {
        std::cout << grid << std::endl;
        grid.AdvanceToTurn(1);
        std::cout << grid << std::endl;
        grid.AdvanceToTurn(2);
        std::cout << grid << std::endl;
        grid.AdvanceToTurn(6);
        std::cout << grid << std::endl;
        grid.AdvanceToTurn(7);
        std::cout << grid << std::endl;
    }
    grid.AdvanceToTurn(70);
    if (g_PrintGrid) {
        std::cout << grid << std::endl;
    }
    const auto infections70 = grid.TurnsThatInfectedNode();
    const uint64_t expected70 = 41;
    if (infections70 != expected70) std::cerr << "Test 22A1 Error: Got " << infections70 << ", Expected " << expected70 << std::endl;
    grid.AdvanceToTurn(10000);
    const auto infections10000 = grid.TurnsThatInfectedNode();
    const uint64_t expected10000 = 5587;
    if (infections10000 != expected10000) std::cerr << "Test 22A2 Error: Got " << infections10000 << ", Expected " << expected10000 << std::endl;
}

void Day22Problems()
{
    std::cout << "Day 22:\n";
    Day22Tests();
    const auto start = std::chrono::steady_clock::now();
    const auto input = Helpers::ReadFileLines("input_day22.txt");
    Day22::VirusGrid grid(input);
    grid.AdvanceToTurn(10000);
    const auto infections10000 = grid.TurnsThatInfectedNode();
    const auto end = std::chrono::steady_clock::now();
    std::cout << infections10000 << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl << std::endl;
}