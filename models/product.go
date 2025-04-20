package models

import "time"

// Product represents a product in our e-commerce platform.
type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Price     float64   `gorm:"type:decimal(10,2)" json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
