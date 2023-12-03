package brianbeck.aoc2023.day01

import brianbeck.aoc2023.util.readInput

fun main() {
    fun part1(input: List<String>): Int {
        return input.size
    }

    fun part2(input: List<String>): Int {
        return input.size
    }

    val testInput = readInput("2023/01/input.txt")
    println(testInput)
    println("Part 1:")
    println(part1(testInput))
    println("\nPart 2:")
    println(part2(testInput))
}