package beckbria.aoc2023.day01

import beckbria.aoc2023.util.readInput

fun part1(input: List<String>): Int {
    return input.map(::firstAndLastDigit).sum()
}

fun part2(input: List<String>): Int {
    return input.map(::numericize).map(::firstAndLastDigit).sum()
}

fun numericize(input: String): String {
    val numbers = mapOf(
        "one" to "1",
        "two" to "2",
        "three" to "3",
        "four" to "4",
        "five" to "5",
        "six" to "6",
        "seven" to "7",
        "eight" to "8",
        "nine" to "9")

    // Find the first and last string number and convert them
    var firstToReplace = ""
    var firstReplaceWith = ""
    var firstNumberIdx = -1;
    var lastToReplace = ""
    var lastReplaceWith = ""
    var lastNumberIdx = -1;

    for (entry in numbers.entries.iterator()) {
        var idx = input.indexOf(entry.key)
        while (idx >= 0) {
            if (idx >= 0 && (idx < firstNumberIdx || firstNumberIdx < 0)) {
                firstNumberIdx = idx;
                firstToReplace = entry.key;
                firstReplaceWith = entry.value;
            }
            if (idx > lastNumberIdx) {
                lastNumberIdx = idx;
                lastToReplace = entry.key;
                lastReplaceWith = entry.value;
            }
            idx = input.indexOf(entry.key, idx + 1)
        }
    }
    var result = input
    if (firstNumberIdx >= 0) {
        result = result.replaceFirst(firstToReplace, firstReplaceWith)
    }
    if (lastNumberIdx >= 0) {
        result = result.replace(lastToReplace, lastReplaceWith)
    }
    return result
}

fun firstAndLastDigit(input: String): Int {
    var d = digits(input)
    if (d.size == 0) {
        return 0
    }
    return d[0] * 10 + d[d.size - 1]
}

fun digits(input: String): List<Int> {
    return input.asIterable().filter{ it >= '0' && it <= '9' }.map(Char::digitToInt).toList()
}

fun main() {
    val testInput = readInput("2023/01/input.txt")
    //val testInput = listOf("1abc2","pqr3stu8vwx","a1b2c3d4e5f","treb7uchet")
    //val testInput = listOf("two1nine","eightwothree","abcone2threexyz","xtwone3four","4nineeightseven2","zoneight234","7pqrstsixteen")
    println("Part 1:")
    println(part1(testInput))
    println("\nPart 2:")
    println(part2(testInput))
}