#include "Problems.h"

/*
Wandering further through the circuits of the computer, you come upon a tower of programs 
that have gotten themselves into a bit of trouble. A recursive algorithm has gotten out 
of hand, and now they're balanced precariously in a large tower.

One program at the bottom supports the entire tower. It's holding a large disc, and on the 
disc are balanced several more sub-towers. At the bottom of these sub-towers, standing on 
the bottom disc, are other programs, each holding their own disc, and so on. At the very 
tops of these sub-sub-sub-...-towers, many programs stand simply keeping the disc below 
them balanced but with no disc of their own.

You offer to help, but first you need to understand the structure of these towers. You ask 
each program to yell out their name, their weight, and (if they're holding a disc) the names 
of the programs immediately above them balancing on that disc. You write this information 
down (your puzzle input). Unfortunately, in their panic, they don't do this in an orderly 
fashion; by the time you're done, you're not sure which program gave which information.

For example, if your list is the following :

pbga(66)
xhth(57)
ebii(61)
havc(66)
ktlj(57)
fwft(72)->ktlj, cntj, xhth
qoyq(66)
padx(45)->pbga, havc, qoyq
tknk(41)->ugml, padx, fwft
jptl(61)
ugml(68)->gyxo, ebii, jptl
gyxo(61)
cntj(57)

...then you would be able to recreate the structure of the towers that looks like this:

               gyxo
             /
        ugml - ebii
       /     \
      |        jptl
      |
      |         pbga
     /        /
tknk --- padx - havc
     \        \
      |         goyq
      |
      |         ktlj
       \      /
         fwft - cntj
              \
                xhth

In this example, tknk is at the bottom of the tower(the bottom program), and is holding up ugml, padx, and fwft.
Those programs are, in turn, holding up other programs; in this example, none of those programs are holding up 
any other programs, and are all the tops of their own towers. (The actual tower balancing in front of you is much larger.)

Before you're ready to help them, you need to make sure your information is correct. What is the name of the bottom program?

**********************

--- Part Two ---

The programs explain the situation: they can't get down. Rather, they could get down, if they 
weren't expending all of their energy trying to keep the tower balanced. Apparently, one 
program has the wrong weight, and until it's fixed, they're stuck here.

For any program holding a disc, each program standing on that disc forms a sub-tower. Each of 
those sub-towers are supposed to be the same weight, or the disc itself isn't balanced. The 
weight of a tower is the sum of the weights of the programs in that tower.

In the example above, this means that for ugml's disc to be balanced, 
gyxo, ebii, and jptl must all have the same weight, and they do: 61.

However, for tknk to be balanced, each of the programs standing on its disc and all programs 
above it must each match. This means that the following sums must all be the same:

ugml + (gyxo + ebii + jptl) = 68 + (61 + 61 + 61) = 251
padx + (pbga + havc + qoyq) = 45 + (66 + 66 + 66) = 243
fwft + (ktlj + cntj + xhth) = 72 + (57 + 57 + 57) = 243

As you can see, tknk's disc is unbalanced: ugml's stack is heavier than the other two. 
Even though the nodes above ugml are balanced, ugml itself is too heavy: it needs to 
be 8 units lighter for its stack to weigh 243 and keep the towers balanced. 
If this change were made, its weight would be 60.

Given that exactly one program is the wrong weight, what would its weight need to be to balance the entire tower?
*/

// Topological Sort!

struct Node
{
    // Hold a weak ref to the parent node to avoid a circular dependency
    std::weak_ptr<Node> parent;
    std::string name;

    // A true implementation should encapsulate children so that we could call InvalidateTowerWeight every time a child is added
    // I'm skipping that here due to time constraints since this is a single use case where we build only before we calculate and
    // never add/remove children once we've successfully built the tree
    std::vector<std::shared_ptr<Node>> children;

    int GetWeight() { return m_weight; }
    void SetWeight(int weight) {
        m_weight = weight;
        InvalidateTowerWeight();
    }

    int TowerWeight()
    {
        // Cache this calculation so we don't have to do it repeatedly
        if (m_towerWeight < 0)
        {
            int tower = m_weight;
            m_balancedChildren = true;
            int const childWeight = (children.size() > 0) ? children[0]->TowerWeight() : 0;
            for (auto &child : children) {
                tower += child->TowerWeight();
                if ((child->TowerWeight() != childWeight) || (!child->BalancedChildren())) m_balancedChildren = false;
            }
            m_towerWeight = tower;
        }
        return m_towerWeight;
    }

    bool BalancedChildren()
    {
        int unused = TowerWeight();     // Calculate if we have balanced children
        return m_balancedChildren;
    }

private:
    bool m_balancedChildren = false;
    int m_weight = 0;
    int m_towerWeight = -1;
    void InvalidateTowerWeight()
    {
        m_towerWeight = -1;
        if (auto p = parent.lock()) {
            p->InvalidateTowerWeight();
        }
    }
};

std::shared_ptr<Node> BuildTree(std::vector<std::string>& input)
{
    std::map<std::string, std::shared_ptr<Node>> entry;

    for (auto &line : input) {
        // Assume well formed input.  If this was a real program, write a better parser.
        // If you have to write a better parser, use a better language

        auto tokens = Tokenize(line);

        // The first token is the name
        auto parent = entry[tokens[0]];
        if (parent == nullptr) {
            parent = std::make_shared<Node>();
            entry[tokens[0]] = parent;
            parent->name = std::move(tokens[0]);
        }

        // The second is the weight in parenthesis
        parent->SetWeight(atoi(tokens[1].substr(1, tokens[1].size() - 2).c_str()));

        if (tokens.size() > 2) {
            if (tokens[2] != "->") {
                std::cerr << "Bad Children Token - Got " << tokens[2] << std::endl;
            }

            for (int i = 3; i < tokens.size(); ++i) {
                std::string name = std::move(tokens[i]);
                if (name[name.size() - 1] == ',') {
                    name.resize(name.size() - 1);
                }

                auto child = entry[name];
                if (child == nullptr) {
                    child = std::make_shared<Node>();
                    entry[name] = child;
                    child->name = std::move(name);
                }
                child->parent = parent;

                parent->children.emplace_back(std::move(child));
            }
        }
    }

    // Since we're told that everything is connected, we should be able to take any node and walk its
    // ancestry to get the root
    auto root = entry.begin()->second;
    while (auto p = root->parent.lock()) root = p;

    return root;
}

int BalanceTree(std::shared_ptr<Node> root, int desiredShift)
{
    int newWeight = -1;

    int childWeight = root->children.size() > 0 ? root->children[0]->GetWeight() : 0;
    std::map<int, IntDefaultToZero> weights;
    std::shared_ptr<Node> childWithUnequalWeight;    // As compared to child 0

    for (auto& child : root->children) {
        weights[child->TowerWeight()].val++;

        if (!child->BalancedChildren()) {
            // The problem is at a deeper level.

            // If we already know what change in weight we're looking for, just recurse directly
            if (desiredShift != 0) return BalanceTree(child, desiredShift);

            // Otherwise, what change in weight are we looking for?
            int shift = 0;

            for (auto &c : root->children) {
                shift = c->TowerWeight() - child->TowerWeight();
                if (shift != 0) break;
            }

            return BalanceTree(child, shift);
        }
        
        if (child->GetWeight() != childWeight) {
            childWithUnequalWeight = child;
        }
    }

    if (childWithUnequalWeight) {
        // Our children have unequal weights, and yet none of them reported having unbalanced children.  Thus, one of them is the problem
        if (root->children.size() == 0) {
            return root->GetWeight() + desiredShift;
        }
        else if (root->children.size() == 1) {
            return root->children[0]->GetWeight() + desiredShift;
        }
        else {
            auto child0Weight = root->children[0]->TowerWeight();

            if (desiredShift != 0) {
                for (auto& child : root->children) {
                    if (child->TowerWeight() + desiredShift == child0Weight) return child->GetWeight() + desiredShift;
                    if (child->TowerWeight() == child0Weight + desiredShift) return root->children[0]->GetWeight() + desiredShift;
                }
            }
            else {
                // We haven't had an opportunity to compute the desired shift.  Pick the majority
                int modeCount = INT_MIN;
                int mode = 0;
                for (auto keyPair : weights) {
                    if (keyPair.second.val > modeCount) {
                        modeCount = keyPair.second.val;
                        mode = keyPair.first;
                    }
                }
                
                // Now find the element whose value doesn't correspond to the mode
                for (auto& child : root->children) {
                    if (child->TowerWeight() != mode) {
                        return child->GetWeight() + (mode - child->TowerWeight());
                    }
                }
            }
        }
    }

    return newWeight;
}

void PrintNode(std::shared_ptr<Node> node, std::ostream& out, bool children = true) {
    out << node->name << "(" << node->GetWeight() << ", " << node->TowerWeight() << ") ";
    if (children) {
        for (auto &child : node->children) PrintNode(child, out, false);
    }
}

void PrintTree(std::shared_ptr<Node> root, std::ostream& out) {
    std::queue<std::shared_ptr<Node>> toPrint;
    toPrint.push(root);
    while (!toPrint.empty()) {
        // Breadth First Printing
        auto node = toPrint.front();
        toPrint.pop();
        out << "===";
        PrintNode(node, out, true);
        out << std::endl << std::endl;
        for (auto &child : node->children) toPrint.push(child);
    }
}

void Day7Tests()
{
    std::vector<std::string> input = {
    "pbga (66)", "xhth (57)", "ebii (61)", "havc (66)", "ktlj (57)", "fwft (72) -> ktlj, cntj, xhth",
    "qoyq (66)", "padx (45) -> pbga, havc, qoyq", "tknk (41) -> ugml, padx, fwft", "jptl (61)",
    "ugml (68) -> gyxo, ebii, jptl", "gyxo (61)", "cntj (57)" };

    auto root = BuildTree(input);
    if (root->name != "tknk") std::cout << "Test 7A Failure: Got " << root->name << ", Expected tknk" << std::endl;
    auto newWeight = BalanceTree(root, 0);
    if (newWeight != 60) std::cout << "Test 7B Failure: Got " << newWeight << ", Expected 60" << std::endl;
}

void Day7()
{
    Day7Tests();
    auto input = ReadFileLines("input_day7.txt");
    auto root = BuildTree(input);
    std::cout << "Day 7:\n";
    std::cout << root->name << std::endl;
    std::cout << BalanceTree(root, 0) << std::endl << std::endl;
}