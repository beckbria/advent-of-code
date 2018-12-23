# Day 23: Experimental Emergency Teleportation

[https://adventofcode.com/2018/day/23#part2](https://adventofcode.com/2018/day/23#part2)

## Description

### Part One

Using your torch to search the darkness of the rocky cavern, you finally locate the man's friend: a small _reindeer_.

You're not sure how it got so far in this cave. It looks sick - too sick to walk - and too heavy for you to carry all the way back. Sleighs won't be invented for another 1500 years, of course.

The only option is _experimental emergency teleportation_.

You hit the "experimental emergency teleportation" <span title="We've always had this button; we've just been too scared to press it.">button</span> on the device and push _I accept the risk_ on no fewer than 18 different warning messages. Immediately, the device deploys hundreds of tiny _nanobots_ which fly around the cavern, apparently assembling themselves into a very specific _formation_. The device lists the `X,Y,Z` position (`pos`) for each nanobot as well as its _signal radius_ (`r`) on its tiny screen (your puzzle input).

Each nanobot can transmit signals to any integer coordinate which is a distance away from it _less than or equal to_ its signal radius (as measured by [Manhattan distance](https://en.wikipedia.org/wiki/Taxicab_geometry)). Coordinates a distance away of less than or equal to a nanobot's signal radius are said to be _in range_ of that nanobot.

Before you start the teleportation process, you should determine which nanobot is the _strongest_ (that is, which has the largest signal radius) and then, for that nanobot, the _total number of nanobots that are in range_ of it, _including itself_.

For example, given the following nanobots:

    pos=<0,0,0>, r=4
    pos=<1,0,0>, r=1
    pos=<4,0,0>, r=3
    pos=<0,2,0>, r=1
    pos=<0,5,0>, r=3
    pos=<0,0,3>, r=1
    pos=<1,1,1>, r=1
    pos=<1,1,2>, r=1
    pos=<1,3,1>, r=1
    

The strongest nanobot is the first one (position `0,0,0`) because its signal radius, `4` is the largest. Using that nanobot's location and signal radius, the following nanobots are in or out of range:

*   The nanobot at `0,0,0` is distance `0` away, and so it is _in range_.
*   The nanobot at `1,0,0` is distance `1` away, and so it is _in range_.
*   The nanobot at `4,0,0` is distance `4` away, and so it is _in range_.
*   The nanobot at `0,2,0` is distance `2` away, and so it is _in range_.
*   The nanobot at `0,5,0` is distance `5` away, and so it is _not_ in range.
*   The nanobot at `0,0,3` is distance `3` away, and so it is _in range_.
*   The nanobot at `1,1,1` is distance `3` away, and so it is _in range_.
*   The nanobot at `1,1,2` is distance `4` away, and so it is _in range_.
*   The nanobot at `1,3,1` is distance `5` away, and so it is _not_ in range.

In this example, in total, `_7_` nanobots are in range of the nanobot with the largest signal radius.

Find the nanobot with the largest signal radius. _How many nanobots are in range_ of its signals?

### Part Two

Now, you just need to figure out where to position yourself so that you're actually teleported when the nanobots activate.

To increase the probability of success, you need to find the coordinate which puts you _in range of the largest number of nanobots_. If there are multiple, choose one _closest to your position_ (`0,0,0`, measured by manhattan distance).

For example, given the following nanobot formation:

    pos=<10,12,12>, r=2
    pos=<12,14,12>, r=2
    pos=<16,12,12>, r=4
    pos=<14,14,14>, r=6
    pos=<50,50,50>, r=200
    pos=<10,10,10>, r=5
    

Many coordinates are in range of some of the nanobots in this formation. However, only the coordinate `12,12,12` is in range of the most nanobots: it is in range of the first five, but is not in range of the nanobot at `10,10,10`. (All other coordinates are in range of fewer than five nanobots.) This coordinate's distance from `0,0,0` is `_36_`.

Find the coordinates that are in range of the largest number of nanobots. _What is the shortest manhattan distance between any of those points and `0,0,0`?_
