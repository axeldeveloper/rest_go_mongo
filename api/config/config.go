package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	hosts      string
	database   string
	username   string
	password   string
	collection string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			hosts:      "localhost:27017",
			database:   "mestre",
			username:   "",
			password:   "",
			collection: "Employee",
		},
	}
}

func GetUri() {

}
