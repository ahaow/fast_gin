package config

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database Database
	Redis    Redis
	Jwt      Jwt
}
