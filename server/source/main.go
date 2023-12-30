package main

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"server/source/cmd"
	"server/source/routes/api"
	"server/source/types"
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
	// Database
	db, err := gorm.Open(sqlite.Open("store/data.db"))

	if err != nil {
		log.Fatalf("Couldn't open database, err: %s\n", err)
	}

	// Migrate
	err = db.AutoMigrate(types.Agent{}, types.Listener{}, types.Operator{})

	if err != nil {
		log.Fatalf("Couldn't migrate database, err: %s\n", err)
	}

	// Webserver
	router := gin.Default()

	apiRoute := router.Group("/api/v1")
	{
		if err = api.SetupAPI(db, apiRoute); err != nil {
			log.Fatalf("Couldn't setup api route, err: %s\n", err)
		}
	}

	log.Infof("Listening at %s", config.Server.Bind)

	router.Run(config.Server.Bind)
}
