package dibnah

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T)	{
	TestingT(t)
}

type S struct {}

var _ = Suite(&S{})

func (s *S) TestEmptyHeading(c *C) {
	h := Heading{}
	c.Check(h.Degree(), Equals, 0)
}

func (s *S) TestHeading(c *C) {
	attributes := map[string] Type {
		"A": Type{"A"},
		"B": Type{"A"},
		"A1": Type{"A"},
	}
	h := Heading{attributes}
	c.Check(h.Degree(), Equals, len(attributes))
}
