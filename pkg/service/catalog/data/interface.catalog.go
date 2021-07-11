package data

import "time"

type Item struct {
	Id        uint64    `json:"id" db:"id"`
	SKU       string    `json:"sku" db:"sku"`
	Name      string    `json:"name" db:"name"`
	Price     float64   `json:"price" db:"price"`
	QTY       int       `json:"qty" db:"qty"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CatalogData interface {
	GetItem(name string) (*Item, error)
}
