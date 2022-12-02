package main

import "fmt"

type Pick byte

const (
	Rock Pick = iota
	Paper
	Scissors
)

func (p Pick) winsOver() Pick {
	switch p {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	}
	panic("unknown pick")
}

func (p Pick) loosesTo() Pick {
	switch p {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	}
	panic("unknown pick")
}

func (p Pick) score() int {
	switch p {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	panic("unknown pick")
}

type Outcome byte

const (
	Loose Outcome = iota
	Draw
	Win
)

type Play struct {
	opponent Pick
	you      Pick
}

type Play2 struct {
	opponent Pick
	outcome  Outcome
}

func dec2() {
	lines, _ := readLines("dec2.txt")

	var plays []Play
	for _, line := range lines {
		plays = append(plays, getPlay(line))
	}

	var plays2 []Play2
	for _, line := range lines {
		plays2 = append(plays2, getPlay2(line))
	}

	dec2_1(plays)
	dec2_2(plays2)
}

func dec2_1(plays []Play) {
	var score int
	for _, play := range plays {
		score = score + play.you.score()
		score = score + scoreResult(play.you, play.opponent)
	}
	fmt.Println("dec2_1", score)
}

func dec2_2(plays []Play2) {
	var score int
	for _, play := range plays {
		pickToPlay := pickToPlay(play.outcome, play.opponent)
		score = score + pickToPlay.score()
		score = score + scoreResult(pickToPlay, play.opponent)
	}
	fmt.Println("dec2_2", score)
}

func pickToPlay(outcome Outcome, opponent Pick) Pick {
	switch outcome {
	case Loose:
		return opponent.winsOver()
	case Draw:
		return opponent
	case Win:
		return opponent.loosesTo()
	}
	panic("unknown pick")
}

func getPlay(line string) Play {
	return Play{opponent: newPick(line[0]), you: newPick(line[2])}
}

func getPlay2(line string) Play2 {
	return Play2{opponent: newPick(line[0]), outcome: newOutcome(line[2])}
}

func newOutcome(u uint8) Outcome {
	switch u {
	case 'X':
		return Loose
	case 'Y':
		return Draw
	case 'Z':
		return Win
	}
	panic("unknown pick")
}

func newPick(u uint8) Pick {
	switch u {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	case 'C', 'Z':
		return Scissors
	}
	panic("unknown pick")
}

func scoreResult(you Pick, opponent Pick) int {
	if you == opponent {
		return 3
	}
	if you.winsOver() == opponent {
		return 6
	}
	return 0
}
