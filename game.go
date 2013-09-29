package connect_four

type Direction struct {
  x,y int
}

type Game struct {
  board *Board
  directions []*Direction
}

func NewDirections() []*Direction {
  directions := make([]*Direction, 8)
  counter := 0
  for i:=-1; i<=1; i++ {
    for j:=-1; j<=1; j++ {
      if(i != 0 && j != 0) {
        directions[counter] = NewDirection(i, j)
        counter = counter + 1
      }
    }
  }

  return directions
}

func NewGame() *Game {
  return &Game{board: NewDefaultBoard(), directions: NewDirections()}
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

func IndexInDirection(index *Index, dir *Direction) *Index {
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

