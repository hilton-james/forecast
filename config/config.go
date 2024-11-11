package config

type Config struct {
	Debug    bool     `koanf:"debug"`
	Forecast Forecast `koanf:"forecast"`
}
