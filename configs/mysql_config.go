package configs

type MySQLConfig struct {
	UserName     string
	UserPassword string
	DatabaseName string
	Port         string
}

func newMySQLConfig() MySQLConfig {
	userName, err1 := getEnv("MYSQL_USER_NAME")
	if err1 != nil {
		panic(err1)
	}

	userPassword, err2 := getEnv("MYSQL_USER_PASSWORD")
	if err2 != nil {
		panic(err2)
	}

	databaseName, err3 := getEnv("MYSQL_DATABASE_NAME")
	if err3 != nil {
		panic(err3)
	}

	port, err4 := getEnv("MYSQL_PORT")
	if err4 != nil {
		panic(err4)
	}

	return MySQLConfig{
		UserName:     userName,
		UserPassword: userPassword,
		DatabaseName: databaseName,
		Port:         port,
	}
}
