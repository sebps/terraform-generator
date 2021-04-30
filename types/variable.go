package types

type Variable struct {
	Name string
}

func (v Variable) Flattify() string {
	return "var." + v.Name
}
