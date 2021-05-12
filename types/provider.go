package types

type Provider struct {
	Registry  string `default: "registry.terraform.io"`
	Namespace string
	Name      string
	Version   string
}

func (p Provider) Flattify() string {
	return p.Namespace + "/" + p.Name
}
