package beckbria.aoc2023.day05

import beckbria.aoc2023.util.readInput

fun part1(input: List<String>): Int {
    return input.size
}

fun part2(input: List<String>): Int {
    return input.size
}

fun bothParts(testInput: List<String>) {
    println("Part 1:")
    println(part1(testInput))
    println("\nPart 2:")
    println(part2(testInput))
}

fun main() {
    val testInputExample = readInput("2023/05/input_example.txt")
    val testInput = readInput("2023/05/input.txt")
    println("Example Input:")
    bothParts(testInputExample)
    println("\nYour Input:")
    bothParts(testInput)
}

