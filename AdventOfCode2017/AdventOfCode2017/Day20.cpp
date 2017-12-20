#include "Problems.h"
/*
--- Day 20: Particle Swarm ---

Suddenly, the GPU contacts you, asking for help. Someone has asked it to simulate too many particles, and it won't be able to finish them all in time to render the next frame at this rate.

It transmits to you a buffer (your puzzle input) listing each particle in order (starting with particle 0, then particle 1, particle 2, and so on). For each particle, it provides the X, Y, and Z coordinates for the particle's position (p), velocity (v), and acceleration (a), each in the format <X,Y,Z>.

Each tick, all particles are updated simultaneously. A particle's properties are updated in the following order:

    Increase the X velocity by the X acceleration.
    Increase the Y velocity by the Y acceleration.
    Increase the Z velocity by the Z acceleration.
    Increase the X position by the X velocity.
    Increase the Y position by the Y velocity.
    Increase the Z position by the Z velocity.

Because of seemingly tenuous rationale involving z-buffering, the GPU would like to know which particle will stay closest to position <0,0,0> in the long term. Measure this using the Manhattan distance, which in this situation is simply the sum of the absolute values of a particle's X, Y, and Z position.

For example, suppose you are only given two particles, both of which stay entirely on the X-axis (for simplicity). Drawing the current states of particles 0 and 1 (in that order) with an adjacent a number line and diagram of current X positions (marked in parenthesis), the following would take place:

p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>    -4 -3 -2 -1  0  1  2  3  4
p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>                         (0)(1)

p=< 4,0,0>, v=< 1,0,0>, a=<-1,0,0>    -4 -3 -2 -1  0  1  2  3  4
p=< 2,0,0>, v=<-2,0,0>, a=<-2,0,0>                      (1)   (0)

p=< 4,0,0>, v=< 0,0,0>, a=<-1,0,0>    -4 -3 -2 -1  0  1  2  3  4
p=<-2,0,0>, v=<-4,0,0>, a=<-2,0,0>          (1)               (0)

p=< 3,0,0>, v=<-1,0,0>, a=<-1,0,0>    -4 -3 -2 -1  0  1  2  3  4
p=<-8,0,0>, v=<-6,0,0>, a=<-2,0,0>                         (0)   

At this point, particle 1 will never be closer to <0,0,0> than particle 0, and so, in the long run, particle 0 will stay closest.

Which particle will stay closest to position <0,0,0> in the long term?
*/
namespace Day20 {

struct Vector3 {
    Vector3(int X, int Y, int Z) : x(X), y(Y), z(Z) {}
    Vector3(const std::string& str);    // Takes input as "X,Y,Z"
    Vector3() {}
    int Length() const { return abs(x) + abs(y) + abs(z); }
    Vector3& operator+=(const Vector3& rhs);

    int x = 0;
    int y = 0;
    int z = 0;
};

Vector3::Vector3(const std::string& str)
{
    // Takes input as "X,Y,Z"
    auto directions = Helpers::Tokenize(str, ',', true);
    x = std::stoi(directions[0]);
    y = std::stoi(directions[1]);
    z = std::stoi(directions[2]);
}

Vector3& Vector3::operator+=(const Vector3& rhs)
{
    x += rhs.x;
    y += rhs.y;
    z += rhs.z;
    return *this;
}

class Particle {
public:
    Particle(const std::string& str, int id);   // Processes from the format used in the input file

    int Id() const { return m_id; }
    Vector3 Acceleration() const { return m_acceleration; }
    Vector3 Velocity() const { return m_velocity; }
    Vector3 Position() const { return m_position; }

    void Tick();     // Time passes; velocity and position are updated

protected:
    Vector3 m_acceleration;
    Vector3 m_velocity;
    Vector3 m_position;
    int m_id;
};

Particle::Particle(const std::string& str, int id)
    : m_id(id)
{
    auto tokens = Helpers::Tokenize(str);
    // If we get a token that just contains "p=<", take note of its first character, but we can discard the rest as we want to remove it anyway
    char prefixType = '\0';    
    for (auto &t : tokens) {
        if (t[t.size() - 1] == '<') {
            prefixType = t[0];
            continue;
        } 

        // Before we trim the ends off, take note of what vector we're processing
        bool const strippedPrefix = (prefixType != '\0');
        char type = strippedPrefix ? prefixType : t[0];
        prefixType = '\0';
        Helpers::RemoveTrailingCharacter(t, ',');

        // We want to remove p=< from the start of each token and > from the end
        const auto contents = t.substr(strippedPrefix ? 0 : 3, t.size() - 4);
        switch (t[0]) {
        case 'p':
            m_position = Vector3(contents);
            break;
        case 'v':
            m_velocity = Vector3(contents);
            break;
        case 'a':
            m_acceleration = Vector3(contents);
            break;
        }
    }
}

void Particle::Tick()
{
    // Time passes; velocity and position are updated
    m_velocity += m_acceleration;
    m_position += m_velocity;
}

std::vector<Particle> ReadParticles(const std::vector<std::string>& input)
{
    std::vector<Particle> particles;
    for (int i = 0; i < input.size(); ++i) {
        if (!input[i].empty()) particles.emplace_back(input[i], i);
    }
    return particles;
}

// We manipulate the vector, so take a copy
int ClosestToOriginInLongTerm(std::vector<Particle> particles)
{
    // "In the long term" depends entirely on acceleration.  Initial Velocity breaks ties, and initial position breaks ties for that.
    std::sort(particles.begin(), particles.end(),
        [](const Particle& left, const Particle& right) -> bool
    {
        const auto leftAcceleration = left.Acceleration().Length();
        const auto rightAcceleration = right.Acceleration().Length();
        if (leftAcceleration < rightAcceleration) {
            return true;
        } else if (leftAcceleration > rightAcceleration) {
            return false;
        } else {
            const auto leftVelocity = left.Velocity().Length();
            const auto rightVelocity = right.Velocity().Length();
            if (leftVelocity < rightVelocity) {
                return true;
            }
            else if (leftVelocity > rightVelocity) {
                return false;
            }
            else {
                return (left.Position().Length() < right.Position().Length());
            }
        }
    });
    return particles[0].Id();
}

} // namespace Day20

void Day20Tests()
{
    const std::vector<std::string> input = {
        "p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>",
        "p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>"
    };
    const auto particles = Day20::ReadParticles(input);
    const auto closest = Day20::ClosestToOriginInLongTerm(particles);
    const int expectedClosest = 0;
    if (closest != expectedClosest) std::cerr << "Test 20A Error: Got " << closest << ", Expected " << expectedClosest << std::endl;
}

void Day20Problems()
{
    std::cout << "Day 20:\n";
    Day20Tests();
    const auto start = std::chrono::steady_clock::now();
    const auto input = Helpers::ReadFileLines("input_day20.txt");
    const auto particles = Day20::ReadParticles(input);
    const auto closest = Day20::ClosestToOriginInLongTerm(particles);
    const auto end = std::chrono::steady_clock::now();
    std::cout << closest << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl << std::endl;
}