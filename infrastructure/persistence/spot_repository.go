package persistence

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/wcarlsen/seatrout/domain/entity"
)

type SpotRepo struct {
	db *gorm.DB
}

func NewSpotRepository(db *gorm.DB) *SpotRepo {
	return &SpotRepo{db}
}

func (r *SpotRepo) SaveSpot(spot *entity.Spot) (*entity.Spot, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Create(&spot).Error
	if err != nil {
		//If the spot already exists
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["spot_exists"] = "spot already exists"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return spot, nil
}

func (r *SpotRepo) GetSpots() (*entity.Spots, error) {
	var spots entity.Spots
	err := r.db.Find(&spots).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("spots not found")
	}
	return &spots, nil
}

func (r *SpotRepo) GetSpot(id uint64) (*entity.Spot, error) {
	var spot entity.Spot
	err := r.db.Where("id = ?", id).Take(&spot).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("spot not found")
	}
	return &spot, nil
}
