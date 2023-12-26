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
// @Tags address
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields name, longitude and latitude are required"})
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

	if err := database.DB.QueryRow(`
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

// GetAddresses godoc
// @Summary Get all addresses
// @Description GetAddresses is an example controller that fetches addresses.
// @Tags address
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
	location_type := c.Query("location_type")

	query := "SELECT * FROM addresses WHERE 1=1"
	args := []interface{}{}

	if isActive != "" {
		query += " AND active = $1"
		activeValue, err := strconv.ParseBool(isActive)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'active' parameter"})
			return
		}
		args = append(args, activeValue)
	}

	if location_type != "" {
		query += " AND location_type = $2"
		args = append(args, location_type)
	}

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	var addresses []models.Address
	for rows.Next() {
		var address models.Address
		err := rows.Scan(
			&address.ID,
			&address.Name,
			&address.Longitude,
			&address.Latitude,
			&address.Active,
			&address.CreatedAt,
			&address.UpdatedAt,
			&address.TimeZone,
			&address.ComplementaryInfo,
			&address.Floor,
			&address.Lift,
			&address.LocationType,
			&address.Yard,
			&address.DoorCode,
			&address.LoadingDock,
			&address.SideLoading,
		)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		addresses = append(addresses, address)
	}

	c.JSON(http.StatusOK, gin.H{"addresses": addresses})
}

// GetAddressByID godoc
// @Summary Get an address by ID
// @Description GetAddressByID is an example controller that fetches an address by its ID.
// @Tags address
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
		c.JSON(400, gin.H{"error": "Invalid address ID"})
		return
	}

	var address models.Address
	err = database.DB.QueryRow("SELECT * FROM addresses WHERE id = $1", addressID).
		Scan(
			&address.ID,
			&address.Name,
			&address.Longitude,
			&address.Latitude,
			&address.Active,
			&address.CreatedAt,
			&address.UpdatedAt,
			&address.TimeZone,
			&address.ComplementaryInfo,
			&address.Floor,
			&address.Lift,
			&address.LocationType,
			&address.Yard,
			&address.DoorCode,
			&address.LoadingDock,
			&address.SideLoading,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows found")
			c.JSON(404, gin.H{"error": "Address not found"})
			return
		}
		log.Println(err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"address": address})

}

// UpdateAddress godoc
// @Summary Update an address by ID
// @Description UpdateAddress is an example controller that updates an address by its ID.
// @Tags address
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

	var existingAddress models.Address
	err = database.DB.QueryRow("SELECT * FROM addresses WHERE id = $1", addressID).
		Scan(
			&existingAddress.ID,
			&existingAddress.Name,
			&existingAddress.Longitude,
			&existingAddress.Latitude,
			&existingAddress.Active,
			&existingAddress.CreatedAt,
			&existingAddress.UpdatedAt,
			&existingAddress.TimeZone,
			&existingAddress.ComplementaryInfo,
			&existingAddress.Floor,
			&existingAddress.Lift,
			&existingAddress.LocationType,
			&existingAddress.Yard,
			&existingAddress.DoorCode,
			&existingAddress.LoadingDock,
			&existingAddress.SideLoading,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows found")
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

	_, err = database.DB.Exec(`
		UPDATE addresses
		SET name=$1, longitude=$2, latitude=$3, active=$4, created_at=$5, updated_at=$6,
			time_zone=$7, complementary_informations=$8, floor=$9, lift=$10,
			location_type=$11, yard=$12, door_code=$13, loading_dock=$14, side_loading=$15
		WHERE id=$16
		RETURNING id, name, longitude, latitude, active, created_at, updated_at,
			time_zone, complementary_informations, floor, lift, location_type, yard,
			door_code, loading_dock, side_loading`,
		updatedAddress.Name, updatedAddress.Longitude, updatedAddress.Latitude,
		updatedAddress.Active, updatedAddress.CreatedAt, updatedAddress.UpdatedAt,
		updatedAddress.TimeZone, updatedAddress.ComplementaryInfo, updatedAddress.Floor,
		updatedAddress.Lift, updatedAddress.LocationType, updatedAddress.Yard,
		updatedAddress.DoorCode, updatedAddress.LoadingDock, updatedAddress.SideLoading,
		addressID,
	)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update address"})
		return
	}

	err = database.DB.QueryRow("SELECT * FROM addresses WHERE id = $1", addressID).
		Scan(
			&updatedAddress.ID,
			&updatedAddress.Name,
			&updatedAddress.Longitude,
			&updatedAddress.Latitude,
			&updatedAddress.Active,
			&updatedAddress.CreatedAt,
			&updatedAddress.UpdatedAt,
			&updatedAddress.TimeZone,
			&updatedAddress.ComplementaryInfo,
			&updatedAddress.Floor,
			&updatedAddress.Lift,
			&updatedAddress.LocationType,
			&updatedAddress.Yard,
			&updatedAddress.DoorCode,
			&updatedAddress.LoadingDock,
			&updatedAddress.SideLoading,
		)
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
// @Tags address
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
		c.JSON(400, gin.H{"error": "Invalid address ID"})
		return
	}

	var existingAddress models.Address
	err = database.DB.QueryRow("SELECT * FROM addresses WHERE id = $1", addressID).
		Scan(
			&existingAddress.ID,
			&existingAddress.Name,
			&existingAddress.Longitude,
			&existingAddress.Latitude,
			&existingAddress.Active,
			&existingAddress.CreatedAt,
			&existingAddress.UpdatedAt,
			&existingAddress.TimeZone,
			&existingAddress.ComplementaryInfo,
			&existingAddress.Floor,
			&existingAddress.Lift,
			&existingAddress.LocationType,
			&existingAddress.Yard,
			&existingAddress.DoorCode,
			&existingAddress.LoadingDock,
			&existingAddress.SideLoading,
		)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"error": "Address not found"})
		return
	}

	_, err = database.DB.Exec("DELETE FROM addresses WHERE id = $1", addressID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Failed to delete address"})
		return
	}

	c.JSON(200, gin.H{"message": "Address deleted successfully"})

}
