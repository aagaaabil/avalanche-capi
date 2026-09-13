package config

type Config struct {
	Coin       string
	Path       string
	Prefix     string
	StorageDir string
}

func Default() Config {
	return Config{
		Coin:       "AVAX",
		Path:       "m/44'/60'/0'",
		Prefix:     "0x",
		StorageDir: ".wallets",
	}
}
