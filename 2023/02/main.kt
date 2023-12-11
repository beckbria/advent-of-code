package beckbria.aoc2023.day02

import beckbria.aoc2023.util.readInput
import kotlin.math.max

class Turn(var count: Map<String, Int>) {
    companion object {
        /**
         * Parse AoC Input such as `4 red, 8 green` into a map of counts
         */
        fun fromString(input: String): Turn {
            val count = mutableMapOf<String, Int>()
            val regex = """(\d+) ([a-z]+)""".toRegex()
            for (raw in input.split(",")) {
                val (num, name) = regex.find(raw.trim())!!.destructured
                count.put(name, num.toInt())
            }
            return Turn(count)
        }
    }

    fun valid(max: Turn): Boolean {
        for (entry in max.count.entries.iterator()) {
            if (entry.value < count.getOrDefault(entry.key, 0)) {
                return false
            }
        }
        return true
    }

    fun power(): Int {
        val p = count.getOrDefault("red", 0) * count.getOrDefault("green", 0) * count.getOrDefault("blue", 0)
        return p
    }

    override fun toString(): String = "" + count
}

class Game(var id: Int, var turns: List<Turn>) {
    companion object {
        /**
         * Parse AoC Input such as `Game 1: 4 red, 8 green; 8 green, 6 red` into a Game object
         */
        fun fromString(input: String): Game {
            val regex = """Game (\d+): (.*)""".toRegex()
            val (id, rawTurns) = regex.find(input)!!.destructured
            val turns = rawTurns.split(";").map{ Turn.fromString(it) }
            return Game(id.toInt(), turns)
        }
    }

    fun minCubes(): Turn {
        var red = 0
        var green = 0
        var blue = 0
        turns.forEach{
            red = max(red, it.count.getOrDefault("red", 0))
            green = max(green, it.count.getOrDefault("green", 0))
            blue = max(blue, it.count.getOrDefault("blue", 0))
        }
        return Turn(mapOf("red" to red, "green" to green, "blue" to blue))
    }

    override fun toString(): String = "Game " + id + ": " + turns
}

fun part1(games: List<Game>): Int {
    val max = Turn(mapOf("red" to 12, "green" to 13, "blue" to 14))
    return games.filter{ it.turns.all{ it.valid(max) } }.map{ it.id }.sum()
}

fun part2(games: List<Game>): Int {
    return games.map{ it.minCubes().power() }.sum()
}

fun bothParts(testInput: List<String>) {
    val games = testInput.map{ Game.fromString(it) }
    println("Part 1:")
    println(part1(games))
    println("\nPart 2:")
    println(part2(games))
}

fun main() {
    val testInputExample = readInput("2023/02/input_example.txt")
    val testInput = readInput("2023/02/input.txt")
    println("Example Input:")
    bothParts(testInputExample)
    println("\nYour Input:")
    bothParts(testInput)
}