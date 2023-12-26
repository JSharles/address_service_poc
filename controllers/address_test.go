package controllers

import (
	m "address/models"
	"address/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestCreateValidAddress(t *testing.T) {
	utils.InitDB()

	router := gin.Default()
	router.POST("/addresses", CreateAddress)

	testAddress := m.Address{
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

	assert.Equal(t, http.StatusCreated, w.Code)
}

// func TestCreateAddressWithMissingName(t *testing.T) {
// 	utils.InitDB()

// 	router := gin.Default()
// 	router.POST("/addresses", CreateAddress)

// 	testAddress := m.Address{
// 		Name:      "TEST",
// 		Longitude: 100,
// 		Latitude:  50,
// 	}

// 	jsonValue, err := json.Marshal(testAddress)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req, err := http.NewRequest("POST", "/addresses", bytes.NewBuffer(jsonValue))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// }
