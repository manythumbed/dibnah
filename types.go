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

type Tuple struct {
	Heading Heading
	Values map[string] Value
}
