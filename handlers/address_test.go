package handlers

import (
	"address/database"
	"address/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/go-playground/assert/v2"
)

func TestCreateValidAddress(t *testing.T) {
	database.InitDB()

	router := gin.Default()
	router.POST("/addresses", CreateAddress)

	testAddress := models.Address{
		Name:      "TEST",
		Longitude: 100,
		Latitude:  50,
	}

	jsonValue, err := json.Marshal(testAddress)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/addresses", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	Equal(t, http.StatusCreated, w.Code)
}

func TestCreateAddressWithMissingName(t *testing.T) {
	database.InitDB()

	router := gin.Default()
	router.POST("/addresses", CreateAddress)

	testAddress := models.Address{
		Name:      "",
		Longitude: 100,
		Latitude:  50,
	}

	jsonValue, err := json.Marshal(testAddress)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/addresses", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	NotEqual(t, http.StatusCreated, w.Code)
}

func TestCreateAddressWithMissingLongitude(t *testing.T) {
	database.InitDB()

	router := gin.Default()
	router.POST("/addresses", CreateAddress)

	testAddress := models.Address{
		Name:     "Test",
		Latitude: 50,
	}

	jsonValue, err := json.Marshal(testAddress)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/addresses", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	NotEqual(t, http.StatusCreated, w.Code)
}

func TestCreateAddressWithMissingLatitude(t *testing.T) {
	database.InitDB()

	router := gin.Default()
	router.POST("/addresses", CreateAddress)

	testAddress := models.Address{
		Name:      "Test",
		Longitude: 100,
	}

	jsonValue, err := json.Marshal(testAddress)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/addresses", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	NotEqual(t, http.StatusCreated, w.Code)
}

func TestCreateAddressWithWrongLatitudeValue(t *testing.T) {
	database.InitDB()

	router := gin.Default()
	router.POST("/addresses", CreateAddress)

	testAddress := models.Address{
		Name:      "Test",
		Longitude: 100,
		Latitude:  91, // Should be between -90 and 90
	}

	jsonValue, err := json.Marshal(testAddress)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/addresses", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	NotEqual(t, http.StatusCreated, w.Code)
}

func TestCreateAddressWithWrongLongitudeValue(t *testing.T) {
	database.InitDB()

	router := gin.Default()
	router.POST("/addresses", CreateAddress)

	testAddress := models.Address{
		Name:      "Test",
		Longitude: 181, // Should be between -180 and 180
		Latitude:  91,
	}

	jsonValue, err := json.Marshal(testAddress)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/addresses", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	NotEqual(t, http.StatusCreated, w.Code)
}

func TestGetAddresses(t *testing.T) {
	database.InitDB()

	router := gin.Default()
	router.GET("/addresses", GetAddresses)

	req, err := http.NewRequest("GET", "/addresses", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	Equal(t, http.StatusOK, w.Code)

	var responseBody struct {
		Addresses []string `json:"addresses"`
	}

	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	NotEqual(t, 0, len(responseBody.Addresses))

}

func TestGetAddressByID(t *testing.T) {
	database.InitDB()

	router := gin.Default()
	router.GET("/addresses/:id", GetAddressByID)

	req, err := http.NewRequest("GET", "/addresses/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	Equal(t, http.StatusOK, w.Code)
}

// func TestUpdateValidAddress(t *testing.T) {
// 	database.InitDB()

// 	router := gin.Default()
// 	router.PUT("/addresses/:id", UpdateAddress)

// 	updatedAddress := models.Address{
// 		Name:         "Updated Test",
// 		Longitude:    40.7128,
// 		Latitude:     -74.0060,
// 		Floor:        "1",
// 		LocationType: "Office",
// 		Yard:         "Front",
// 	}

// 	updatedAddressJSON, err := json.Marshal(updatedAddress)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req, err := http.NewRequest("PUT", "/addresses/2", bytes.NewBuffer(updatedAddressJSON))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	Equal(t, http.StatusOK, w.Code)

// }
