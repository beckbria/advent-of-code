#include "2017/lib/Helpers.h"
/*
--- Day 1: Inverse Captcha ---

The night before Christmas, one of Santa's Elves calls you in a panic. "The printer's broken! We can't print the 
Naughty or Nice List!" By the time you make it to sub-basement 17, there are only a few minutes until midnight. 
"We have a big problem," she says; "there must be almost fifty bugs in this system, but nothing else can print 
The List. Stand in this square, quick! There's no time to explain; if you can convince them to pay you in stars, 
you'll be able to--" She pulls a lever and the world goes blurry.

When your eyes can focus again, everything seems a lot more pixelated than before. She must have sent you inside 
the computer! You check the system clock: 25 milliseconds until midnight. With that much time, you should be able 
to collect all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day millisecond in the advent 
calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You're standing in a room with "digitization quarantine" written in LEDs along one wall. The only door is locked, 
but it includes a small interface. "Restricted Area - Strictly No Digitized Users Allowed."

It goes on to explain that you may only leave by solving a captcha to prove you're not a human. Apparently, 
you only get one millisecond to solve the captcha: too fast for a normal human, but it feels like hours to you.

The captcha requires you to review a sequence of digits (your puzzle input) and find the sum of all digits 
that match the next digit in the list. The list is circular, so the digit after the last digit is the first 
digit in the list.

For example:

1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit.
1111 produces 4 because each digit (all 1) matches the next.
1234 produces 0 because no digit matches the next.
91212129 produces 9 because the only digit that matches the next one is the last digit, 9.

What is the solution to your captcha?

--- Part Two ---

You notice a progress bar that jumps to 50% completion. Apparently, the door isn't yet satisfied, but it did emit a 
star as encouragement. The instructions change:

Now, instead of considering the next digit, it wants you to consider the digit halfway around the circular list. 
That is, if your list contains 10 items, only include a digit in your sum if the digit 10/2 = 5 steps forward 
matches it. Fortunately, your list has an even number of elements.

For example:

1212 produces 6: the list contains 4 items, and all four digits match the digit 2 items ahead.
1221 produces 0, because every comparison is between a 1 and a 2.
123425 produces 4, because both 2s match each other, but no other digit has a match.
123123 produces 12.
12131415 produces 4.

What is the solution to your new captcha?
*/
namespace Day1
{
    int sumRepeatingDigits(const std::string &s)
    {
        if (s.size() < 1)
            return 0;

        char previous = s[s.size() - 1];
        int sum = 0;
        for (char c : s)
        {
            if (c == previous)
            {
                sum += (c - '0');
            }
            previous = c;
        }
        return sum;
    }

    int sumHalfwayAroundDigits(const std::string &s)
    {
        if (s.size() < 1)
            return 0;

        int sum = 0;
        int offset = s.size() / 2;
        for (size_t i = 0; i < s.size(); ++i)
        {
            int oppositeIndex = (i + offset) % s.size();
            if (s[i] == s[oppositeIndex])
            {
                sum += (s[i] - '0');
            }
        }
        return sum;
    }
} // namespace Day1

void Day1ATests()
{
    const struct
    {
        std::string input;
        int answer;
    } testCases[] = {
        {"1122", 3},
        {"12345", 0},
        {"1", 1},
        {"", 0},
        {"11111", 5},
        {"91212129", 9}};
    for (auto &t : testCases)
    {
        int const count = Day1::sumRepeatingDigits(t.input);
        if (count != t.answer)
        {
            std::cerr << "TEST 1A FAILED: " << t.input << " => " << count << " (expected " << t.answer << ")" << std::endl;
        }
    }
}

void Day1BTests()
{
    const struct
    {
        std::string input;
        int answer;
    } testCases[] = {
        {"1212", 6},
        {"1221", 0},
        {"123425", 4},
        {"123123", 12},
        {"12131415", 4}};
    for (auto &t : testCases)
    {
        int const count = Day1::sumHalfwayAroundDigits(t.input);
        if (count != t.answer)
        {
            std::cerr << "TEST 1B FAILED: " << t.input << " => " << count << " (expected " << t.answer << ")" << std::endl;
        }
    }
}

void Day1Problems()
{
    std::cout << "Day 1:\n";
    Day1ATests();
    Day1BTests();
    const auto start = std::chrono::steady_clock::now();
    std::string source = "951484596541141557316984781494999179679767747627132447513171626424561779662873157761442952212296685573452311263445163233493199211387838461594635666699422982947782623317333683978438123261326863959719777179228599319321138948466562743761584836184512984131635354116264899181952748224523953976485816295227945792555726121913344959454458829485471174415775278865324142733339789878929596275998341778873889585819916457474773252249179366599951454182657225576277834669222982366884688565754691273745959468648957498511326215934353963981471593984617554514519623785326888374742147318993423214834751785956958395133486656388454552769722562524415715913869946325551396638593398729938526424994348267935153555851552287223313383583669912941364344694725478258297498969517632881187394141593479818536194597976519254215932257653777455227477617957833273463216593642394215275314734914719726618923177918342664351954252667253233858814365351722938716621544226598956257753212248859258351363174782742336961425325381561575992352415514168782816173861148859478285339529151631429536819286498721812323861771638574344416879476255929929157912984151742613268754779685396125954595318134933366626594498249956388771723777242772654678448815844555372892574747735672368299826548254744359377667294764559334659523233146587568261116253155189394188696831691284711264872914348961888253386971994431352474717376878745948769171243242621219912378731755544387249443997382399714738351857752329367997665166956467544459817582915478514486541453932175598413554259672117364863112592515988922747164842668361925135551248923449968328385889877512156952725198691746951431443497496455761516486573476185321748523644283494181119399874324683922393547682851931435931276267766772798261563117954648576421741384823494187895272582575669685279986988357796138794326125852772995446355723211161523161886222562853546488411563473998633847953246787557146187696947831335722888918172961256498971868946237299523474841983527391489962357196433927251798764362493965894995592683296651874787384247326643886774966828657393717626591578321174832222434128817871765347278152799425565633521152643686221411129463425496425385516719682884157452772141585743166647191938727971366274357874252166721759";
    const auto repeating = Day1::sumRepeatingDigits(source);
    const auto halfway = Day1::sumHalfwayAroundDigits(source);
    const auto end = std::chrono::steady_clock::now();
    std::cout << repeating << std::endl;
    std::cout << halfway << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl
              << std::endl;
}

int main()
{
    Day1Problems();
    return 0;
}