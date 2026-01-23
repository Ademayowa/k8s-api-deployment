// Test for retrieving all properties - Test GET /properties endpoint

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

func TestGetProperties(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	db.DB = mockDB

	rows := sqlmock.NewRows([]string{"id", "title", "description", "type", "status", "price", "bedrooms", "bathrooms", "size_sqm", "address", "images", "created_at"})
	mock.ExpectQuery("SELECT (.+) FROM properties").WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/properties", nil)
	w := httptest.NewRecorder()

	router := gin.Default()
	router.GET("/properties", handlers.GetProperties)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
