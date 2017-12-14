#include "Problems.h"
/*
--- Day 14: Disk Defragmentation ---

Suddenly, a scheduled job activates the system's disk defragmenter. Were the situation different, you might sit and watch it 
for a while, but today, you just don't have that kind of time. It's soaking up valuable system resources that are needed 
elsewhere, and so the only option is to help it finish its task as soon as possible.

The disk in question consists of a 128x128 grid; each square of the grid is either free or used. On this disk, the state of 
the grid is tracked by the bits in a sequence of knot hashes.

A total of 128 knot hashes are calculated, each corresponding to a single row in the grid; each hash contains 128 bits 
which correspond to individual grid squares. Each bit of a hash indicates whether that square is free (0) or used (1).

The hash inputs are a key string (your puzzle input), a dash, and a number from 0 to 127 corresponding to the row. 
For example, if your key string were flqrgnkx, then the first row would be given by the bits of the knot hash of 
flqrgnkx-0, the second row from the bits of the knot hash of flqrgnkx-1, and so on until the last row, flqrgnkx-127.

The output of a knot hash is traditionally represented by 32 hexadecimal digits; each of these digits correspond to 
4 bits, for a total of 4 * 32 = 128 bits. To convert to bits, turn each hexadecimal digit to its equivalent binary 
value, high-bit first: 0 becomes 0000, 1 becomes 0001, e becomes 1110, f becomes 1111, and so on; a hash that begins 
with a0c2017... in hexadecimal would begin with 10100000110000100000000101110000... in binary.

Continuing this process, the first 8 rows and columns for key flqrgnkx appear as follows, using # to denote used 
squares, and . to denote free ones:

##.#.#..-->
.#.#.#.#   
....#.#.   
#.#.##.#   
.##.#...   
##..#..#   
.#...#..   
##.#.##.-->
|      |   
V      V   

In this example, 8108 squares are used across the entire 128x128 grid.

Given your actual key string, how many squares are used?

--- Part Two ---

Now, all the defragmenter needs to know is the number of regions. A region is a group of used squares that are all adjacent, not including diagonals. Every used square is in exactly one region: lone used squares form their own isolated regions, while several adjacent squares all count as a single region.

In the example above, the following nine regions are visible, each marked with a distinct digit:

11.2.3..-->
.1.2.3.4
....5.6.
7.8.55.9
.88.5...
88..5..8
.8...8..
88.8.88.-->
|      |
V      V

Of particular interest is the region marked 8; while it does not appear contiguous in this small view, all of the squares marked 8 are connected when considering the whole 128x128 grid. In total, in this example, 1242 regions are present.

How many regions are present given your key string?

*/
namespace Day14 {

class Grid
{
public:
    Grid(const std::string& seed);
    int UsedSquares();
    int Regions();
    int Rows() const;
    int Columns() const;

private:
    // Grid storage is a vector of vector of bytes.  Its dimensions should be 128x16.
    typedef std::vector<std::vector<unsigned int>> GridStorage;
    GridStorage m_grid;
    static constexpr int BitsPerByte = 8;

    struct Point {
        Point(int X, int Y) : x(X), y(Y) {}
        int x;
        int y;
    };

    // This checks the current value of a bit in the 128x128 array
    inline bool IsCellUsed(const GridStorage& grid, const Point& pt) {
        // xCell is the index into the x array that contains the bit we're interested in.
        const unsigned int xCell = pt.x / BitsPerByte;
        const unsigned int bit = pt.x % BitsPerByte;
        // If we want the 0th bit, we want the most significant bit.  Use masks to extract the desired result
        static constexpr unsigned int bitMask[BitsPerByte] = { 0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01 };
        return (grid[pt.y][xCell] & bitMask[bit]);
    }

    inline void ClearCell(GridStorage& grid, const Point& pt) {
        // xCell is the index into the x array that contains the bit we're interested in.
        const unsigned int xCell = pt.x / BitsPerByte;
        const unsigned int bit = pt.x % BitsPerByte;

        // If we want the 0th bit, we want to clear the most significant bit.  Use masks to extract the desired result
        static constexpr unsigned int bitMask[BitsPerByte] = { 
            0xff - 0x80, 
            0xff - 0x40, 
            0xff - 0x20, 
            0xff - 0x10, 
            0xff - 0x08, 
            0xff - 0x04, 
            0xff - 0x02, 
            0xff - 0x01 };

        grid[pt.y][xCell] &= bitMask[bit];
    }

    // Clears all of the bits in a region (a region is a set of bits connected horizontally or vertically)
    void ClearRegion(GridStorage& grid, const Point& pt);
};

Grid::Grid(const std::string& seed)
{
    for (int i = 0; i < 128; ++i) {
        std::stringstream rowSeed;
        rowSeed << seed << '-' << i;
        m_grid.emplace_back(std::move(Day10::KnotHash(rowSeed.str())));
    }
}

int Grid::Rows() const
{
    return m_grid.size();
}

int Grid::Columns() const
{
    if (m_grid.size() < 1) {
        return 0;
    } else {
        return m_grid[0].size() * BitsPerByte;
    }
}

int Grid::UsedSquares()
{
    int used = 0;
    for (const auto &row : m_grid) {
        for (const auto &cell : row) {
            used += Helpers::CountBits(cell);
        }
    }
    return used;
}

int Grid::Regions()
{
    // Create a copy of the grid so that we can mark where we've been
    auto grid = m_grid;
    int regions = 0;

    const auto rows = Rows();
    const auto cols = Columns();
    for (int x = 0; x < cols; ++x) {
        for (int y = 0; y < rows; ++y) {
            Point here(x, y);
            if (IsCellUsed(grid, here)) {
                ++regions;
                ClearRegion(grid, here);
            }
        }
    }

    return regions;
}

void Grid::ClearRegion(GridStorage& grid, const Point& pt)
{
    // Depth first preorder traversal to clear any bits in the region
    std::stack<Point> region;
    region.push(pt);
    const auto lastRow = Rows() - 1;
    const auto lastColumn = Columns() - 1;
    while (!region.empty()) {
        auto current = region.top();
        region.pop();

        if (IsCellUsed(grid, current))
        {
            ClearCell(grid, current);

            if (current.x > 0) {
                Point left(current.x - 1, current.y);
                if (IsCellUsed(grid, left)) region.push(left);
            }

            if (current.x < lastColumn) {
                Point right(current.x + 1, current.y);
                if (IsCellUsed(grid, right)) region.push(right);
            }

            if (current.y > 0) {
                Point up(current.x, current.y - 1);
                if (IsCellUsed(grid, up)) region.push(up);
            }

            if (current.y < lastRow) {
                Point down(current.x, current.y + 1);
                if (IsCellUsed(grid, down)) region.push(down);
            }
        }
    }
}
} // namespace Day14

void Day14Tests()
{
    const std::string input = "flqrgnkx";
    Day14::Grid grid(input);
    const auto used = grid.UsedSquares();
    if (used != 8108) std::cerr << "Test 14A Error: Got " << used << ", Expected 8108\n";
    const auto regions = grid.Regions();
    if (regions != 1242) std::cerr << "Test 14B Error: Got " << regions << ", Expected 1242\n";
}

void Day14Problems()
{
    std::cout << "Day 14:\n";
    Day14Tests();
    const std::string input = "uugsqrei";
    Day14::Grid grid(input);
    std::cout << grid.UsedSquares() << std::endl;
    std::cout << grid.Regions() << std::endl << std::endl;
}