package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/Ademayowa/k8s-api-deployment/internal/database"
	"github.com/Ademayowa/k8s-api-deployment/internal/handlers"
	"github.com/Ademayowa/k8s-api-deployment/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test for creating a new property - Test POST /properties endpoint
func TestCreateProperty(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db.DB = mockDB

	mock.ExpectExec("INSERT INTO properties").WillReturnResult(sqlmock.NewResult(1, 1))

	property := models.Property{
		Title:       "Test Property",
		Description: "Test Description",
		Type:        "apartment",
		Status:      "for_sale",
		Price:       50000000,
		Bedrooms:    2,
		Bathrooms:   1,
		SizeSqm:     80,
		Address:     "Test Address",
		Images:      []string{"test.jpg"},
	}

	body, _ := json.Marshal(property)
	req := httptest.NewRequest(http.MethodPost, "/properties", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/properties", handlers.CreateProperty)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
