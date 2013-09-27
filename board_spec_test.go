package connect_four

import (
	"github.com/orfjackal/gospec/src/gospec"
	. "github.com/orfjackal/gospec/src/gospec"
)

func BoardSpec(c gospec.Context) {
  c.Specify("It should have a size of 7 columns and 6 rows", func() {
    board := NewDefaultBoard()
    c.Expect(board.width, Equals, 7)
    c.Expect(board.height, Equals, 6)
  })
  c.Specify("It should be empty", func() {
    c.Expect(true, gospec.IsFalse)
  })
  c.Specify("It should have width * height indices", func() {
    c.Expect(true, gospec.IsFalse)
  })
  c.Specify("It should return true for empty check", func() {
    c.Expect(true, gospec.IsFalse)
  })
  c.Specify("It should keep track of markers at indices", func() {
    c.Expect(true, gospec.IsFalse)
  })
  c.Specify("It should know when there are no moves left", func() {
    c.Expect(true, gospec.IsFalse)
  })
}

