package connect_four

import (
	. "github.com/orfjackal/gospec/src/gospec"
)

func BoardSpec(c Context) {
	c.Specify("It should have a size of 7 columns and 6 rows", func() {
		board := NewDefaultBoard()
		c.Expect(board.width, Equals, 7)
		c.Expect(board.height, Equals, 6)
	})
	c.Specify("It should have width * height indices", func() {
		c.Expect(len(NewDefaultBoard().Indices()), Equals, 42)
	})
	c.Specify("It should return true for empty check", func() {
		board := NewDefaultBoard()
		c.Expect(board.IsEmpty(), IsTrue)
		board.SetMarkerAt(NewIndex(0, 0), A)
		c.Expect(board.IsEmpty(), IsFalse)
	})
	c.Specify("It should keep track of markers at indices", func() {
		board := NewDefaultBoard()
		index1 := NewIndex(0, 0)
		index2 := NewIndex(0, 1)
		board.SetMarkerAt(index1, A)
		board.SetMarkerAt(index2, B)
		c.Expect(board.MarkerAt(index1), Equals, A)
	})
	c.Specify("It should know when there are no moves left", func() {
		board := NewDefaultBoard()
		c.Expect(board.HasMoves(), IsTrue)

		indices := board.Indices()
		for i := range indices {
			board.SetMarkerAt(indices[i], A)
		}

		c.Expect(board.HasMoves(), IsFalse)
	})
}
