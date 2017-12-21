#include "Problems.h"
/*
--- Day 21: Fractal Art ---

You find a program trying to generate some art. It uses a strange process that involves repeatedly 
enhancing the detail of an image through a set of rules.

The image consists of a two-dimensional square grid of pixels that are either on (#) or off (.). 
The program always begins with this pattern:

.#.
..#
###

Because the pattern is both 3 pixels wide and 3 pixels tall, it is said to have a size of 3.

Then, the program repeats the following process:

If the size is evenly divisible by 2, break the pixels up into 2x2 squares, and convert each 2x2 square 
into a 3x3 square by following the corresponding enhancement rule.
Otherwise, the size is evenly divisible by 3; break the pixels up into 3x3 squares, and convert each 
3x3 square into a 4x4 square by following the corresponding enhancement rule.

Because each square of pixels is replaced by a larger one, the image gains pixels and so its size increases.

The artist's book of enhancement rules is nearby (your puzzle input); however, it seems to be missing rules. 
The artist explains that sometimes, one must rotate or flip the input pattern to find a match. (Never rotate 
or flip the output pattern, though.) Each pattern is written concisely: rows are listed as single units, 
ordered top-down, and separated by slashes. For example, the following rules correspond to the adjacent patterns:

../.#  =  ..
          .#

                .#.
.#./..#/###  =  ..#
                ###

                        #..#
#..#/..../#..#/.##.  =  ....
                        #..#
                        .##.

When searching for a rule to use, rotate and flip the pattern as necessary. For example, 
all of the following patterns match the same rule:

.#.   .#.   #..   ###
..#   #..   #.#   ..#
###   ###   ##.   .#.

Suppose the book contained the following two rules:

../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#

As before, the program begins with this pattern:

.#.
..#
###

The size of the grid (3) is not divisible by 2, but it is divisible by 3. It divides evenly into a 
single square; the square matches the second rule, which produces:

#..#
....
....
#..#

The size of this enhanced grid (4) is evenly divisible by 2, so that rule is used. It divides evenly into four squares:

#.|.#
..|..
--+--
..|..
#.|.#

Each of these squares matches the same rule (../.# => ##./#../...), three of which require some flipping 
and rotation to line up with the rule. The output for the rule is the same in all four cases:

##.|##.
#..|#..
...|...
---+---
##.|##.
#..|#..
...|...

Finally, the squares are joined into a new grid:

##.##.
#..#..
......
##.##.
#..#..
......

Thus, after 2 iterations, the grid contains 12 pixels that are on.

How many pixels stay on after 5 iterations?

--- Part Two ---

How many pixels stay on after 18 iterations?
*/
namespace Day21 {

    struct Rectangle {
        Rectangle(int Top, int Left, int Bottom, int Right) : top(Top), left(Left), bottom(Bottom), right(Right) {}
        int top;
        int left;
        int bottom;
        int right;

        bool Valid() const { return (right >= left) && (bottom >= top); }
    };
    
    class Image {
    public:
        Image(const std::string& pixels);   // Parses the image from its representation in the input file i.e. ".#/##"
        Image(unsigned int height = 0, unsigned int width = 0);
        std::string str() const;  // Converts the image into its string representation i.e. ".#/##"

        void RotateRight();
        void FlipHorizontal();
        void FlipVertical();
        std::vector<std::string> AllPermutations();     // Generates all combinations of rotating and flipping the image

        Image SubImage(const Rectangle& dimensions) const;
        int Height() const { return m_contents.size(); }
        int Width() const { return m_contents.size() > 0 ? m_contents[0].size() : 0; }
        unsigned int SetBits() const;  // Returns the number of pixels turned on
        void Set(int top, int left, const Image& content);  // Paints the content of an image into another

    protected:
        static constexpr char PixelOff = '.';
        static constexpr char PixelOn = '#';

        // Used to generate subimages without having to reparse the string
        Image(const std::vector<std::string>& pixels) : m_contents(pixels) {};
        void AddFlipPermutations(std::vector<std::string>& permutations);
        // For very large images, it would be worthwhile to store the bits using a bit vector instead of using an entire
        // char for each pixel.  That said:
        // 1) The problem set calls for small image sizes (starting with a 3x3 and iterating a few steps)
        // 2) The dimensions grow by either (4/3)^2 or (3/2)^2 each time, which is to say that they roughly double.  
        //    Thus, using 8x the memory to store a char instead of a bit only buys us 3 extra turns.  What does it gain us?
        // 3) Using strings gives us convenient substr functionality to get the subsquares, concatenation, and means that we
        //    don't have to transform the input nearly as much.
        std::vector<std::string> m_contents;
    };

    class Rule {
    public:
        Rule(const std::string& str);     // Parses the format of a rule in the input file
        int InputDimension() const { return m_input.Height(); }
        int OutputDimension() const { return m_output.Height(); }
        const Image& Input() const { return m_input; }
        const Image& Output() const { return m_output; }
    protected:
        Image m_input;
        Image m_output;
    };

    class Rulebook {
    public:
        Rulebook(const std::vector<std::string>& rules);
        Image Match(const std::string& pattern);
    protected:
        std::vector<Rule> m_rules;
        // Store a copy of all the transformed versions of each rule for quick lookup
        std::unordered_map<std::string, int> m_ruleTransforms;
    };

    class Grid {
    public:
        Grid(const Rulebook& book, Image initialContents) : m_rules(book), m_contents(initialContents) {};
        int SetBits() { return m_contents.SetBits(); }
        int Iteration() const { return m_iteration; }
        void Tick(); // Advance to the next iteration
        std::string Contents() const { return m_contents.str(); }
    protected:
        Rulebook m_rules;
        Image m_contents;
        int m_iteration = 0;
    };

    Image::Image(const std::string& pixels)
    {
        m_contents = Helpers::Tokenize(pixels, '/');
    }

    Image::Image(unsigned int height, unsigned int width)
    {
        m_contents.resize(height);
        for (auto &line : m_contents) line.resize(width, PixelOff);
    }

    Image Image::SubImage(const Rectangle& dimensions) const
    {
        std::vector<std::string> pixels;
        if (dimensions.Valid() &&
            (dimensions.top >= 0) && (dimensions.left >= 0) &&
            (dimensions.bottom < (int)m_contents.size()) &&
            (dimensions.right < (int)m_contents[0].size()))
        {
            const auto width = (dimensions.right - dimensions.left) + 1;
            for (int y = dimensions.top; y <= dimensions.bottom; ++y) {
                pixels.emplace_back(m_contents[y].substr(dimensions.left, width));
            }
        }
        else {
            throw std::invalid_argument("Received invalid rectangle");
        }
        return Image(std::move(pixels));
    }

    std::string Image::str() const
    {
        std::stringstream output;
        // This should be:
        // std::copy(m_contents.begin(), m_contents.end(), std::ostream_iterator<std::string>(output, "/"));
        // But that appears to insert a trailing slash
        const int count = m_contents.size() - 1;
        for (int i = 0; i <= count; ++i) {
            output << m_contents[i];
            if (i != count) output << "/";
        }
        return output.str();
    }

    void Image::RotateRight()
    {
        if (m_contents.size() < 1) return;

        // Potential perf optimization: We can rotate squares in place
        // For a non-square rectangle, we'll need to make a copy.
        const int currentColumns = m_contents[0].size();
        const int currentRows = m_contents.size();

        std::vector<std::string> rotated;
        rotated.resize(currentColumns);
        for (auto &s : rotated) s.resize(currentRows, ' ');

        for (int x = 0; x < currentRows; ++x) {
            for (int y = 0; y < currentColumns; ++y) {
                rotated[x][y] = m_contents[currentColumns - (y+1)][x];
            }
        }
        std::swap(m_contents, rotated);
    }

    void Image::FlipHorizontal()
    {
        if (m_contents.size() < 1) return;
        int left = 0;
        int right = m_contents[0].size() - 1;
        while (right > left) {
            for (auto &line : m_contents) std::swap(line[left], line[right]);
            ++left;
            --right;
        }
    }

    void Image::FlipVertical()
    {
        int top = 0;
        int bottom = static_cast<int>(m_contents.size()) - 1;
        while (bottom > top) {
            std::swap(m_contents[top++], m_contents[bottom--]);
        }
    }

    std::vector<std::string> Image::AllPermutations()
    {
        // The problem description is ambiguous as to whether you're allowed to "flip OR rotate" the image "flip AND rotate".
        // They use both expressions.  For now we're doing all combinations of the two.  We're intentionally NOT doing generating
        // the patern where an image is flipped both vertically and horizontally because it's never mentioned.
        // If we find a missing pattern, consider adding that.  If we find a double match, fall back to only adding just rotation
        // or just flipping.
        std::vector<std::string> permutations;
        for (int i = 0; i < 4; ++i) {
            AddFlipPermutations(permutations);
            RotateRight();  // The final rotation returns us to where we began
        }
        return permutations;
    }

    void Image::AddFlipPermutations(std::vector<std::string>& permutations)
    {
        permutations.emplace_back(str());   // Base Case
        FlipHorizontal();
        permutations.emplace_back(str());   // Horizontal Flip
        FlipHorizontal();                   // Undo
        FlipVertical();
        permutations.emplace_back(str());   // Vertical Flip
        FlipVertical();                     // Undo
    }

    unsigned int Image::SetBits() const
    {
        unsigned int setBits = 0;
        for (const auto &line : m_contents) {
            for (const auto c : line) {
                if (c == PixelOn) ++setBits;
            }
        }
        return setBits;
    }

    void Image::Set(int top, int left, const Image& content)
    {
        for (unsigned int y = 0; (y < content.m_contents.size()) && ((top + y) < this->m_contents.size()); ++y) {
            for (unsigned int x = 0; (x < content.m_contents[0].size()) && ((left + x) < this->m_contents[0].size()); ++x) {
                this->m_contents[top + y][left + x] = content.m_contents[y][x];
            }
        }
    }

    Rule::Rule(const std::string& str)
    {
        // Split into input and output around =>
        auto images = Helpers::Tokenize(str);
        m_input = Image(images[0]);
        // images[1] == "=>"
        m_output = Image(images[2]);
    }

    Rulebook::Rulebook(const std::vector<std::string>& rules)
    {
        for (const auto& r : rules) {
            m_rules.emplace_back(r);
            const int ruleIndex = m_rules.size() - 1;
            // AllPermutations returns the object to its original state when it's done, but modifies its contents in the process.
            // As such, it's not const in the compiler definition, but it is const in practice.
            for (auto inputRule : const_cast<Image&>(m_rules[ruleIndex].Input()).AllPermutations()) {
                if ((m_ruleTransforms.count(inputRule) > 0) && (m_ruleTransforms[inputRule] != ruleIndex)) {
                    throw std::invalid_argument("Conflicting Rules");
                }
                m_ruleTransforms[inputRule] = ruleIndex;
            }
        }
    }

    Image Rulebook::Match(const std::string& pattern)
    {
        if (m_ruleTransforms.count(pattern) == 0) {
            throw std::invalid_argument("Unknown Pattern");
        }
        return m_rules[m_ruleTransforms[pattern]].Output().str();
    }

    void Grid::Tick()
    {
        const int inputBlockSize = (m_contents.Height() % 2) ? 3 : 2;
        const int outputBlockSize = (inputBlockSize == 3) ? 4 : 3;
        const int newDimension = m_contents.Height() / inputBlockSize * outputBlockSize;
        Image newContents(newDimension, newDimension);

        int outputTop = 0;
        for (int inputTop = 0; inputTop < m_contents.Height(); inputTop += inputBlockSize) {
            const int inputBottom = inputTop + inputBlockSize - 1;
            int outputLeft = 0;
            for (int left = 0; left < m_contents.Width(); left += inputBlockSize) {
                const int inputRight = left + inputBlockSize - 1;
                Image sector = m_contents.SubImage(Rectangle(inputTop, left, inputBottom, inputRight));
                auto pattern = sector.str();
                auto replacement = m_rules.Match(pattern);
                newContents.Set(outputTop, outputLeft, replacement);
                outputLeft += outputBlockSize;
            }
            outputTop += outputBlockSize;
        }

        std::swap(newContents, m_contents);
        ++m_iteration;
    }

} // namespace Day21

const Day21::Image g_StartingPosition(".#./..#/###");
constexpr bool g_printGrid = false;

void PrintGrid(const Day21::Grid& grid) {
    if (g_printGrid) {
        auto lines = Helpers::Tokenize(grid.Contents(), '/');
        for (auto &l : lines) std::cout << l << std::endl;
        std::cout << std::endl;
    }
}

void Day21Tests()
{ 
    const std::vector <std::string> rules = { "../.# => ##./#../...", ".#./..#/### => #..#/..../..../#..#" };
    Day21::Rulebook rulebook(rules);
    Day21::Grid grid(rulebook, g_StartingPosition);
    for (unsigned int i = 0; i < 2; ++i) {
        PrintGrid(grid);
        grid.Tick();
    }
    PrintGrid(grid);
    const auto pixels = grid.SetBits();
    const int expectedPixels = 12;
    if (pixels != expectedPixels) std::cerr << "Test 21A Error: Got " << pixels << ", Expected " << expectedPixels << std::endl;
}

void Day21Problems()
{
    std::cout << "Day 21:\n";
    Day21Tests();
    const auto start = std::chrono::steady_clock::now();
    const auto input = Helpers::ReadFileLines("input_day21.txt");
    Day21::Rulebook rulebook(input);
    Day21::Grid grid(rulebook, g_StartingPosition);
    for (unsigned int i = 0; i < 5; ++i) {
        PrintGrid(grid);
        grid.Tick();
    }
    PrintGrid(grid);
    const auto pixelsAfter5 = grid.SetBits();
    for (auto i = grid.Iteration(); i < 18; ++i) grid.Tick();
    const auto pixelsAfter18 = grid.SetBits();
    const auto end = std::chrono::steady_clock::now();
    std::cout << pixelsAfter5 << std::endl << pixelsAfter18 << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl << std::endl;
}