#include "Problems.h"
/*
--- Day 20: Particle Swarm ---

Suddenly, the GPU contacts you, asking for help. Someone has asked it to simulate too many particles, 
and it won't be able to finish them all in time to render the next frame at this rate.

It transmits to you a buffer (your puzzle input) listing each particle in order (starting with particle 0, 
then particle 1, particle 2, and so on). For each particle, it provides the X, Y, and Z coordinates for the 
particle's position (p), velocity (v), and acceleration (a), each in the format <X,Y,Z>.

Each tick, all particles are updated simultaneously. A particle's properties are updated in the following order:

    Increase the X velocity by the X acceleration.
    Increase the Y velocity by the Y acceleration.
    Increase the Z velocity by the Z acceleration.
    Increase the X position by the X velocity.
    Increase the Y position by the Y velocity.
    Increase the Z position by the Z velocity.

Because of seemingly tenuous rationale involving z-buffering, the GPU would like to know which particle will 
stay closest to position <0,0,0> in the long term. Measure this using the Manhattan distance, which in this 
situation is simply the sum of the absolute values of a particle's X, Y, and Z position.

For example, suppose you are only given two particles, both of which stay entirely on the X-axis (for 
simplicity). Drawing the current states of particles 0 and 1 (in that order) with an adjacent a number 
line and diagram of current X positions (marked in parenthesis), the following would take place:

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

--- Part Two ---

To simplify the problem further, the GPU would like to remove any particles that collide. Particles collide if their positions ever 
exactly match. Because particles are updated simultaneously, more than two particles can collide at the same time and place. Once 
particles collide, they are removed and cannot collide with anything else after that tick.

For example:

p=<-6,0,0>, v=< 3,0,0>, a=< 0,0,0>
p=<-4,0,0>, v=< 2,0,0>, a=< 0,0,0>    -6 -5 -4 -3 -2 -1  0  1  2  3
p=<-2,0,0>, v=< 1,0,0>, a=< 0,0,0>    (0)   (1)   (2)            (3)
p=< 3,0,0>, v=<-1,0,0>, a=< 0,0,0>

p=<-3,0,0>, v=< 3,0,0>, a=< 0,0,0>
p=<-2,0,0>, v=< 2,0,0>, a=< 0,0,0>    -6 -5 -4 -3 -2 -1  0  1  2  3
p=<-1,0,0>, v=< 1,0,0>, a=< 0,0,0>             (0)(1)(2)      (3)
p=< 2,0,0>, v=<-1,0,0>, a=< 0,0,0>

p=< 0,0,0>, v=< 3,0,0>, a=< 0,0,0>
p=< 0,0,0>, v=< 2,0,0>, a=< 0,0,0>    -6 -5 -4 -3 -2 -1  0  1  2  3
p=< 0,0,0>, v=< 1,0,0>, a=< 0,0,0>                       X (3)
p=< 1,0,0>, v=<-1,0,0>, a=< 0,0,0>

------destroyed by collision------
------destroyed by collision------    -6 -5 -4 -3 -2 -1  0  1  2  3
------destroyed by collision------                      (3)
p=< 0,0,0>, v=<-1,0,0>, a=< 0,0,0>

In this example, particles 0, 1, and 2 are simultaneously destroyed at the time and place marked X. On the next tick, particle 3 passes 
through unharmed.

How many particles are left after all collisions are resolved?
*/
namespace Day20 {

struct Vector3 {
    Vector3(int64_t X, int64_t Y, int64_t Z) : x(X), y(Y), z(Z) {}
    Vector3(const std::string& str);    // Takes input as "X,Y,Z"
    Vector3() {}
    int64_t Length() const { return abs(x) + abs(y) + abs(z); }
    Vector3& operator+=(const Vector3& rhs);

    // Orders by X, then Y, then Z.
    friend bool operator< (const Vector3& left, const Vector3& right);

    int64_t x = 0;
    int64_t y = 0;
    int64_t z = 0;
};

class Particle {
public:
    Particle(const std::string& str, unsigned int id);   // Processes from the format used in the input file
    Particle() {}

    unsigned int Id() const { return m_id; }
    Vector3 Acceleration() const { return m_acceleration; }
    Vector3 Velocity() const { return m_velocity; }
    Vector3 Position() const { return m_position; }

    // Orders by acceleration, then velocity, then position.  Used for sorting
    friend bool operator< (const Particle& left, const Particle& right);   

    void Tick();     // Time passes; velocity and position are updated

protected:
    Vector3 m_acceleration;
    Vector3 m_velocity;
    Vector3 m_position;
    unsigned int m_id = -1;
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

bool operator< (const Vector3& left, const Vector3& right)
{
    if (left.x < right.x) {
        return true;
    } else if (left.x > right.x) {
        return false;
    } else {
        if (left.y < right.y) {
            return true;
        } else if (left.y > right.y) {
            return false;
        } else {
            return (left.z < right.z);
        }
    }
}

Particle::Particle(const std::string& str, unsigned int id)
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

        const auto contents = strippedPrefix ? t.substr(0, t.size() - 1) : t.substr(3, t.size() - 4);
        switch (type) {
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

bool operator< (const Particle& left, const Particle& right)
{
    const auto leftAcceleration = left.Acceleration().Length();
    const auto rightAcceleration = right.Acceleration().Length();
    if (leftAcceleration < rightAcceleration) {
        return true;
    }
    else if (leftAcceleration > rightAcceleration) {
        return false;
    }
    else {
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
}

std::vector<Particle> ReadParticles(const std::vector<std::string>& input)
{
    std::vector<Particle> particles;
    for (unsigned int i = 0; i < input.size(); ++i) {
        if (!input[i].empty()) particles.emplace_back(input[i], i);
    }
    return particles;
}

// We manipulate the vector, so take a copy
int64_t ClosestToOriginInLongTerm(std::vector<Particle> particles)
{
    // "In the long term" depends entirely on acceleration.  Initial Velocity breaks ties, and initial position breaks ties for that.
    std::sort(particles.begin(), particles.end());
    return particles[0].Id();
}

inline bool SameDirection(const Vector3& left, const Vector3& right)
{
    return ((left.x < 0) == (right.x < 0)) && ((left.y < 0) == (right.y < 0)) && ((left.z < 0) == (right.z < 0));
}

bool WillNeverCollide(const std::vector<Particle>& particles)
{
    // All collisions are complete once:
    //   * The particles' distance from the origin is in the same order as the magnitude of their acceleration. (that is,
    //         the fastest accelerating particles are the furthest away
    //   * The particles' velocity is in the same direction as their acceleration (they're not slowing down)
    //   * The particles' velocity is in the same direction as their position (they're not moving back towards the origin)
    // Once we have reached that state, no particle can reach any other particle.  

    // TODO: This can take quite a while to get to a known end state.  The right solution may involve vector analysis.

    if (particles.size() < 2) return true;

    int64_t previousDistance = INT64_MIN;
    int64_t previousVelocity = INT64_MIN;
    for (const auto &p : particles) {
        const auto currentDistance = p.Position().Length();
        if ((currentDistance < previousDistance) || 
            !SameDirection(p.Acceleration(), p.Velocity()) || 
            !SameDirection(p.Velocity(), p.Position())) return false;
        previousDistance = currentDistance;
    }
    return true;
}

int64_t RemainingAfterCollisions(std::vector<Particle> particles)
{
    std::sort(particles.begin(), particles.end());
    //while (!WillNeverCollide(particles)) {
    // TODO: Finish debugging the collision mechanism.  In practice, our values are small enough that 1000 iterations is more than enough
    for (int i = 0; i < 1000; ++i) {
        // Any particles with duplicate positions are destroyed
        std::map<Vector3, std::list<unsigned int>> positions;
        for (unsigned int i = 0; i < particles.size(); ++i) {
            positions[particles[i].Position()].push_back(i);
        }

        std::vector<unsigned int> indexesToErase;
        for (const auto& pos : positions) {
            if (pos.second.size() > 1) {
                // There are duplicates here.  Make note of them to be erased
                indexesToErase.insert(indexesToErase.end(), pos.second.begin(), pos.second.end());
            }
        }
        Helpers::RemoveIndexes(particles, indexesToErase);

        // Advance the particles
        for (auto &p : particles) p.Tick();
    }
    return particles.size();
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
    const int64_t expectedClosest = 0;
    if (closest != expectedClosest) std::cerr << "Test 20A Error: Got " << closest << ", Expected " << expectedClosest << std::endl;

    const std::vector<std::string> partBInput = {
        "p=<-6,0,0>, v=< 3,0,0>, a=< 0,0,0>",
        "p=<-4,0,0>, v=< 2,0,0>, a=< 0,0,0>",
        "p=<-2,0,0>, v=< 1,0,0>, a=< 0,0,0>",
        "p=< 3,0,0>, v=<-1,0,0>, a=< 0,0,0>"
    };
    const auto particlesB = Day20::ReadParticles(partBInput);
    const auto remaining = Day20::RemainingAfterCollisions(particlesB);
    const auto expectedRemaining = 1;
    if (remaining != expectedRemaining) std::cerr << "Test 20B Error: Got " << remaining << ", Expected " << expectedRemaining << std::endl;
}

void Day20Problems()
{
    std::cout << "Day 20:\n";
    Day20Tests();
    const auto start = std::chrono::steady_clock::now();
    const auto input = Helpers::ReadFileLines("input_day20.txt");
    const auto particles = Day20::ReadParticles(input);
    const auto closest = Day20::ClosestToOriginInLongTerm(particles);
    const auto remaining = Day20::RemainingAfterCollisions(particles);
    const auto end = std::chrono::steady_clock::now();
    std::cout << closest << std::endl;
    std::cout << remaining << std::endl;
    std::cout << "Took " << std::chrono::duration<double, std::milli>(end - start).count() << " ms" << std::endl << std::endl;
}