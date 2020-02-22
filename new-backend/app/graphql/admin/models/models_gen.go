// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models_gen

import (
	"github.com/99designs/gqlgen/graphql"
)

type PayloadEntity interface {
	IsPayloadEntity()
}

type InputSetProduct struct {
	ID     *string           `json:"id" db:"id"`
	Name   string            `json:"name" db:"name"`
	Desp   string            `json:"desp" db:"desp"`
	Images []*graphql.Upload `json:"images" db:"images"`
}

type Member struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Product struct {
	ID     string          `json:"id" db:"id"`
	Name   string          `json:"name" db:"name"`
	Desp   *string         `json:"desp" db:"desp"`
	Images []*ProductImage `json:"images" db:"images"`
}

type ProductImage struct {
	ID      string   `json:"id" db:"id"`
	No      string   `json:"no" db:"no"`
	Product *Product `json:"product" db:"product"`
	Image   string   `json:"image" db:"image"`
	Seq     int      `json:"seq" db:"seq"`
	IsMain  bool     `json:"is_main" db:"is_main"`
}

type SetProductPayload struct {
	Code int    `json:"code" db:"code"`
	Msg  string `json:"msg" db:"msg"`
	Data string `json:"data" db:"data"`
}

func (SetProductPayload) IsPayloadEntity() {}
