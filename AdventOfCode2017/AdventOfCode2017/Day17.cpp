#include "Problems.h"
/*
--- Day 17: Spinlock ---

Suddenly, whirling in the distance, you notice what looks like a massive, pixelated hurricane: 
a deadly spinlock. This spinlock isn't just consuming computing power, but memory, too; vast, 
digital mountains are being ripped from the ground and consumed by the vortex.

If you don't move quickly, fixing that printer will be the least of your problems.

This spinlock's algorithm is simple but efficient, quickly consuming everything in its path. 
It starts with a circular buffer containing only the value 0, which it marks as the current 
position. It then steps forward through the circular buffer some number of steps (your puzzle 
input) before inserting the first new value, 1, after the value it stopped on. The inserted 
value becomes the current position. Then, it steps forward from there the same number of steps, 
and wherever it stops, inserts after it the second new value, 2, and uses that as the new 
current position again.

It repeats this process of stepping forward, inserting a new value, and using the location of 
the inserted value as the new current position a total of 2017 times, inserting 2017 as its 
final operation, and ending with a total of 2018 values (including 0) in the circular buffer.

For example, if the spinlock were to step 3 times per insert, the circular buffer would begin 
to evolve like this (using parentheses to mark the current position after each iteration of 
the algorithm):

    (0), the initial state before any insertions.
    0 (1): the spinlock steps forward three times (0, 0, 0), and then inserts the first value, 1, 
           after it. 1 becomes the current position.
    0 (2) 1: the spinlock steps forward three times (0, 1, 0), and then inserts the second value, 
             2, after it. 2 becomes the current position.
    0  2 (3) 1: the spinlock steps forward three times (1, 0, 2), and then inserts the third value, 
                3, after it. 3 becomes the current position.

And so on:

    0  2 (4) 3  1
    0 (5) 2  4  3  1
    0  5  2  4  3 (6) 1
    0  5 (7) 2  4  3  6  1

    0  5  7  2  4  3 (8) 6  1
    0 (9) 5  7  2  4  3  8  6  1

Eventually, after 2017 insertions, the section of the circular buffer near the last insertion 
looks like this:

1512  1134  151 (2017) 638  1513  851

Perhaps, if you can identify the value that will ultimately be after the last value written (2017), 
you can short-circuit the spinlock. In this example, that would be 638.

What is the value after 2017 in your completed circular buffer?

Your puzzle input is 328.
*/
namespace Day17 {
    class Spinlock : public std::list<unsigned int> {
    public:
        Spinlock(unsigned int advance, unsigned int elementsToInsert);
        const_iterator CurrentPosition() const { return m_currentPosition; }

    private:
        inline void AdvanceCurrentPosition()
        {
            ++m_currentPosition;
            // This is a circular buffer, so loop back if we hit the end
            if (m_currentPosition == end()) m_currentPosition = begin();
        }
        const_iterator m_currentPosition;
        unsigned int m_advance;
    };

    Spinlock::Spinlock(unsigned int advance, unsigned int elementsToInsert)
        : m_advance(advance)
    {
        // Insert the initial value
        m_currentPosition = insert(begin(), 0);
        for (unsigned int i = 1; i <= elementsToInsert; ++i) {
            unsigned int positionsToAdvance = m_advance % size();
            for (unsigned int i = 0; i < positionsToAdvance; ++i) AdvanceCurrentPosition();
            // We want to insert after the selected element
            auto insertPos = m_currentPosition;
            ++insertPos;
            m_currentPosition = insert(insertPos, i);
        }
    }
} // namespace Day17

void Day17Tests()
{
    auto spinlock = Day17::Spinlock(3, 9);
    const std::vector<int> expected = { 0, 9, 5, 7, 2, 4, 3, 8, 6, 1 };
    auto expectedIterator = expected.begin();
    for (auto it = spinlock.begin(); it != spinlock.end(); ++it) {
        if (*it != *expectedIterator) {
            std::cerr << "Test 17A1 Error: Got {";
            for (auto outputIt = spinlock.begin(); outputIt != spinlock.end(); ++outputIt) {
                std::cerr << *outputIt << ",";
            }
            std::cerr << "}, Expected {";
            for (auto outputIt = expected.begin(); outputIt != expected.end(); ++outputIt) {
                std::cerr << *outputIt << ",";
            }
            std::cerr << "}\n";
            break;
        }
        ++expectedIterator;
    }

    auto spinlock2017 = Day17::Spinlock(3, 2017);
    auto pos = ++spinlock2017.CurrentPosition();
    if (*pos != 638) std::cerr << "Test 17A2 Error: Got " << *pos << ", Expected 638\n";    
}

void Day17Problems()
{
    std::cout << "Day 17:\n";
    Day17Tests();
    const auto start = std::chrono::steady_clock::now();
    auto spinlock = Day17::Spinlock(328, 2017);
    auto nextValue = ++spinlock.CurrentPosition();
    const auto end = std::chrono::steady_clock::now();
    std::cout << *nextValue << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl << std::endl;
}