package stage

import "gorm.io/gorm"

// Migrate - Migrate all route necessary types.
func Migrate(db *gorm.DB) error {
	return nil
}

func SetupStager(db *gorm.DB) error {
	if err := Migrate(db); err != nil {
		return err
	}

	return nil
}
