package config

const (
	Dialect  = "postgres"
	Host     = "localhost"
	DbPort   = "5432"
	User     = "nmathur"
	DbName   = "to_do"
	Password = "abc123"
)

type Todo struct {
	Task    string `json:"task"`
	Checked bool   `json:"checked"`
}
