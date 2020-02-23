package models

import "time"

type ProductImage struct {
	ID        int        `json:"id" db:"id"`
	ProductID int        `json:"-" db:"product_id"`
	Image     string     `json:"image" db:"image"`
	Seq       string     `json:"seq" db:"seq"`
	IsMain    string     `json:"is_main" db:"is_main"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	Product   *Product   `json:"product" belongs_to:"product"`
}
