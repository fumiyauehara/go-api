package model

type DBConfig struct {
	Type   string
	DBHost string
	DBPort int
	DBName string
	DBUser string
	DBPw   string
}

type RedisConfig struct {
	RedisHost string
	RedisPort int
}

type Config struct {
	Postgres    DBConfig
	MysqlWriter DBConfig
	MysqlReader DBConfig
	RedisConfig
	Port int
}
