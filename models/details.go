package models

import (
	"context"

	"github.com/SDC-Paprika/go-products/db"
)

type Details struct {
	ID           int    `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Slogan       string `db:"slogan" json:"slogan"`
	Description  string `db:"description" json:"description"`
	Category     string `db:"category" json:"category"`
	DefaultPrice string `db:"default_price" json:"default_price"`
	Features     Feats  `json:"features"`
}

type Feats struct {
	Feature string `db:"feature" json:"feature"`
	Value   string `db:"value" json:"value"`
}

type DetailsModel struct{}

func (d DetailsModel) Get(productId int) (details Details, err error) {
	// query := "SELECT products.*, (SELECT json_agg(f) FROM (SELECT feature, value FROM features WHERE product_id = $1) AS f) AS features FROM products WHERE id = $1"
	query := `
		SELECT
			p.id,
			p.name,
			p.slogan,
			p.description,
			p.category,
			p.default_price,
			(
				SELECT json_agg(f)
				FROM (
					SELECT feature, value FROM features WHERE product_id = $1
				) AS f
			) AS features
		FROM products p
		WHERE id = $1`

	err = db.GetDB().QueryRow(context.Background(), query, productId).Scan(&details)
	return
}
