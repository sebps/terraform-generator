package types

type Provider struct {
	Name string
}

func (p Provider) Flattify() string {
	return "provider." + p.Name
}
