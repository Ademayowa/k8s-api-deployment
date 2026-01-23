package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/Ademayowa/k8s-api-deployment/internal/database"
	"github.com/Ademayowa/k8s-api-deployment/internal/handlers"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test for retrieving a single property - Test GET /properties/:id endpoint
func TestGetProperty(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db.DB = mockDB

	rows := sqlmock.NewRows([]string{"id", "title", "description", "type", "status", "price", "bedrooms", "bathrooms", "size_sqm", "address", "images", "created_at"}).
		AddRow("123", "Test", "Desc", "apartment", "for_sale", 50000000, 2, 1, 80, "Address", "{}", "2026-01-23")
	mock.ExpectQuery("SELECT (.+) FROM properties WHERE id").WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/properties/123", nil)
	w := httptest.NewRecorder()

	router := gin.Default()
	router.GET("/properties/:id", handlers.GetProperty)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
