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
	c.Check(h.Equals(h), Equals, true)
}

func (s *S) TestHeading(c *C) {
	h := Heading{s.attributes}
	c.Check(h.Degree(), Equals, len(s.attributes))
	c.Check(len(h.Types()), Equals, 3)
	c.Check(h.Equals(h), Equals, true)
}

func (s *S) TestHeadingEquality(c *C)	{
	h := Heading{s.attributes}
	c.Check(h.Equals(h), Equals, true)

	h1 := Heading{}
	c.Check(h.Equals(h1), Equals, false)
	c.Check(h1.Equals(h), Equals, false)

	a1 := map[string] Type {
		"a" : Type{"a"},
		"b" : Type{"b"},
	}
	a2 := map[string] Type {
		"a" : Type{"b"},
		"b" : Type{"a"},
	}

	h2 := Heading{a1}
	c.Check(h2.Equals(h2), Equals, true)

	h3 := Heading{a2}
	c.Check(h3.Equals(h3), Equals, true)

	c.Check(h2.Equals(h3), Equals, false)
	c.Check(h3.Equals(h2), Equals, false)

	
}
