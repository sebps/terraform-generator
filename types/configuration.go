package types

type Configuration struct {
	Resources     []*Resource
	ResourceTypes []*Type
	Data          []*Data
	DataTypes     []*Type
	Modules       []*Module
	Variables     []*Variable
}
