import re
from collections import defaultdict
from z3 import *

NANOBOT_REGEX = re.compile("^pos=<(-?\\d+),(-?\\d+),(-?\\d+)>, r=(\\d+)$")

class Nanobot:
    x = 0
    y = 0
    z = 0
    r = 0 # Radius of impact

    def __init__(self, input):
        match = NANOBOT_REGEX.match(input)
        self.x = int(match.group(1))
        self.y = int(match.group(2))
        self.z = int(match.group(3))
        self.r = int(match.group(4))

    def distance(self, other):
        return abs(self.x - other.x) + abs(self.y - other.y) + abs(self.z - other.z)

    def inRange(self, other):
        return self.distance(other) <= self.r

def inRangeOfStrongest(bots):
    strongest = max(bots, key=lambda b: b.r)
    return len(list(filter(lambda x: strongest.inRange(x), bots)))

def absZ3(x):
    return If(x < 0,-x,x)

def optimalLocationDistance(bots):
    # Throw everything at the SAT-solver
    o = Optimize()
    # Define variables
    x = Int('x') # Location X
    y = Int('y') # Location Y
    z = Int('z') # Location Z
    distance = Int('distance')  # Distance of our found location from zero
    inRange = Int('inRange')    # Number of nanobots in range of our current location
    o.add(distance == absZ3(x) + absZ3(y) + absZ3(z)) # Manhattan distance of point to (0,0,0)

    # Keep track of the individual components of the score
    constraints = []
    for i in range(len(bots)):
        b = bots[i]
        constraint = Int('constraint%d' % i)    # Each variable requires a unique name
        # TODO: Should probably factor this logic out into some sort of Z3 distance function, which
        # would imply moving the location into a Vector3-style class
        o.add(constraint == If(absZ3(x - b.x) + absZ3(y - b.y) + absZ3(z - b.z) <= b.r, 1, 0))
        constraints.append(constraint)
    o.add(inRange == sum(constraints))
    # Add the scoring criteria
    maxRangeConstraint = o.maximize(inRange)        # Maximize the number of points in rage
    minDistanceConstraint = o.minimize(distance)    # As a tiebreaker, minimize the distance from zero
    # Find an optimal solution
    o.check()
    xVal = o.model()[x].as_long()
    yVal = o.model()[y].as_long()
    zVal = o.model()[z].as_long()
    return [xVal, yVal, zVal]

if __name__ == "__main__":
    with open("input.txt", 'r') as input_file:
        file_content = input_file.readlines()
        nanobots = list(map(lambda x: Nanobot(x), file_content))
        print("Part 1: %d" % (inRangeOfStrongest(nanobots)))
        pt = optimalLocationDistance(nanobots)
        print("Part 2: (%d,%d,%d) = %d" % (pt[0], pt[1], pt[2], abs(pt[0]) + abs(pt[1]) + abs(pt[2])))
