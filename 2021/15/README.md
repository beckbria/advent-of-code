# Day 15: Chiton

[https://adventofcode.com/2021/day/15](https://adventofcode.com/2021/day/15)

## Description

### Part One

You've almost reached the exit of the cave, but the walls are getting closer together. Your submarine can barely still fit, though; the main problem is that the walls of the cave are covered in [chitons](https://en.wikipedia.org/wiki/Chiton), and it would be best not to bump any of them.

The cavern is large, but has a very low ceiling, restricting your motion to two dimensions. The shape of the cavern resembles a square; a quick scan of chiton density produces a map of _risk level_ throughout the cave (your puzzle input). For example:

    1163751742
    1381373672
    2136511328
    3694931569
    7463417111
    1319128137
    1359912421
    3125421639
    1293138521
    2311944581
    

You start in the top left position, your destination is the bottom right position, and you <span title="Can't go diagonal until we can repair the caterpillar unit. Could be the liquid helium or the superconductors.">cannot move diagonally</span>. The number at each position is its _risk level_; to determine the total risk of an entire path, add up the risk levels of each position you _enter_ (that is, don't count the risk level of your starting position unless you enter it; leaving it adds no risk to your total).

Your goal is to find a path with the _lowest total risk_. In this example, a path with the lowest total risk is highlighted here:

    1163751742
    1381373672
    2136511328
    3694931569
    7463417111
    1319128137
    1359912421
    3125421639
    1293138521
    2311944581
    

The total risk of this path is _`40`_ (the starting position is never entered, so its risk is not counted).

_What is the lowest total risk of any path from the top left to the bottom right?_
