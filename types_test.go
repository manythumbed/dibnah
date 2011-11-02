package dibnah

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T)	{
	TestingT(t)
}

type S struct {
	attributes map[string] Type
}

func(s *S) SetUpSuite(c *C)	{
	s.attributes = map[string] Type {
		"A": Type{"A"},
		"B": Type{"B"},
		"A1": Type{"A"},
	}
}

var _ = Suite(&S{})

func (s *S) TestEmptyHeading(c *C) {
	h := Heading{}
	c.Check(h.Degree(), Equals, 0)
	c.Check(len(h.Types()), Equals, 0)
}

func (s *S) TestHeading(c *C) {
	h := Heading{s.attributes}
	c.Check(h.Degree(), Equals, len(s.attributes))
	c.Check(len(h.Types()), Equals, 3)
}
