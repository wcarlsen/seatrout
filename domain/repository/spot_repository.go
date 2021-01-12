package repository

import "github.com/wcarlsen/seatrout/domain/entity"

type SpotRepository interface {
	SaveSpot(*entity.Spot) (*entity.Spot, map[string]string)
	GetSpot(uint64) (*entity.Spot, error)
	GetSpots() (*entity.Spots, error)
}
