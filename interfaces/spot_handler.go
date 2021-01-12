package interfaces

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wcarlsen/seatrout/application"
	"github.com/wcarlsen/seatrout/domain/entity"
)

type Spots struct {
	sp application.SpotAppInterface
}

func NewSpots(sp application.SpotAppInterface) *Spots {
	return &Spots{sp: sp}
}

func (s *Spots) SaveSpot(c *gin.Context) {
	var spot entity.Spot
	if err := c.ShouldBindJSON(&spot); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	newSpot, err := s.sp.SaveSpot(&spot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newSpot)
}

func (s *Spots) GetSpots(c *gin.Context) {
	spots := &entity.Spots{}
	var err error
	spots, err = s.sp.GetSpots()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, spots)
}

func (s *Spots) GetSpot(c *gin.Context) {
	spotID, err := strconv.ParseUint(c.Param("spot_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	spot, err := s.sp.GetSpot(spotID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, spot)
}
