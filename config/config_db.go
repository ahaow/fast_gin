package config

type Database struct {
	Mode         string
	Dsn          string
	MaxIdleConns int
	MaxOpenCons  int
}
