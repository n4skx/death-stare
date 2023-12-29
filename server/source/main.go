package main

import (
	"github.com/charmbracelet/log"
	"os"
	"server/source/cmd"
	"time"
)

var (
	err    error
	config *cmd.Config

	// Logger
	logger = log.NewWithOptions(os.Stdout, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
)

func init() {
	logger.Info("Parsing configuration")

	// Parse
	config, err = cmd.ParseConfig("config.yaml")

	if err != nil {
		logger.Fatalf("Couldn't parse configuration config.yaml, err: %s", err)
	}
}

func main() {
	for _, listener := range config.Listeners {
		logger.Infof("Listener %s started", listener.Name)
	}
}
