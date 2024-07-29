package main

import (
	"os"
  "fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  p := tea.NewProgram(InitModel())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Something went wrong :( -> %v", err)
    os.Exit(1)
  }
}

