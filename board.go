package connect_four

type Marker int

const (
  EMPTY Marker = iota
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
    for j := range spaces[i] {
      spaces[i][j] = EMPTY
    }
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
  counter := 0
  for i := 0; i < board.height; i++ {
    for j := 0; j < board.width; j++ {
      indices[counter] = NewIndex(i, j)
      counter = counter + 1
    }
  }
  return indices
}

func (board *Board) SetMarkerAt(index *Index, marker Marker) {
  board.spaces[index.row][index.column] = marker
}

func (board *Board) MarkerAt(index *Index) Marker {
  return board.spaces[index.row][index.column]
}

func (board *Board) IsEmpty() bool{
  indices := board.Indices()
  accumulator := true
  for i := range indices {
    index := indices[i]
    accumulator = accumulator && board.MarkerAt(index) == EMPTY
  }
  return accumulator
}

func (board *Board) HasMoves() bool {
  indices := board.Indices()
  accumulator := true
  for i := range indices {
    index := indices[i]
    accumulator = accumulator && board.MarkerAt(index) != EMPTY
  }
  return !accumulator
}

func (board *Board) IndexValid(index *Index) bool {
  return index != nil &&
         index.column < board.width &&
         index.column >= 0 &&
         index.row < board.height &&
         index.row >= 0
}

