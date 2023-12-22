package utils

import (
	"address/models"
)

func IsValidFloorType(floor models.FloorType) bool {
	switch floor {
	case models.Basement, models.Sidewalk, models.GroundFloor, models.First, models.Second, models.Third, models.Fourth, models.Fifth, models.Sixth, models.Seventh:
		return true
	default:
		return false
	}
}

func IsValidLocationType(locationType models.LocationType) bool {
	switch locationType {
	case models.Individual, models.Company, models.RetailStore, models.Event, models.Supermarket, models.Warehouse, models.DistributionPlatform:
		return true
	default:
		return false
	}
}

func IsValidYardType(yard models.YardType) bool {
	switch yard {
	case models.None, models.Inf_10m, models.Between_10_30m, models.Sup_30m:
		return true
	default:
		return false
	}
}

func IsValidCoordinates(latitude, longitude float64) bool {
	return latitude >= -90 && latitude <= 90 && longitude >= -180 && longitude <= 180
}
