package superadmin

import (
	"github.com/hngprojects/hng_boilerplate_golang_web/internal/models"
	"gorm.io/gorm"
)

func AddToRegion(region *models.Region, db *gorm.DB) error {

	if err := region.CreateRegion(db); err != nil {
		return err
	}

	return nil
}

func AddToTimeZone(timezone *models.Timezone, db *gorm.DB) error {

	if err := timezone.CreateTimeZone(db); err != nil {
		return err
	}

	return nil
}

func AddToLanguage(language *models.Language, db *gorm.DB) error {

	if err := language.CreateLanguage(db); err != nil {
		return err
	}

	return nil
}

func GetRegions(db *gorm.DB) ([]models.Region, error) {

	var region models.Region

	regionData, err := region.GetRegions(db)
	if err != nil {
		return nil, err
	}

	return regionData, nil
}

func GetTimeZones(db *gorm.DB) ([]models.Timezone, error) {

	var timezone models.Timezone

	timezoneData, err := timezone.GetTimeZones(db)
	if err != nil {
		return nil, err
	}

	return timezoneData, nil
}

func GetLanguages(db *gorm.DB) ([]models.Language, error) {

	var language models.Language

	languageData, err := language.GetLanguages(db)
	if err != nil {
		return nil, err
	}

	return languageData, nil
}
