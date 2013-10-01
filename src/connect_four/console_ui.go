package connect_four

import (
  "fmt"
)

func console_ui() {
  game := NewGame()
  var column int
  someoneWon := false
  hasMoves := true
  for !someoneWon && hasMoves {
    fmt.Printf("%s\nWhat column would you like to play? ", game.ToString())
    fmt.Scanf("%d", &column)
    if !game.Move(column) {
      fmt.Printf("Invalid move.")
    }
    someoneWon = game.IsWin() != nil
    hasMoves = game.HasMoves()
  }

  if someoneWon {
    fmt.Printf("Someone won!\n%s", game.ToString())
  } else if hasMoves {
    fmt.Printf("There are no moves left")
  }
}

