package controllers

import (
	m "address/models"
	u "address/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCreateAddress(t *testing.T) {
	// gin.SetMode(gin.ReleaseMode)
	r := u.SetupTestRouter()
	testAddress := m.Address{
		Name:      "TEST",
		Longitude: 100,
		Latitude:  50,
	}
	jsonValue, _ := json.Marshal(testAddress)
	fmt.Println(testAddress)
	req, _ := http.NewRequest("POST", "/api/addresses/", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

}
