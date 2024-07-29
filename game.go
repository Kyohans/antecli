package main

import (
	"github.com/charmbracelet/lipgloss"
)

const (
  N int = 5
  DECK_SIZE = 25
  STRAIGHT_FLUSH_SCORE int = 75
  FOUR_OF_A_KIND_SCORE int = 50
  FULL_HOUSE_SCORE int = 25
  FLUSH_SCORE int = 20
  STRAIGHT_SCORE int = 15
  THREE_OF_A_KIND_SCORE int = 10
  TWO_PAIR_SCORE int = 5
  PAIR_SCORE int = 2
  HIGH_CARD_SCORE int = 1
)

type Grid [N][N]Card
type Game struct {
  Board Grid
  Deck []Card
  CurrentCard int
  NumPlaced int
  ScoredHands [2][]string // [scoredRows, scoredCols]
  Score int
}

var Hand2Str map[int]string = map[int]string{
  HIGH_CARD_SCORE: "High Card (+1)",
  PAIR_SCORE: "Pair (+2)",
  TWO_PAIR_SCORE: "Two Pair (+5)",
  THREE_OF_A_KIND_SCORE: "Three Of A Kind (+10)",
  STRAIGHT_SCORE: "Straight (+15)",
  FLUSH_SCORE: "Flush (+20)",
  FULL_HOUSE_SCORE: "Full House (+25)",
  FOUR_OF_A_KIND_SCORE: "Four Of A Kind (+50)",
  STRAIGHT_FLUSH_SCORE: "Straight Flush (+75)",
}

func (g *Game) TallyScore() (error) {
  var rows, cols int
  for _, row := range g.Board {
    var hand Hand
    for _, card := range row {
      hand = append(hand, card)
    }
    score := hand.ScoreHand()
    g.ScoredHands[0] = append(g.ScoredHands[0], Hand2Str[score])
    rows += score
  }

  for i := 0; i < N; i++ {
    var hand Hand
    for row := 0; row < N; row++ {
      hand = append(hand, g.Board[row][i])
    }
    score := hand.ScoreHand()
    g.ScoredHands[1] = append(g.ScoredHands[1], Hand2Str[score])
    cols += score
  }

  g.Score = rows + cols
  return nil
}

func (g *Game) PlaceCard(cursor pos) {
  g.Board[cursor.x][cursor.y] = g.Deck[g.CurrentCard]
  g.CurrentCard++
  g.NumPlaced++
}

func (g *Game) BoardState(cursor pos) string {
  s := ""
  rows := ""
  for i, row := range g.Board {
    rowStyle := []string{}
    for j, card := range row {
      c := ""
      cardFmt := ""
      if card.Rank != "" {
        cardFmt = CardFormat(card)
        if i == cursor.x && j == cursor.y {
          c = highlightedCard.Render(cardFmt)
        } else {
          c = setCard.Render(cardFmt)
        }
      } else {
        cardFmt = CardFormat(BlankCard)
        if i == cursor.x && j == cursor.y {
          c = highlightedCard.Render(cardFmt)
        }else {
          c = unsetCard.Render(cardFmt)
        }
      }

      rowStyle = append(rowStyle, c)
    }
    rows += lipgloss.JoinHorizontal(lipgloss.Bottom, rowStyle...)
  }
  s = lipgloss.JoinVertical(lipgloss.Top, rows)
  return s
}
