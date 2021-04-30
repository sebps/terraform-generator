package types

type Type struct {
	Name string
}

func (t Type) Flattify() string {
	return t.Name
}
