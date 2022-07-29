package models

import (
	"github.com/SDC-Paprika/go-products/db"
)

type Product struct {
	ID           int    `db:"id, primarykey" json:"id"`
	Name         string `db:"name" json:"name"`
	Slogan       string `db:"slogan" json:"slogan"`
	Description  string `db:"description" json:"description"`
	Category     string `db:"category" json"category"`
	DefaultPrice string `db:"default_price" json:"`
}

type ProductsModel struct{}

// FIXME: `products` has a placeholder type
func (p ProductsModel) Get(page, count int) (products []int, err error) {
	offset := count * (page - 1)
	query := "SELECT * FROM products LIMIT $1 OFFSET $2"

	_, err = db.GetDB().Select(&products, query, count, offset)

	return
}
