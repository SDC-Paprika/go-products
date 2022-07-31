package models

import (
	"context"

	"github.com/SDC-Paprika/go-products/db"
)

type Product struct {
	ID           int    `db:"id, primarykey" json:"id"`
	Name         string `db:"name" json:"name"`
	Slogan       string `db:"slogan" json:"slogan"`
	Description  string `db:"description" json:"description"`
	Category     string `db:"category" json:"category"`
	DefaultPrice string `db:"default_price" json:"default_price"`
}

type ProductsModel struct{}

func (p ProductsModel) Get(page, count int) (products []Product, err error) {
	offset := count * (page - 1)
	query := "SELECT * FROM products LIMIT $1 OFFSET $2"

	rows, err := db.GetDB().Query(context.Background(), query, count, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p Product
		if err = rows.Scan(&p); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return
}
