package models

import (
	db "github.com/Ademayowa/k8s-api-deployment/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Property struct {
	ID          string   `json:"id"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Type        string   `json:"type" binding:"required"`
	Status      string   `json:"status" binding:"required"`
	Price       int64    `json:"price" binding:"required"`
	Bedrooms    int      `json:"bedrooms" binding:"required"`
	Bathrooms   int      `json:"bathrooms" binding:"required"`
	SizeSqm     int      `json:"size_sqm" binding:"required"`
	Address     string   `json:"address" binding:"required"`
	Images      []string `json:"images"`
	CreatedAt   string   `json:"created_at,omitempty"`
}

// Save Property into the database
func (p *Property) Save() error {
	p.ID = uuid.New().String()

	query := `
		INSERT INTO properties (id, title, description, type, status, price, bedrooms, bathrooms, size_sqm, address, images)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := db.DB.Exec(
		query,
		p.ID,
		p.Title,
		p.Description,
		p.Type,
		p.Status,
		p.Price,
		p.Bedrooms,
		p.Bathrooms,
		p.SizeSqm,
		p.Address,
		pq.Array(p.Images),
	)

	return err
}

// Retrieves all properties from the database
func GetAllProperties() ([]Property, error) {
	query := `
		SELECT id, title, description, type, status, price, bedrooms, bathrooms, size_sqm, address, images, created_at
		FROM properties
		ORDER BY created_at DESC
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []Property

	for rows.Next() {
		var property Property
		err := rows.Scan(
			&property.ID,
			&property.Title,
			&property.Description,
			&property.Type,
			&property.Status,
			&property.Price,
			&property.Bedrooms,
			&property.Bathrooms,
			&property.SizeSqm,
			&property.Address,
			pq.Array(&property.Images),
			&property.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		properties = append(properties, property)
	}

	return properties, rows.Err()
}

// Retrieves a single property by ID from the database
func GetPropertyByID(id string) (*Property, error) {
	query := `
		SELECT id, title, description, type, status, price, bedrooms, bathrooms, size_sqm, address, images, created_at
		FROM properties
		WHERE id = $1
	`

	var property Property
	err := db.DB.QueryRow(query, id).Scan(
		&property.ID,
		&property.Title,
		&property.Description,
		&property.Type,
		&property.Status,
		&property.Price,
		&property.Bedrooms,
		&property.Bathrooms,
		&property.SizeSqm,
		&property.Address,
		pq.Array(&property.Images),
		&property.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &property, nil
}
