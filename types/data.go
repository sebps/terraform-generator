package types

type Data struct {
	Typ  string
	Name string
}

func (d Data) Flattify() string {
	return "data." + d.Typ + "." + d.Name
}
