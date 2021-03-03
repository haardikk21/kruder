package config

// Config defines the configuration for the Kruder API
type Config struct {
	LogLevel    string `default:"DEBUG"`
	Environment string `default:"DEV"`

	RouterConfig
	DatabaseConfig
}

// RouterConfig defines the configuration for the router
type RouterConfig struct {
	Port int `default:"6000"`
}

// DatabaseConfig defines the configuration for the database
type DatabaseConfig struct {
	DBAddress  string `required: "true"`
	DBPassword string `required: "true"`
	DBUsername string `required: "true"`
	DBName     string `required: "true"`
}
