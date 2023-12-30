package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Migrate - Migrate all route necessary types.
func Migrate(db *gorm.DB) error {
	return nil
}

func SetupAPI(db *gorm.DB, group *gin.RouterGroup) error {
	if err := Migrate(db); err != nil {
		return err
	}

	return nil
}
