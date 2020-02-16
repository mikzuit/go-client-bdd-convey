package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
	Host	string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect: "postgres",
			Username: "root",
			Password: "password",
			Name: "interview_accountapi",
			Charset: "utf8",
			Host: "192.168.160.3",
		},
	}
}
