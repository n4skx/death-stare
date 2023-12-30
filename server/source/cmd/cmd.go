package cmd

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Core struct {
	Bind string `yaml:"bind"`
}

type Operator struct {
	IsAdmin  bool   `yaml:"is_admin"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Listener struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Bind string `yaml:"bind"`
	Cert struct {
		FullChain  string `yaml:"full_chain"`
		PrivateKey string `yaml:"private_key"`
	}
	ReverseProxy struct {
		Enabled bool `yaml:"enabled"`
		Headers struct {
			RealIP string `yaml:"real_ip"`
		}
	}
}

type Config struct {
	Server    Core       `yaml:"server"`
	Operators []Operator `yaml:"operators"`
	Listeners []Listener `yaml:"listeners"`
}

func ParseConfig(path string) (*Config, error) {
	var config *Config

	// Read
	content, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	// Parse
	err = yaml.Unmarshal(content, &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
