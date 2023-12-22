package controllers

import (
	"address/models"
	u "address/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context) {

	var address models.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		fmt.Println("Error while binding JSON:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requiredFields := []struct {
		name  string
		value string
	}{
		{"name", address.Name},
		{"longitude", fmt.Sprintf("%f", address.Longitude)},
		{"latitude", fmt.Sprintf("%f", address.Latitude)},
	}

	for _, field := range requiredFields {
		if field.value == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Field '" + field.name + "' is required."})
			return
		}
	}

	if address.Name == "" || address.Longitude == 0 || address.Latitude == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields name, longitude and latitude) are required"})
		return
	}

	if !u.IsValidCoordinates(address.Latitude, address.Longitude) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and longitude values should fall within the specified range."})
		return
	}

	if !u.IsValidFloorType(address.Floor) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'floor' is not valid."})
		return
	}

	if !u.IsValidLocationType(address.LocationType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'location_type' is not valid."})
		return
	}

	if !u.IsValidYardType(address.Yard) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'yard' is not valid."})
		return
	}

	if err := u.DB.QueryRow(`
		INSERT INTO addresses (name, longitude, latitude, active, created_at, updated_at, time_zone, complementary_informations, floor, lift, location_type, yard, door_code, loading_dock, side_loading)
		VALUES ($1, $2, $3, $4, NOW(), NOW(), $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id`,
		address.Name, address.Longitude, address.Latitude, address.Active, address.TimeZone, address.ComplementaryInfo,
		address.Floor, address.Lift, address.LocationType, address.Yard, address.DoorCode,
		address.LoadingDock, address.SideLoading).Scan(&address.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, address)
}

func GetAddresses(c *gin.Context) {

}

func GetAddressByID(c *gin.Context) {

}

func UpdateAddress(c *gin.Context) {

}

func DeleteAddress(c *gin.Context) {

}
