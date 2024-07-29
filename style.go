package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

var (
  // Terminal Dimensions
  termWidth, _, _ = term.GetSize(os.Stdout.Fd())

  // Colors
  defaultColor = lipgloss.Color("#e5c5a6")
  subtitleColor = lipgloss.Color("#8f7f70")
  highlightColor = lipgloss.Color("#c2a07f")
  suitColors = map[string]lipgloss.Color{
    "Spades": lipgloss.Color("#5a9ad6"),
    "Clubs": lipgloss.Color("#87d65a"),
    "Hearts": lipgloss.Color("#f7795a"),
    "Diamonds": lipgloss.Color("#f7b45a"),
  }
  defaultStyle = lipgloss.NewStyle().Foreground(defaultColor)

  // Card Styles
  cardFmt = "%2s\n\n%5s\n\n%8s "
  cardStyle = lipgloss.NewStyle().
    BorderStyle(lipgloss.RoundedBorder()).
    MarginTop(1)
  unsetCard = cardStyle.Foreground(subtitleColor).BorderForeground(subtitleColor)
  setCard = cardStyle.Foreground(defaultColor).BorderForeground(defaultColor)
  highlightedCard = cardStyle.Foreground(highlightColor).BorderForeground(highlightColor)
  currentCard = setCard.
    MarginTop(15)

  helpStyle = lipgloss.NewStyle().
    Align(lipgloss.Center).
    Foreground(subtitleColor).
    Faint(true)
)

func CardFormat(card Card) string {
  rank := card.Rank
  symbol, color := card.SuitSymbol()
  suit := lipgloss.NewStyle().Foreground(color).SetString(symbol)

  cardFmt := fmt.Sprintf("%2s\n\n%5s\n\n%8s ", rank, fmt.Sprintf("\t%s", suit), rank)
  return cardFmt
}

func Results(scoredHands [2][]string, board string, score int) string {
  detailsStyle := lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).
                  BorderForeground(defaultColor).
                  Foreground(defaultColor).Margin(3, 2).Padding(1)
  scoreStyle := detailsStyle.Width(40).AlignHorizontal(lipgloss.Center).Padding(2, 4).UnsetBorderStyle()

  rowsLabeled := append([]string{"\tRows 👉"}, scoredHands[0]...)
  colsLabeled := append([]string{"\tColumns 👇"}, scoredHands[1]...)
  helpText := [2]lipgloss.Style{defaultStyle.Foreground(defaultColor), defaultStyle.Foreground(subtitleColor)}
  finalScore := strings.Join(
    []string{
      fmt.Sprintf("Score: %d\n", score), 
      fmt.Sprint(
        helpText[0].Render("p "), helpText[1].Render("Restart\t"),
        helpText[0].Render("ctrl+c/q "), helpText[1].Render("Quit")),
    },
  "\n")
  scoredRows := strings.Join(rowsLabeled, "\n\n")
  scoredCols := strings.Join(colsLabeled, "\n\n")

  s := lipgloss.JoinVertical(lipgloss.Top, lipgloss.JoinHorizontal(lipgloss.Left, detailsStyle.Render(scoredRows), detailsStyle.Render(scoredCols)), scoreStyle.Render(finalScore))
  return lipgloss.JoinHorizontal(lipgloss.Left, s, board)
}
