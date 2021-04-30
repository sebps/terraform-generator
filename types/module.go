package types

type Module struct {
	Name string
}

func (m Module) Flattify() string {
	return "module." + m.Name
}
