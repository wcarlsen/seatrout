package entity

type Spot struct {
	ID            uint64  `gorm:"primary_key;auto_increment"`
	Name          string  `gorm:"size:100;not null;"`
	ParkingAdress string  `gorm:"size:100;unique;not null;"`
	Latitude      float32 `gorm:"not null;unique;"`
	Longitude     float32 `gorm:"not null;unique;"`
	Direction     string  `gorm:"not null;"`
	// CreatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	// UpdatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	// DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}

type Spots []Spot
