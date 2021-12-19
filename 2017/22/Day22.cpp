#include "2017/lib/Helpers.h"
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

--- Part Two ---

As you go to remove the virus from the infected nodes, it evolves to resist your attempt.

Now, before it infects a clean node, it will weaken it to disable your defenses. If it encounters an infected node, 
it will instead flag the node to be cleaned in the future. So:

Clean nodes become weakened.
Weakened nodes become infected.
Infected nodes become flagged.
Flagged nodes become clean.

Every node is always in exactly one of the above states.

The virus carrier still functions in a similar way, but now uses the following logic during its bursts of action:

Decide which way to turn based on the current node:
If it is clean, it turns left.
If it is weakened, it does not turn, and will continue moving in the same direction.
If it is infected, it turns right.
If it is flagged, it reverses direction, and will go back the way it came.
Modify the state of the current node, as described above.
The virus carrier moves forward one node in the direction it is facing.

Start with the same map (still using . for clean and # for infected) and still with the virus carrier starting in the middle and facing up.

Using the same initial state as the previous example, and drawing weakened as W and flagged as F, the middle of the 
infinite grid looks like this, with the virus carrier's position again marked with [ ]:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . # . . .
. . . #[.]. . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .

This is the same as before, since no initial nodes are weakened or flagged. The virus carrier is on a clean node, 
so it still turns left, instead weakens the node, and moves left:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . # . . .
. . .[#]W . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .

The virus carrier is on an infected node, so it still turns right, instead flags the node, and moves up:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . .[.]. # . . .
. . . F W . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .

This process repeats three more times, ending on the previously-flagged node and facing right:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . W W . # . . .
. . W[F]W . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .

Finding a flagged node, it reverses direction and cleans the node:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . W W . # . . .
. .[W]. W . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .

The weakened node becomes infected, and it continues in the same direction:

. . . . . . . . .
. . . . . . . . .
. . . . . . . . .
. . W W . # . . .
.[.]# . W . . . .
. . . . . . . . .
. . . . . . . . .
. . . . . . . . .

Of the first 100 bursts, 26 will result in infection. Unfortunately, another feature of this evolved virus is speed; 
of the first 10000000 bursts, 2511944 will result in infection.

Given your actual map, after 10000000 bursts of activity, how many bursts cause a node to become infected? 
(Do not count nodes that begin infected.)
*/
namespace Day22
{

    struct Point
    {
        Point(int64_t X = 0, int64_t Y = 0) : x(X), y(Y) {}
        int64_t x;
        int64_t y;
    };

    bool operator==(const Point &left, const Point &right);

    class PointHash
    {
    public:
        std::size_t operator()(const Point &p) const
        {
            // This isn't a particularly good hash.  A much better choice would be something like boost::hash_combine.
            // I'm avoiding doing that here because I don't want to pull in a massive set of dependencies and licensing
            // constraints.
            return std::hash<uint64_t>()(p.x) + std::hash<uint64_t>()(p.y);
        }
    };

    class VirusCarrier
    {
    public:
        Point Position() const { return m_position; }
        void RotateLeft() { m_direction = static_cast<Direction>((static_cast<int>(m_direction) + 3) % 4); }
        void RotateRight() { m_direction = static_cast<Direction>((static_cast<int>(m_direction) + 1) % 4); }
        void ReverseDirection() { m_direction = static_cast<Direction>((static_cast<int>(m_direction) + 2) % 4); }
        void MoveForward();

    private:
        enum class Direction
        {
            Up = 0,
            Right = 1,
            Down = 2,
            Left = 3
        };

        Point m_position; // The carrier always starts at (0,0)
        Direction m_direction = Direction::Up;
    };

    class VirusGrid
    {
    public:
        VirusGrid(const std::vector<std::string> &input);
        void AdvanceToTurn(unsigned int turn);
        uint64_t TurnsThatInfectedNode() { return m_infectedTurns; }
        friend std::ostream &operator<<(std::ostream &out, const VirusGrid &grid);

    protected:
        enum class NodeState : char
        {
            Clean = '.',
            Weakened = 'W',
            Infected = '#',
            Flagged = 'F'
        };
        NodeState ClassifyPoint(const Point &pt) const;
        virtual void AdvanceTurn();

        // There's an infinite grid of clean nodes.  We don't need to store them - only keep track of the unclean ones.
        std::unordered_map<Point, NodeState, PointHash> m_unclean;
        VirusCarrier m_carrier;
        uint64_t m_infectedTurns = 0;
        uint64_t m_currentTurn = 0;
    };

    class EvolvedVirusGrid : public VirusGrid
    {
    public:
        EvolvedVirusGrid(const std::vector<std::string> &input) : VirusGrid(input) {}

    protected:
        virtual void AdvanceTurn() override;
    };

    bool operator==(const Point &left, const Point &right)
    {
        return (left.x == right.x) && (left.y == right.y);
    }

    VirusGrid::NodeState VirusGrid::ClassifyPoint(const Point &pt) const
    {
        if (m_unclean.count(pt) > 0)
        {
            return m_unclean.at(pt);
        }
        else
        {
            return NodeState::Clean;
        }
    }

    void VirusCarrier::MoveForward()
    {
        switch (m_direction)
        {
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

    VirusGrid::VirusGrid(const std::vector<std::string> &input)
    {
        if (input.size() > 0)
        {
            Point middle((input[0].size() / 2), (input.size() / 2)); // The +1 to get to the center of odd shapes is implied by 0 indexing
            for (int y = 0; y < (int)input.size(); ++y)
            {
                for (int x = 0; x < (int)input[0].size(); ++x)
                {
                    const NodeState state = static_cast<NodeState>(input[y][x]);
                    if (state != NodeState::Clean)
                    {
                        m_unclean[Point(x - middle.x, y - middle.y)] = state;
                    }
                }
            }
        }
    }

    void VirusGrid::AdvanceToTurn(unsigned int turn)
    {
        while (m_currentTurn < static_cast<uint64_t>(turn))
        {
            AdvanceTurn();
        }
    }

    void VirusGrid::AdvanceTurn()
    {
        const auto currentPosition = m_carrier.Position();
        if (ClassifyPoint(currentPosition) == NodeState::Infected)
        {
            m_unclean.erase(currentPosition); // Clean the current node
            m_carrier.RotateRight();
        }
        else
        {
            m_unclean[currentPosition] = NodeState::Infected;
            ++m_infectedTurns;
            m_carrier.RotateLeft();
        }
        m_carrier.MoveForward();
        ++m_currentTurn;
    }

    void EvolvedVirusGrid::AdvanceTurn()
    {
        const auto currentPosition = m_carrier.Position();
        switch (ClassifyPoint(currentPosition))
        {
        case NodeState::Clean:
            m_unclean[currentPosition] = NodeState::Weakened;
            m_carrier.RotateLeft();
            break;
        case NodeState::Weakened:
            m_unclean[currentPosition] = NodeState::Infected;
            ++m_infectedTurns;
            // Do not turn - keep moving in the same direction
            break;
        case NodeState::Infected:
            m_unclean[currentPosition] = NodeState::Flagged;
            m_carrier.RotateRight();
            break;
        case NodeState::Flagged:
            m_unclean.erase(currentPosition); // Clean the current node
            m_carrier.ReverseDirection();
            break;
        }
        m_carrier.MoveForward();
        ++m_currentTurn;
    }

    std::ostream &operator<<(std::ostream &out, const VirusGrid &grid)
    {
        if (grid.m_unclean.size() > 0)
        {
            // Find the corners
            int64_t minX = INT64_MAX, minY = INT64_MAX, maxX = INT64_MIN, maxY = INT64_MIN;
            for (const auto &point : grid.m_unclean)
            {
                minX = std::min(minX, point.first.x);
                minY = std::min(minY, point.first.y);
                maxX = std::max(maxX, point.first.x);
                maxY = std::max(maxY, point.first.y);
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
            for (int64_t y = minY; y <= maxY; ++y)
            {
                pt.y = y;
                for (int64_t x = minX; x <= maxX; ++x)
                {
                    pt.x = x;
                    out << static_cast<char>(grid.ClassifyPoint(pt));
                    if (pt == carrierPos)
                    {
                        out << ']';
                    }
                    else if (pt == leftOfCarrier)
                    {
                        out << '[';
                    }
                    else
                    {
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
    const std::vector<std::string> input = {"..#", "#..", "..."};
    Day22::VirusGrid grid(input);
    if (g_PrintGrid)
    {
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
    if (g_PrintGrid)
    {
        std::cout << grid << std::endl;
    }
    const auto infections70 = grid.TurnsThatInfectedNode();
    const uint64_t expected70 = 41;
    if (infections70 != expected70)
        std::cerr << "Test 22A1 Error: Got " << infections70 << ", Expected " << expected70 << std::endl;

    grid.AdvanceToTurn(10000);
    const auto infections10000 = grid.TurnsThatInfectedNode();
    const uint64_t expected10000 = 5587;
    if (infections10000 != expected10000)
        std::cerr << "Test 22A2 Error: Got " << infections10000 << ", Expected " << expected10000 << std::endl;

    Day22::EvolvedVirusGrid evolvedGrid(input);
    evolvedGrid.AdvanceToTurn(100);
    const auto infections100 = evolvedGrid.TurnsThatInfectedNode();
    const uint64_t expected100 = 26;
    if (infections100 != expected100)
        std::cerr << "Test 22B1 Error: Got " << infections100 << ", Expected " << expected100 << std::endl;

    evolvedGrid.AdvanceToTurn(10000000);
    const auto infections10000000 = evolvedGrid.TurnsThatInfectedNode();
    const auto expected10000000 = 2511944;
    if (infections10000000 != expected10000000)
        std::cerr << "Test 22B2 Error: Got " << infections10000000 << ", Expected " << expected10000000 << std::endl;
}

void Day22Problems()
{
    std::cout << "Day 22:\n";
    Day22Tests();
    const auto start = std::chrono::steady_clock::now();
    const auto input = Helpers::ReadFileLines("2017/22/input_day22.txt");

    Day22::VirusGrid grid(input);
    grid.AdvanceToTurn(10000);
    const auto infections10000 = grid.TurnsThatInfectedNode();

    Day22::EvolvedVirusGrid evolvedGrid(input);
    evolvedGrid.AdvanceToTurn(10000000);
    const auto infections10000000 = evolvedGrid.TurnsThatInfectedNode();

    const auto end = std::chrono::steady_clock::now();
    std::cout << infections10000 << std::endl
              << infections10000000 << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl
              << std::endl;
}

int main()
{
    Day22Problems();
    return 0;
}