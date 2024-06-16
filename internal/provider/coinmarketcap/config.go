package coinmarketcap

import "time"

type Config struct {
	APIUrl     string        `toml:"APIUrl"`
	APIKey     string        `toml:"APIKey"`
	TimeoutSec time.Duration `toml:"TimeoutSec"`
}
