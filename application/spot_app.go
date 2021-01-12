package application

import (
	"github.com/wcarlsen/seatrout/domain/entity"
	"github.com/wcarlsen/seatrout/domain/repository"
)

type spotApp struct {
	sp repository.SpotRepository
}

type SpotAppInterface interface {
	SaveSpot(*entity.Spot) (*entity.Spot, map[string]string)
	GetSpots() (*entity.Spots, error)
	GetSpot(uint64) (*entity.Spot, error)
}

func (s *spotApp) SaveSpot(spot *entity.Spot) (*entity.Spot, map[string]string) {
	return s.sp.SaveSpot(spot)
}

func (s *spotApp) GetSpots() (*entity.Spots, error) {
	return s.sp.GetSpots()
}

func (s *spotApp) GetSpot(spotId uint64) (*entity.Spot, error) {
	return s.sp.GetSpot(spotId)
}
