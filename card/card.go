package card

import (
	"errors"
	"fmt"
)

const (
	SuitDiamonds        = "diamonds"
	SuitDiamondsUnicode = "\u2666"

	SuitClubs        = "clubs"
	SuitClubsUnicode = "\u2663"

	SuitHearts        = "hearts"
	SuitHeartsUnicode = "\u2665"

	SuitSpades        = "spades"
	SuitSpadesUnicode = "\u2660"
)

const (
	Face2     = "2"
	Face3     = "3"
	Face4     = "4"
	Face5     = "5"
	Face6     = "6"
	Face7     = "7"
	Face8     = "8"
	Face9     = "9"
	Face10    = "10"
	FaceJack  = "J"
	FaceQueen = "Q"
	FaceKing  = "K"
	FaceAce   = "A"
)

type Card struct {
	Suit string
	Face string
}

func isValidSuit(suit string) bool {
	switch suit {
	case SuitClubs, SuitDiamonds, SuitHearts, SuitSpades:
		return true
	default:
		return false
	}
}

func isValidFace(face string) bool {
	switch face {
	case Face2, Face3, Face4, Face5, Face6, Face7, Face8, Face9, Face10, FaceJack, FaceQueen, FaceKing, FaceAce:
		return true
	default:
		return false
	}
}

func (c Card) SuitUnicode() (string, error) {
	switch c.Suit {
	case SuitClubs:
		return SuitClubsUnicode, nil
	case SuitSpades:
		return SuitSpadesUnicode, nil
	case SuitHearts:
		return SuitHeartsUnicode, nil
	case SuitDiamonds:
		return SuitDiamondsUnicode, nil
	default:
		return "", errors.New(fmt.Sprintf("unrecognized suit %s", c.Suit))
	}
}

func New(suit string, face string) (*Card, error) {
	if isValidSuit(suit) && isValidFace(face) {
		return &Card{
			Suit: suit,
			Face: face,
		}, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Cannot construct Card with suit %s, face %s", suit, face))
	}
}