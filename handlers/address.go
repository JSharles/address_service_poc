package handlers

import (
	"address/database"
	"address/models"
	u "address/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateAddress godoc
// @Summary Create a new address
// @Description CreateAddress is an example controller to create a new address.
// @Tags addresses
// @Accept json
// @Produce json
// @Param address body models.AddressRequest  true "Address information"
// @Success 201 {object} models.Address "Successfully created"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/addresses [post]
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields name, longitude, and latitude are required"})
		return
	}

	if !u.IsValidCoordinates(address.Latitude, address.Longitude) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and longitude values should fall within the specified range."})
		return
	}

	if address.Floor != "" && !u.IsValidFloorType(address.Floor) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'floor' is not valid."})
		return
	}

	if address.LocationType != "" && !u.IsValidLocationType(address.LocationType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'location_type' is not valid."})
		return
	}

	if address.LocationType != "" && !u.IsValidYardType(address.Yard) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'yard' is not valid."})
		return
	}

	newAddress, err := database.InsertAddress(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newAddress)
}

// GetAddresses godoc
// @Summary Get all addresses
// @Description GetAddresses is an example controller that fetches addresses.
// @Tags addresses
// @Produce json
// @Param active query boolean false "Filter by active status"
// @Param location_type query string false "Filter by location type"
// @Success 200 {array} models.Address "Successfully retrieved"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/addresses [get]
func GetAddresses(c *gin.Context) {
	isActive := c.Query("active")
	locationType := c.Query("location_type")

	addresses, err := database.GetAddresses(isActive, locationType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"addresses": addresses})
}

// GetAddressByID godoc
// @Summary Get an address by ID
// @Description GetAddressByID is an example controller that fetches an address by its ID.
// @Tags addresses
// @Produce json
// @Param id path int true "Address ID"
// @Success 200 {object} models.Address "Successfully retrieved"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Address not found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/addresses/{id} [get]
func GetAddressByID(c *gin.Context) {
	addressIDStr := c.Param("id")
	addressID, err := strconv.Atoi(addressIDStr)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
		return
	}

	address, err := database.GetAddressByID(addressID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
			return
		}
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": address})
}

// UpdateAddress godoc
// @Summary Update an address by ID
// @Description UpdateAddress is an example controller that updates an address by its ID.
// @Tags addresses
// @Accept json
// @Produce json
// @Param id path int true "Address ID"
// @Param address body models.Address true "Updated address information"
// @Success 200 {object} string "Address updated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Address not found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/addresses/{id} [put]
func UpdateAddress(c *gin.Context) {
	addressIDStr := c.Param("id")
	addressID, err := strconv.Atoi(addressIDStr)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
		return
	}

	_, err = database.GetAddressByID(addressID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
			return
		}
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var updatedAddress models.Address
	if err := c.ShouldBindJSON(&updatedAddress); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	requiredFields := []struct {
		name  string
		value string
	}{
		{"name", updatedAddress.Name},
		{"longitude", fmt.Sprintf("%f", updatedAddress.Longitude)},
		{"latitude", fmt.Sprintf("%f", updatedAddress.Latitude)},
	}

	for _, field := range requiredFields {
		if field.value == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Field '" + field.name + "' is required."})
			return
		}
	}

	if !u.IsValidCoordinates(updatedAddress.Latitude, updatedAddress.Longitude) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and longitude values should fall within the specified range."})
		return
	}

	if !u.IsValidFloorType(updatedAddress.Floor) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'floor' is not valid."})
		return
	}

	if !u.IsValidLocationType(updatedAddress.LocationType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'location_type' is not valid."})
		return
	}

	if !u.IsValidYardType(updatedAddress.Yard) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'yard' is not valid."})
		return
	}

	err = database.UpdateAddress(addressID, updatedAddress)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update address"})
		return
	}

	updatedAddress, err = database.GetAddressByID(addressID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully", "address": updatedAddress})
}

// DeleteAddress godoc
// @Summary Delete an address by ID
// @Description DeleteAddress is an example controller that deletes an address by its ID.
// @Tags addresses
// @Produce json
// @Param id path int true "Address ID"
// @Success 200 {object} string "Address deleted successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Address not found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /api/v1/addresses/{id} [delete]
func DeleteAddress(c *gin.Context) {
	addressIDStr := c.Param("id")
	addressID, err := strconv.Atoi(addressIDStr)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
		return
	}
	_, err = database.GetAddressByID(addressID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
			return
		}
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	err = database.DeleteAddress(addressID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}
