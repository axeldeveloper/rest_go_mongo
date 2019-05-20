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
	//ds061375.mlab.com
	//mongodb://axelmacnamara:Axel@193@ds061375.mlab.com:61375/mestre

	return &Config{
		DB: &DBConfig{
			hosts:      "localhost:27017",
			database:   "movies_db",
			username:   "",
			password:   "",
			collection: "userdetails",
		},
	}
}

func GetUri() {

}
