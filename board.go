package connect_four

type Marker int

const (
  Empty Marker = iota
  A
  B
)

type Board struct {
  spaces [][]Marker
  width, height int
}

func NewDefaultBoard() *Board {
  return NewBoard(7, 6)
}
func NewBoard(width, height int) *Board {
  spaces := make([][]Marker, height)
  for i := range spaces {
    spaces[i] = make([]Marker, width)
  }
  return &Board{spaces: spaces, width: width, height: height}
}


