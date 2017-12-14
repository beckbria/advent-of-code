#include "Problems.h"
/*
-- - Day 12: Digital Plumber-- -

Walking along the memory banks of the stream, you find a small village that is experiencing a little confusion : some programs can't communicate with each other.

Programs in this village communicate using a fixed system of pipes.Messages are passed between programs using these pipes, but most programs aren't connected to each other directly. Instead, programs pass messages between each other until the message reaches the intended recipient.

For some reason, though, some of these messages aren't ever reaching their intended recipient, and the programs suspect that some pipes are missing. They would like you to investigate.

You walk through the village and record the ID of each program and the IDs with which it can communicate directly(your puzzle input).Each program has one or more programs with which it can communicate, and these pipes are bidirectional; if 8 says it can communicate with 11, then 11 will say it can communicate with 8.

You need to figure out how many programs are in the group that contains program ID 0.

For example, suppose you go door - to - door like a travelling salesman and record the following list :

0 <-> 2
    1 <-> 1
    2 <-> 0, 3, 4
    3 <-> 2, 4
    4 <-> 2, 3, 6
    5 <-> 6
    6 <-> 4, 5

    In this example, the following programs are in the group that contains program ID 0:

Program 0 by definition.
Program 2, directly connected to program 0.
Program 3 via program 2.
Program 4 via program 2.
Program 5 via programs 6, then 4, then 2.
Program 6 via programs 4, then 2.

Therefore, a total of 6 programs are in this group; all but program 1, which has a pipe that connects it to itself.

How many programs are in the group that contains program ID 0 ?

--- Part Two ---

There are more programs than just the ones in the group containing program ID 0. The rest of them have no way of reaching that group, and still might have no way of reaching each other.

A group is a collection of programs that can all communicate via pipes either directly or indirectly. The programs you identified just a moment ago are all part of the same group. Now, they would like you to determine the total number of groups.

In the example above, there were 2 groups: one consisting of programs 0,2,3,4,5,6, and the other consisting solely of program 1.

How many groups are there in total?

*/

namespace Day12
{
struct Node
{
    // Ref count to the nodes are held by the map.  If we hold strong refs to our linked nodes
    // we risk a circular reference
    Node(int i) : value(i) {}
    int value;
    std::vector<std::weak_ptr<Node>> neighbors;
    bool visited = false;
};

std::map<int, std::shared_ptr<Node>> BuildGraph(const std::vector<std::string>& input)
{
    std::map<int, std::shared_ptr<Node>> nodes;
    for (auto &line : input) {
        auto tokens = Helpers::Tokenize(line);
        int originValue = std::stoi(tokens[0]);
        if (nodes.count(originValue) == 0) nodes[originValue] = std::make_shared<Node>(originValue);
        auto origin = nodes[originValue];

        // Skip the <==> token
        for (size_t i = 2; i < tokens.size(); ++i) {
            Helpers::RemoveTrailingCharacter(tokens[i], ',');
            auto neighborValue = std::stoi(tokens[i]);
            if (nodes.count(neighborValue) == 0) nodes[neighborValue] = std::make_shared<Node>(neighborValue);
            origin->neighbors.push_back(nodes[neighborValue]);
        }
    }
    return nodes;
}

int ConnectedNodeCount(std::shared_ptr<Node> root)
{
    // Use Depth-First Search to enumerate the nodes
    std::stack<std::shared_ptr<Node>> remaining;
    remaining.push(root);
    int count = 0;
    while (!remaining.empty()) {
        auto current = remaining.top();
        remaining.pop();
        // If a node is reachable via multiple other nodes before it has been visited, it could appear in the list multiple times
        if (current->visited) continue;
        current->visited = true;
        ++count;

        for (auto node : current->neighbors) {
            auto strong = node.lock();
            if (strong && !strong->visited) {
                remaining.emplace(std::move(strong));
            }
        }
    }

    return count;
}

int TotalGroups(const std::map<int, std::shared_ptr<Node>>& nodes)
{
    int groups = 0;
    int nodeCount = 0;

    // First, mark all nodes not visited
    for (auto &n : nodes) n.second->visited = false;
    for (auto &n : nodes) {
        if (!n.second->visited) {
            // Search out the group and mark it visited
            nodeCount += ConnectedNodeCount(n.second);
            ++groups;
        }
    }

    // We should have visited all nodes
    if (nodeCount != nodes.size()) {
        std::cerr << "TotalGroups Error: Saw " << nodeCount << " nodes, expected " << nodes.size() << std::endl;
    }

    return groups;
}

} // namespace Day12

void Day12Tests()
{
    const std::vector<std::string> input = {
        "0 <-> 2",
        "1 <-> 1",
        "2 <-> 0, 3, 4",
        "3 <-> 2, 4",
        "4 <-> 2, 3, 6",
        "5 <-> 6",
        "6 <-> 4, 5"
    };
    auto nodes = Day12::BuildGraph(input);
    auto count = Day12::ConnectedNodeCount(nodes[0]);
    if (count != 6) std::cerr << "Test 12A Error: Got " << count << ", Expected 6" << std::endl;
    int totalGroups = TotalGroups(nodes);
    if (totalGroups != 2) std::cerr << "Test 12B Error: Got " << totalGroups << ", Expected 2" << std::endl;
}

void Day12Problems()
{
    std::cout << "Day 12:\n";
    Day12Tests();
    auto input = Helpers::ReadFileLines("input_day12.txt");
    auto nodes = Day12::BuildGraph(input);
    std::cout << Day12::ConnectedNodeCount(nodes[0]) << std::endl;
    std::cout << Day12::TotalGroups(nodes) << std::endl << std::endl;
}