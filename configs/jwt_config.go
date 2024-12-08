package configs

type JWTConfig struct {
	Secret string
}

func newJWTConfig() JWTConfig {
	secret, err := getEnv("JWT_SECRET")

	if err != nil {
		panic(err)
	}

	return JWTConfig{
		Secret: secret,
	}
}
