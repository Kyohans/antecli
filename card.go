package main

import (
	"errors"
	"math/rand"
	"sort"

  "github.com/charmbracelet/lipgloss"
)

type Card struct {
  Suit string
  Rank string
  Val int
}

func (c Card) SuitSymbol() (string, lipgloss.Color) {
  switch(c.Suit) {
    case "Spades": return "♠", suitColors["Spades"]
    case "Clubs": return "♣", suitColors["Clubs"]
    case "Diamonds": return "♦", suitColors["Diamonds"]
    case "Hearts": return "♥", suitColors["Hearts"]
    default: return "?", subtitleColor
  }
}

var FaceCardHierarchy map[string]int = map[string]int{ "K": 2, "Q": 1, "J": 0 }
var FullDeck []Card = []Card{
  {"Spades", "A", 11}, {"Clubs", "A", 11}, {"Hearts", "A", 11}, {"Diamonds", "A", 11},
  {"Spades", "K", 10}, {"Clubs", "K", 10}, {"Hearts", "K", 10}, {"Diamonds", "K", 10},
  {"Spades", "Q", 10}, {"Clubs", "Q", 10}, {"Hearts", "Q", 10}, {"Diamonds", "Q", 10},
  {"Spades", "J", 10}, {"Clubs", "J", 10}, {"Hearts", "J", 10}, {"Diamonds", "J", 10},
  {"Spades", "10", 10}, {"Clubs", "10", 10}, {"Hearts", "10", 10}, {"Diamonds", "10", 10},
  {"Spades", "9", 9}, {"Clubs", "9", 9}, {"Hearts", "9", 9}, {"Diamonds", "9", 9},
  {"Spades", "8", 8}, {"Clubs", "8", 8}, {"Hearts", "8", 8}, {"Diamonds", "8", 8},
  {"Spades", "7", 7}, {"Clubs", "7", 7}, {"Hearts", "7", 7}, {"Diamonds", "7", 7},
  {"Spades", "6", 6}, {"Clubs", "6", 6}, {"Hearts", "6", 6}, {"Diamonds", "6", 6},
  {"Spades", "5", 5}, {"Clubs", "5", 5}, {"Hearts", "5", 5}, {"Diamonds", "5", 5},
  {"Spades", "4", 4}, {"Clubs", "4", 4}, {"Hearts", "4", 4}, {"Diamonds", "4", 4},
  {"Spades", "3", 3}, {"Clubs", "3", 3}, {"Hearts", "3", 3}, {"Diamonds", "3", 3},
  {"Spades", "2", 2}, {"Clubs", "2", 2}, {"Hearts", "2", 2}, {"Diamonds", "2", 2},
}
var BlankCard Card = Card{" ", " ", 0}

type Hand []Card
func (h Hand) Len() int { return len(h) }
func (h Hand) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Hand) Less(i, j int) bool { 
  valI, isFaceI := FaceCardHierarchy[h[i].Rank]
  valJ, isFaceJ := FaceCardHierarchy[h[j].Rank]
  if isFaceI && isFaceJ {
    return valI < valJ
  }
  return h[i].Val < h[j].Val
}

func InitDeck() ([]Card, error) {
  if DECK_SIZE > 52 {
    return nil, errors.New("InitDeck(): Deck size cannot be greater than 52")
  }
  m := make(map[int]bool, 0)
  cards := make([]Card, 0, DECK_SIZE)
  for len(cards) < DECK_SIZE {
    idx := rand.Intn(52)
    if _, exists := m[idx]; !exists {
      m[idx] = true
      cards = append(cards, FullDeck[idx])
    }
  }
  return cards, nil
}

func (h Hand) ScoreHand() int {
  m := make(map[string]int, 0)
  sort.Sort(h)

  handInfo := map[string]int{
    "consecutive": 1,
    "max_occurring_rank": 0,
    "max_occurring_suit": 0,
    "three_of_a_kind": 0,
    "four_of_a_kind": 0,
    "twos": 0,
  }
  for i, card := range h {
    if i > 0 {
      left, faceLeft := FaceCardHierarchy[card.Rank]
      right, faceRight := FaceCardHierarchy[card.Rank]
      if faceLeft && faceRight {
        if left == right - 1 {
          handInfo["consecutive"]++
        } else {
          handInfo["consecutive"] = 1
        }
      } else {
        if (card.Rank == "Ace" && h[0].Rank == "Two") || (h[i-1].Val == card.Val - 1) {
          handInfo["consecutive"]++
        } else {
          handInfo["consecutive"] = 1
        }
      }
    }

    m[card.Suit]++
    m[card.Rank]++

    switch m[card.Rank] {
      case 3:
        handInfo["three_of_a_kind"] = 1
      case 4:
        handInfo["four_of_a_kind"] = 1
    }

    if m[card.Suit] > handInfo["max_occurring_suit"] {
      handInfo["max_occurring_suit"] = m[card.Suit]
    }

    if m[card.Rank] > handInfo["max_occurring_rank"] {
      handInfo["max_occurring_rank"] = m[card.Rank]
    }
  }

  // Check for two pairs (doesn't work in first loop as it can conflict with 3oaK)
  for i, card := range h {
    if i > 0 && h[i-1].Rank == card.Rank {
      continue
    }

    if m[card.Rank] == 2 {
      handInfo["twos"]++
    }
  }

  var flush bool = false
  if handInfo["max_occurring_suit"] == 5 {
    flush = true
  }

  var straight bool = false
  if handInfo["consecutive"] == 5 && handInfo["max_occurring_rank"] == 1 {
    straight = true
  }

  if straight || flush {
    if straight && flush {
      return STRAIGHT_FLUSH_SCORE
    } else if straight {
      return STRAIGHT_SCORE
    } else if flush {
      return FLUSH_SCORE
    }
  } else if handInfo["four_of_a_kind"] > 0 {
    return FOUR_OF_A_KIND_SCORE
  } else if handInfo["three_of_a_kind"] > 0 {
    if handInfo["twos"] > 0 {
      return FULL_HOUSE_SCORE
    }
    return THREE_OF_A_KIND_SCORE
  } else if val, _ :=  handInfo["twos"]; val > 0 {
    if val > 1 {
      return TWO_PAIR_SCORE
    }
    return PAIR_SCORE
  }

  return HIGH_CARD_SCORE
}
