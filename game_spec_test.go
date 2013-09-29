package connect_four

import (
	. "github.com/orfjackal/gospec/src/gospec"
)

func GameSpec(c Context) {
  origin := NewIndex(0, 0)
  east := NewDirection(1, 0)

  c.Specify("It should have directions.", func() {
    c.Expect(len(NewGame().directions), Equals, 8)
  })
  c.Specify("should return true when checkPosition is called with 0 steps", func() {
    c.Expect(NewGame().CheckPosition(origin, A, east, 0), IsTrue)
  })
  c.Specify("should return true when checkPosition is called with an index and the marker at that index and 1 step", func() {
    c.Expect(true, IsFalse)
  })
  c.Specify("should return true when checkPosition is called with an index and the marker at that index, the same marker beneath it, and 2 steps", func() {
    c.Expect(true, IsFalse)
  })
  c.Specify("can check whether there are markers in a row", func() {
    c.Expect(true, IsFalse)
  })
  c.Specify("should be a player's win whenever there are four in a row", func() {
    c.Expect(true, IsFalse)
  })
  c.Specify("not allow a move to be made in a full column", func() {
    c.Expect(true, IsFalse)
  })
}

