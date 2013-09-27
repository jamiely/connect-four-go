package connect_four

import (
	"github.com/orfjackal/gospec/src/gospec"
  "testing"
)

func TestAllSpecs(t *testing.T) {
  r := gospec.NewRunner()
  r.AddSpec(BoardSpec)
  gospec.MainGoTest(r, t)
}
