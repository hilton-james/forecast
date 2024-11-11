package config

func Default() Config {
	return Config{
		Debug: false,
		Forecast: Forecast{
			ListenPort: ":5000",
		},
	}
}
