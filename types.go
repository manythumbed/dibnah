package dibnah

type Type struct {
	Name string
}

type Value struct {}

type Heading struct {
	attributes map[string] Type
}

func (h Heading) Degree() int {
	return len(h.attributes)
}

func (h Heading) Types() []Type {
	t :=  []Type{}
	for _, v := range h.attributes {
		t = append(t, v)
	}

	return t
}

type Tuple struct {
	Heading Heading
	Values map[string] Value
}
