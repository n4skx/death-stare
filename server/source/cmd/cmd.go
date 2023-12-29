package cmd

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Operator struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Listener struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Bind struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
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
