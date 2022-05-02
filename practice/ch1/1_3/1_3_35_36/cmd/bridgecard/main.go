package main

import (
	"algo/practice/ch1/1_3/1_3_35_36/randqueue"
	"fmt"
)

type suitType int

const (
	club suitType = iota
	heart
	spade
	diamond
)

func (t suitType) dump() string {
	switch t {
	case club:
		return "♣"
	case heart:
		return "♥"
	case spade:
		return "♠"
	case diamond:
		return "♦"
	}
	panic("invalid")
}

type card struct {
	suit   suitType
	number int
}

func (c *card) dump() string {
	return fmt.Sprintf("{%s, %d}", c.suit.dump(), c.number)
}

// bridge cards
func main() {
	cards := &randqueue.RandQueue[card]{}
	// 13 * 4 cards
	for suit := suitType(0); suit <= diamond; suit++ {
		for i := 1; i <= 13; i++ {
			cards.Enqueue(card{suit: suit, number: i})
		}
	}
	var player int
	cards.DoAndDequeue(func(c card) bool {
		fmt.Printf("player%d: %s\n", player+1, c.dump())
		player = (player + 1) % 4
		return false
	})
}
