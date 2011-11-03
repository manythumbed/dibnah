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

func (h Heading) Names() []string {
	n := []string{}
	for k, _ := range h.attributes {
		n = append(n, k)
	}

	return n
}

func (h Heading) Equals(that Heading) bool {
	if len(h.attributes) == len(that.attributes)	{
		for k, v := range h.attributes	{
			if that.attributes[k].Name != v.Name {
				return false
			}
		}
		return true
	}
	return false
}

type Tuple struct {
	Heading Heading
	Values map[string] Value
}

func (t Tuple) Equals(that Tuple) bool {
	return false
}
