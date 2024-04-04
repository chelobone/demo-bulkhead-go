package model

type Client struct {
	ID       int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Nombre   string `boil:"nombre" json:"nombre" toml:"nombre" yaml:"nombre"`
	Apellido string `boil:"apellido" json:"apellido" toml:"apellido" yaml:"apellido"`
	Direcion string `boil:"direccion" json:"direccion" toml:"direccion" yaml:"direccion"`
	Correo   string `boil:"correo" json:"correo" toml:"correo" yaml:"correo"`
}

type ClientSlice []*Client

type QueryConfig struct {
	PageSize int
	Page     int
}

type Result struct {
	statusCode string
	message    string
}
