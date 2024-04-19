package gopot

// Config
type Config struct {
	// pot server address[localhost:9577]
	Address string `json:"address"`
	// key module prefix[optional]
	Module string `json:"module,omitempty"`
}

// ReloadConfig
func ReloadConfig() *Config {
	return &Config{
		Address: "localhost:9577",
	}
}
