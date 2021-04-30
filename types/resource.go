package types

type Resource struct {
	Typ  string
	Name string
}

func (r Resource) Flattify() string {
	return r.Typ + "." + r.Name
}
