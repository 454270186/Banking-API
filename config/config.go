package config

type ServerConfig struct {
	ProgramName  string         `mapstructure:"name"`
	Port         int            `mapstructure:"port"`
	Host         string         `mapstructure:"host"`
	PostgresInfo PostgresConfig `mapstructure:"postgres"`
}

type PostgresConfig struct {
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBHost     string `mapstructure:"db_host"`
	DBPort     int    `mapstructure:"db_port"`
	DBName     string `mapstructure:"db_name"`
}
