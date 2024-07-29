package main

import (
  "fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type pos struct {
  x int
  y int
}

type Model struct {
  g *Game
  cursor pos
  done bool
}

func InitModel() Model {
  deck, err := InitDeck()
  if err != nil {
    panic(err)
  }

  return Model{
    g: &Game{
      NumPlaced: 0,
      Score: 0,
      Deck: deck,
      Board: Grid{},
      CurrentCard: 0,
    },
    cursor: pos{0, 0},
    done: false,
  }
}

func (m Model) Init() tea.Cmd {
  return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
      case "ctrl+c", "q":
        return m, tea.Quit
      case "p":
        if m.done {
          restart := InitModel()
          return restart, nil
        }
      case "h", "left":
        if m.cursor.y > 0 {
          m.cursor.y--
        }
      case "l", "right":
        if m.cursor.y < N-1 {
          m.cursor.y++
        }
      case "k", "up":
        if m.cursor.x > 0 {
          m.cursor.x--
        }
      case "j", "down":
        if m.cursor.x < N-1 {
          m.cursor.x++
        }
      case "enter", " ", "a":
        if !m.done {
          cardSymbol, _ := m.g.Board[m.cursor.x][m.cursor.y].SuitSymbol()
          if cardSymbol == "?" && m.g.CurrentCard < 25 {
            m.g.PlaceCard(m.cursor)
          }

          if m.g.NumPlaced == 25 {
            m.g.TallyScore()
            m.done = true
          }
        }
    }
  }

  return m, nil
}

func (m Model) View() string {
  var s string

  if m.g.CurrentCard < 25 {
    symbol, color := m.g.Deck[m.g.CurrentCard].SuitSymbol()
    style := defaultStyle.Foreground(color).SetString(symbol).Render()

    cFmt := fmt.Sprintf(cardFmt, m.g.Deck[m.g.CurrentCard].Rank, fmt.Sprintf("\t%s", style), m.g.Deck[m.g.CurrentCard].Rank)
    helpText := fmt.Sprintf("\n%s%s%s%s%s%s",
      defaultStyle.Render("\ta/enter/space "), defaultStyle.Foreground(subtitleColor).Render("Place Card\t"),
      defaultStyle.Render("hjkl/←↓↑→ "), defaultStyle.Foreground(subtitleColor).Render("Move Cursor\t"),
      defaultStyle.Render("ctrl+c/q "), defaultStyle.Foreground(subtitleColor).Render("Quit"),
    )

    s = lipgloss.JoinHorizontal(lipgloss.Left, currentCard.SetString(cFmt).Render(), m.g.BoardState(m.cursor))
    s = lipgloss.JoinVertical(lipgloss.Top, s, helpText)
  }

  if m.done {
    s = Results(m.g.ScoredHands, m.g.BoardState(pos{-1,-1}), m.g.Score)
  }

  return s
}
