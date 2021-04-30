package types

type Data struct {
	Typ  string
	Name string
}

func (d Data) Flattify() string {
	return d.Typ + "." + d.Name
}
