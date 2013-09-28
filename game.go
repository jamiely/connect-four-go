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

