package main

import "testing"

func assert(ok bool) bool {
  return ok
}

/*
  Straight Flush (75): Consecutive ranks and same suit
  Four Of A Kind (50): Four of the same rank
  Full House (25): Three of a kind and a pair (2-2-2-K-K)
  Flush (20): Hand contains the same suit
  Straight (15): Consecutive ranks
  Three Of A Kind (10): Three of the same rank
  Two Pair (5): Two pairs each with similar rank (6-6-J-J)
  Pair (2): Two cards that share the same rank
  High Card (1): No Pair
*/

func assertHand(hand Hand, expected int, t *testing.T) {
  score := hand.ScoreHand()
  if assert(score != expected) {
    t.Errorf("Failed: Expected %d, got %d\n", expected, score)
  }
}

func TestStraightFlush(t *testing.T) {
  hand := Hand{
    {"Spades", "Ace", 11}, {"Spades", "Five", 5}, {"Spades", "Four", 4}, {"Spades", "Three", 3}, {"Spades", "Two", 2}, 
  }
  assertHand(hand, STRAIGHT_FLUSH_SCORE, t)
}

func TestFourOfAKind(t *testing.T) {
  hand := Hand{
    {"Clubs", "King", 10}, {"Spades", "King", 10}, {"Diamonds", "King", 10}, {"Hearts", "King", 10}, {"Spades", "Two", 2}, 
  }
  assertHand(hand, FOUR_OF_A_KIND_SCORE, t)
}

func TestFullHouse(t *testing.T) {
  hand := Hand{
    {"Clubs", "King", 10}, {"Spades", "King", 10}, {"Diamonds", "King", 10}, {"Hearts", "Eight", 8}, {"Spades", "Eight", 8}, 
  }
  assertHand(hand, FULL_HOUSE_SCORE, t)
}

func TestFlush(t *testing.T) {
  hand := Hand{
    {"Clubs", "Ace", 11}, {"Clubs", "Three", 3}, {"Clubs", "King", 10}, {"Clubs", "Three", 3}, {"Clubs", "Two", 2}, 
  }
  assertHand(hand, FLUSH_SCORE, t)
}

func TestStraight(t *testing.T) {
  hand := Hand{
    {"Spades", "Ace", 11}, {"Clubs", "Five", 5}, {"Diamonds", "Four", 4}, {"Hearts", "Three", 3}, {"Spades", "Two", 2}, 
  }
  assertHand(hand, STRAIGHT_SCORE, t)
}

func TestThreeOfAKind(t *testing.T) {
  hand := Hand{
    {"Clubs", "King", 10}, {"Spades", "King", 10}, {"Diamonds", "King", 10}, {"Hearts", "Ace", 11}, {"Spades", "Two", 2}, 
  }
  assertHand(hand, THREE_OF_A_KIND_SCORE, t)
}

func TestTwoPair(t *testing.T) {
  hand := Hand{
    {"Clubs", "King", 10}, {"Spades", "King", 10}, {"Diamonds", "Queen", 10}, {"Hearts", "Queen", 10}, {"Spades", "Two", 2}, 
  }
  assertHand(hand, TWO_PAIR_SCORE, t)
}

func TestPair(t *testing.T) {
  hand := Hand{
    {"Clubs", "King", 10}, {"Spades", "King", 10}, {"Diamonds", "Queen", 10}, {"Hearts", "Jack", 10}, {"Spades", "Two", 2}, 
  }
  assertHand(hand, PAIR_SCORE, t)
}

func TestHighCard(t *testing.T) {
  hand := Hand{
    {"Clubs", "King", 10}, {"Spades", "Ace", 11}, {"Diamonds", "Queen", 10}, {"Hearts", "Jack", 10}, {"Spades", "Two", 2}, 
  }
  assertHand(hand, HIGH_CARD_SCORE, t)
}
