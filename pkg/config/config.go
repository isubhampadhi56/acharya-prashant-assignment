package config

type Config struct {
	accessTokenSecret  []byte
	accessTokenExpiry  int
	refreshTokenSecret []byte
	refreshTokenExpiry int
	port               int
}

func (c *Config) GetAccessTokenSecret() []byte {
	return c.accessTokenSecret
}
func (c *Config) GetRefreshTokenSecret() []byte {
	return c.refreshTokenSecret
}
func (c *Config) GetAccessTokenExpiry() int {
	return c.accessTokenExpiry
}
func (c *Config) GetRefreshTokenExpiry() int {
	return c.refreshTokenExpiry
}
func (c *Config) GetAppPort() int {
	return c.port
}

var config *Config

func GetConfig() *Config {
	if config != nil {
		return config
	}
	config = &Config{
		accessTokenSecret:  []byte("hello"),
		refreshTokenSecret: []byte("hello123"),
		accessTokenExpiry:  5,
		refreshTokenExpiry: 72,
		port:               8080,
	}
	return config
}
