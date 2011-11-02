package dibnah

type Type struct {
	Name string
}

type Attribute struct {
	Type Type
	Name string
}

type Value struct {}

type Heading struct {
	attributes map[string] Attribute
}

type Tuple struct {
	Heading Heading
	Values map[string] Value
}
