package connect_four

import (
  "fmt"
  "strings"
)

type Direction struct {
  x,y int
}

type Game struct {
  board *Board
  directions []*Direction
  current Marker
  gameOver bool
}

func NewDirections() []*Direction {
  directions := make([]*Direction, 8)
  counter := 0
  for i:=-1; i<=1; i++ {
    for j:=-1; j<=1; j++ {
      if(i != 0 || j != 0) {
        directions[counter] = NewDirection(i, j)
        counter = counter + 1
      }
    }
  }

  return directions
}

func NewGame() *Game {
  return &Game{board: NewDefaultBoard(), 
               directions: NewDirections(),
               current: A,
               gameOver: false}
}

func NewDirection(x int, y int) *Direction {
  return &Direction{x: x, y: y}
}

func (game *Game) IndexValid(index *Index) bool {
  return game.board.IndexValid(index)
}

func (game *Game) CheckPosition(index *Index, marker Marker, 
  direction *Direction, steps int) bool {
  if steps == 0 {
    return true
  } else if game.CheckMarkerAt(index, marker) {
    return game.CheckPosition(IndexInDirection(index, direction),
                              marker,
                              direction,
                              steps - 1)
  } else {
    return false
  }
}

func (game *Game) CheckMarkerAt(index *Index, marker Marker) bool {
  return game.IndexValid(index) && game.MarkerAt(index) == marker
}

func (game *Game) HasMoves() bool {
  return game.board.HasMoves()
}

func IndexInDirection(index *Index, dir *Direction) *Index {
  if index == nil || dir == nil { 
    return nil
  }
  return NewIndex(index.row + dir.y, index.column + dir.x)
}

func (game *Game) MarkerAt(index *Index) Marker {
  return game.board.MarkerAt(index)
}

func (game *Game) DropIntoColumn(marker Marker, column int) *Index {
  index := game.FirstEmptyRowIn(column)
  if index == nil { 
    return nil
  }

  game.board.SetMarkerAt(index, marker)
  return index
}

func (game *Game) FirstEmptyRowIn(column int) *Index {
  for i := 0; i < game.board.height; i ++ {
    index := NewIndex(i, column)
    if game.CheckMarkerAt(index, EMPTY) {
      return index
    }
  }
  return nil
}

func (game *Game) ToggleMarker() Marker {
  if game.current == A { 
    game.current = B
  } else { 
    game.current = A
  }
  return game.current
}

func (game *Game) Move(column int) bool {
  result := game.DropIntoColumn(game.current, column) != nil
  if result {
    game.ToggleMarker()
  }
  return result
}

func (game *Game) BoardIndices() []*Index {
  return game.board.Indices()
}

/**
 * Returns an index from which there is a win.
 */
func (game *Game) IsWin() *Index {
  indices := game.BoardIndices()
  for i := range indices {
    if game.WinAtIndex(indices[i]) {
      return indices[i]
    }
  }

  return nil
}

func (game *Game) WinAtIndex(index *Index) bool {
  if ! game.IndexValid(index) {
    return false
  }
  marker := game.MarkerAt(index)
  if marker == EMPTY {
    return false
  }

  for i := range game.directions {
    if game.CheckPosition(index, marker, game.directions[i], 4) {
      return true
    }
  }

  return false
}

func (game *Game) ToString() string {
  board := game.board
  parts := make([]string, board.height)
  header := "  "
  for j := 0; j < board.width; j++ {
    header = fmt.Sprintf("%s%d", header, j)
  }
  header += "\n"
  for i := 0; i < board.height; i ++ {
    str := fmt.Sprintf("%d ", i)
    for j := 0; j < board.width; j ++ {
      index := NewIndex(i, j)
      str = fmt.Sprintf("%s%s", str, MarkerToString(game.MarkerAt(index)))
    }
    parts[board.height - 1 - i] = str
  }
  return header + strings.Join(parts, "\n") + "\n" + header
}

func MarkerToString(marker Marker) string {
  switch marker {
  case EMPTY: 
    return "."
  case A: 
    return "X"
  case B:
    return "O"
  default:
    return "?"
  }
}

