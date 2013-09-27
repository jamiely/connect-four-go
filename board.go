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

type Index struct {
  row, column int
}

func NewIndex(r,c int) *Index {
  return &Index{row: r, column: c}
}

func (board *Board) Indices() []*Index {
  indices := make([]*Index, board.width * board.height)
  for i := 0; i < board.height; i++ {
    for j := 0; j < board.width; j++ {
      indices[i * j + j] = NewIndex(i, j)
    }
  }
  return indices
}


