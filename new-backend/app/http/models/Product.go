package models

import "time"

type Product struct {
	ID            int            `json:"id" db:"id"`
	No            string         `json:"no" db:"no"`
	Name          string         `json:"name" db:"name"`
	Brief         *string        `json:"brief" db:"brief"`
	Desp          *string        `json:"desp" db:"desp"`
	Qty           int            `json:"qty" db:"qty"`
	Price         float32        `json:"price" db:"price"`
	CreatedAt     time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time     `json:"deleted_at" db:"deleted_at"`
	ProductImages []ProductImage `json:"product_images" has_many:"product_images"`
}
