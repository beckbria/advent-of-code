#include "2017/lib/Helpers.h"
/*
--- Day 24: Electromagnetic Moat ---

The CPU itself is a large, black building surrounded by a bottomless pit. Enormous metal tubes
extend outward from the side of the building at regular intervals and descend down into the 
void. There's no way to cross, but you need to get inside.

No way, of course, other than building a bridge out of the magnetic components strewn about nearby.

Each component has two ports, one on each end. The ports come in all different types, and 
only matching types can be connected. You take an inventory of the components by their port 
types (your puzzle input). Each port is identified by the number of pins it uses; more pins 
mean a stronger connection for your bridge. A 3/7 component, for example, has a type-3 port 
on one side, and a type-7 port on the other.

Your side of the pit is metallic; a perfect surface to connect a magnetic, zero-pin port. 
Because of this, the first port you use must be of type 0. It doesn't matter what type of 
port you end with; your goal is just to make the bridge as strong as possible.

The strength of a bridge is the sum of the port types in each component. For example, if 
your bridge is made of components 0/3, 3/7, and 7/4, your bridge has a strength of 
0+3 + 3+7 + 7+4 = 24.

For example, suppose you had the following components:

0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10

With them, you could make the following valid bridges:

    0/1
    0/1--10/1
    0/1--10/1--9/10
    0/2
    0/2--2/3
    0/2--2/3--3/4
    0/2--2/3--3/5
    0/2--2/2
    0/2--2/2--2/3
    0/2--2/2--2/3--3/4
    0/2--2/2--2/3--3/5

(Note how, as shown by 10/1, order of ports within a component doesn't matter. 
However, you may only use each port on a component once.)

Of these bridges, the strongest one is 0/1--10/1--9/10; it has a strength of 
0+1 + 1+10 + 10+9 = 31.

What is the strength of the strongest bridge you can make with the components you have available?

--- Part Two ---

The bridge you've built isn't long enough; you can't jump the rest of the way.

In the example above, there are two longest bridges:

0/2--2/2--2/3--3/4
0/2--2/2--2/3--3/5

Of them, the one which uses the 3/5 component is stronger; its strength is 0+2 + 2+2 + 2+3 + 3+5 = 19.

What is the strength of the longest bridge you can make? If you can make multiple bridges of the 
longest length, pick the strongest one.
*/
namespace Day24
{

    struct Component
    {
        Component(unsigned int A = 0, unsigned int B = 0) : portA(A), portB(B) {}
        unsigned int portA;
        unsigned int portB;
        bool inUse = false;
        unsigned int Value() const { return portA + portB; }
        unsigned int OtherSide(const unsigned int port);
    };

    unsigned int Component::OtherSide(const unsigned int port)
    {
        if (port == portA)
        {
            return portB;
        }
        else if (port == portB)
        {
            return portA;
        }
        else
        {
            throw std::invalid_argument("Must specify one of the ports for this component");
        }
    }

    std::vector<Component> ReadComponents(const std::vector<std::string> &input)
    {
        std::vector<Component> components;
        for (const auto &line : input)
        {
            auto sides = Helpers::Tokenize(line, '/');
            components.emplace_back(std::stoi(sides[0]), std::stoi(sides[1]));
        }
        return components;
    }

    template <typename T>
    unsigned int BridgeValue(const T &bridge)
    {
        unsigned int value = 0;
        std::for_each(bridge.cbegin(), bridge.cend(), [&value](const Component &c)
                      { value += c.Value(); });
        return value;
    }

    typedef std::map<unsigned int, std::list<unsigned int>> PortLookup;

    // Build a list of indexes that have the value in question.  This maps
    // port type to a list of indexes containing that port.  If a component has
    // the same port type, it appears only once in the list
    PortLookup BuildPortLookup(const std::vector<Component> &components)
    {
        PortLookup portLookup;
        for (unsigned int i = 0; i < components.size(); ++i)
        {
            portLookup[components[i].portA].push_back(i);
            if (components[i].portA != components[i].portB)
            {
                portLookup[components[i].portB].push_back(i);
            }
        }
        return portLookup;
    }

    // Recursively search for the strongest bridge
    template <typename T>
    void FindBestBridge(
        std::list<Component> &bridge,       // In/Out - the best bridge gets stored here
        std::vector<Component> &components, // The list of all available components
        PortLookup &portLookup,
        unsigned int startFrom, // The port number we're trying to initially match
        T &comparator)
    {
        std::list<Component> bestBridge;
        for (auto candidate : portLookup[startFrom])
        {
            if (!components[candidate].inUse)
            {
                std::list<Component> candidateBridge;
                candidateBridge.push_back(components[candidate]);
                components[candidate].inUse = true;

                FindBestBridge(candidateBridge, components, portLookup, components[candidate].OtherSide(startFrom), comparator);
                if (comparator(bestBridge, candidateBridge))
                {
                    bestBridge = candidateBridge;
                }
                components[candidate].inUse = false;
            }
        }

        // Add the best values to our bridge
        bridge.insert(bridge.end(), bestBridge.begin(), bestBridge.end());
    }

    template <typename T>
    std::list<Component> BestBridge(std::vector<Component> &components, T &&comparator)
    {
        auto portLookup = BuildPortLookup(components);
        std::list<Component> bridge;
        FindBestBridge(bridge, components, portLookup, 0, comparator);
        return bridge;
    }

    unsigned int StrongestBridgeValue(const std::vector<std::string> &input)
    {
        auto components = ReadComponents(input);
        auto bridge = BestBridge(components,
                                 [](const std::list<Component> &currentBest, const std::list<Component> &candidate) -> bool
                                 {
                                     return BridgeValue(candidate) > BridgeValue(currentBest);
                                 });
        return BridgeValue(bridge);
    }

    unsigned int LongestBridgeValue(const std::vector<std::string> &input)
    {
        auto components = ReadComponents(input);
        auto bridge = BestBridge(components,
                                 [](const std::list<Component> &currentBest, const std::list<Component> &candidate) -> bool
                                 {
                                     return ((candidate.size() > currentBest.size()) ||
                                             ((candidate.size() == currentBest.size()) && (BridgeValue(candidate) > BridgeValue(currentBest))));
                                 });
        return BridgeValue(bridge);
    }

} // namespace Day24

void Day24Tests()
{
    const std::vector<std::string> input = {"0/2", "2/2", "2/3", "3/4", "3/5", "0/1", "10/1", "9/10"};
    const auto bridgeStrength = Day24::StrongestBridgeValue(input);
    const unsigned int expectedStrength = 31;
    if (bridgeStrength != expectedStrength)
        std::cerr << "Test 24A Error: Got " << bridgeStrength << ", expected " << expectedStrength << std::endl;
    const auto longestStrength = Day24::LongestBridgeValue(input);
    const unsigned int expectedLongest = 19;
    if (longestStrength != expectedLongest)
        std::cerr << "Test 24B Error: Got " << longestStrength << ", expected " << expectedLongest << std::endl;
}

void Day24Problems()
{
    std::cout << "Day 24:\n";
    Day24Tests();
    const auto start = std::chrono::steady_clock::now();
    const auto input = Helpers::ReadFileLines("2017/24/input_day24.txt");
    const auto bridgeStrength = Day24::StrongestBridgeValue(input);
    const auto longestStrength = Day24::LongestBridgeValue(input);
    const auto end = std::chrono::steady_clock::now();
    std::cout << bridgeStrength << std::endl
              << longestStrength << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl
              << std::endl;
}

int main()
{
    Day24Problems();
    return 0;
}