package config

//best approach will be loading from config.json or something like that
//but for the sake of simplicity will be hardcoded

const BASE_URL string = "https://api.tiendanube.com/v1/382262/"
const USER_AGENT string = "GoCrazy (cyn@gocrazy.com.ar)"
const BEARER_TOKEN string = "bearer e23247347ffb0f7e2983bec04b991834442b8905"

type Config struct {
	BaseUrl     string
	UserAgent   string
	BearerToken string
}

func CreateConfig() *Config {
	return &Config{
		BaseUrl:     BASE_URL,
		UserAgent:   USER_AGENT,
		BearerToken: BEARER_TOKEN,
	}
}
