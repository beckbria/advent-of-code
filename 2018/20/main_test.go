package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPattern(t *testing.T) {
	assert.Equal(t, SinglePattern("WNE"), ReadPattern("^WNE$"))

	assert.Equal(t,
		ConcatenatePattern(
			SinglePattern("ENWWW"),
			ChoicePattern(
				SinglePattern("NEEE"),
				ConcatenatePattern(
					SinglePattern("SSE"),
					ChoicePattern(SinglePattern("EE"), SinglePattern("N"))))),
		ReadPattern("^ENWWW(NEEE|SSE(EE|N))$"))

	assert.Equal(t,
		ConcatenatePattern(
			SinglePattern("ENNWSWW"),
			ChoicePattern(SinglePattern("NEWS"), SinglePattern("")),
			SinglePattern("SSSEEN"),
			ChoicePattern(SinglePattern("WNSE"), SinglePattern("")),
			SinglePattern("EE"),
			ChoicePattern(SinglePattern("SWEN"), SinglePattern("")),
			SinglePattern("NNN")),
		ReadPattern("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"))

	assert.Equal(t,
		ConcatenatePattern(
			SinglePattern("ESSWWN"),
			ChoicePattern(
				SinglePattern("E"),
				ConcatenatePattern(
					SinglePattern("NNENN"),
					ChoicePattern(
						ConcatenatePattern(
							SinglePattern("EESS"),
							ChoicePattern(SinglePattern("WNSE"), SinglePattern("")),
							SinglePattern("SSS")),
						ConcatenatePattern(
							SinglePattern("WWWSSSSE"),
							ChoicePattern(SinglePattern("SW"), SinglePattern("NNNE"))))))),
		ReadPattern("^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"))

	assert.Equal(t,
		ConcatenatePattern(
			SinglePattern("WSSEESWWWNW"),
			ChoicePattern(
				SinglePattern("S"),
				ConcatenatePattern(
					SinglePattern("NENNEEEENN"),
					ChoicePattern(
						ConcatenatePattern(
							SinglePattern("ESSSSW"),
							ChoicePattern(SinglePattern("NWSW"), SinglePattern("SSEN"))),
						ConcatenatePattern(
							SinglePattern("WSWWN"),
							ChoicePattern(
								SinglePattern("E"),
								ConcatenatePattern(
									SinglePattern("WWS"),
									ChoicePattern(SinglePattern("E"), SinglePattern("SS"))))))))),
		ReadPattern("^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$"))
}

/*
func inputToMap(input []string) RoomMap {
	m := make(RoomMap)

	// Find the origin
	var origin point
findorigin:
	for y, s := range input {
		for x, c := range []rune(s) {
			if c == 'X' {
				origin.x, origin.y = x, y
				break findorigin
			}
		}
	}

	for y, s := range input {
		for x, c := range []rune(s) {
			if (y == origin.y) && (x == origin.x) {
				// Treat this like a normal room
				m.set(0, 0, Room)
			} else {
				m.set(x-origin.x, y-origin.y, Component(c))
			}
		}
	}
	return m
}

func patternToMap(pattern string) RoomMap {
	p := ReadPattern(pattern)
	return BuildMap(&p)
}

func TestBuildMap(t *testing.T) {
	assert.Equal(t, inputToMap([]string{
		"#####",
		"#.|.#",
		"#-###",
		"#.|X#",
		"#####",
	}), patternToMap("^WNE$"))

	assert.Equal(t, inputToMap([]string{
		"#########",
		"#.|.|.|.#",
		"#-#######",
		"#.|.|.|.#",
		"#-#####-#",
		"#.#.#X|.#",
		"#-#-#####",
		"#.|.|.|.#",
		"#########",
	}), patternToMap("^ENWWW(NEEE|SSE(EE|N))$"))

	assert.Equal(t, inputToMap([]string{
		"###########",
		"#.|.#.|.#.#",
		"#-###-#-#-#",
		"#.|.|.#.#.#",
		"#-#####-#-#",
		"#.#.#X|.#.#",
		"#-#-#####-#",
		"#.#.|.|.|.#",
		"#-###-###-#",
		"#.|.|.#.|.#",
		"###########",
	}), patternToMap("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"))

	assert.Equal(t, inputToMap([]string{
		"#############",
		"#.|.|.|.|.|.#",
		"#-#####-###-#",
		"#.#.|.#.#.#.#",
		"#-#-###-#-#-#",
		"#.#.#.|.#.|.#",
		"#-#-#-#####-#",
		"#.#.#.#X|.#.#",
		"#-#-#-###-#-#",
		"#.|.#.|.#.#.#",
		"###-#-###-#-#",
		"#.|.#.|.|.#.#",
		"#############",
	}), patternToMap("^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"))

	assert.Equal(t, inputToMap([]string{
		"###############",
		"#.|.|.|.#.|.|.#",
		"#-###-###-#-#-#",
		"#.|.#.|.|.#.#.#",
		"#-#########-#-#",
		"#.#.|.|.|.|.#.#",
		"#-#-#########-#",
		"#.#.#.|X#.|.#.#",
		"###-#-###-#-#-#",
		"#.|.#.#.|.#.|.#",
		"#-###-#####-###",
		"#.|.#.|.|.#.#.#",
		"#-#-#####-#-#-#",
		"#.#.|.|.|.#.|.#",
		"###############",
	}), patternToMap("^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$"))
}

func TestMostDoors(t *testing.T) {
	assert.Equal(t, 3, MostDoors("^WNE$"))
	assert.Equal(t, 10, MostDoors("^ENWWW(NEEE|SSE(EE|N))$"))
	assert.Equal(t, 18, MostDoors("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"))
	assert.Equal(t, 23, MostDoors("^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"))
	assert.Equal(t, 31, MostDoors("^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$"))
}
*/
