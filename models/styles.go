package models

import (
	"context"

	"github.com/SDC-Paprika/go-products/db"
)

type Styles struct {
	ID  int     `db:"id" json:"product_id"`
	Res Results `json:"results"`
}

type Results struct {
	ID            int             `db:"id" json:"styles_id"`
	Name          string          `db:"name" json:"name"`
	OriginalPrice string          `db:"original_price" json:"original_price"`
	SalePrice     string          `db:"sale_price" json:"sale_price"`
	Default       bool            `db:"default" json:"default?"`
	Photos        []Images        `json:"photos"`
	Skus          map[string]SKUs `json:"skus"`
}

type Images struct {
	Thumbnail string `db:"thumbnail_url" json:"thumbnail_url"`
	URL       string `db:"url" json:"url"`
}

type SKUs struct {
	Quantity int    `db:"quantity" json:"quantity"`
	Size     string `db:"size" json:"size"`
}

type StylesModel struct{}

func (s StylesModel) Get(productId int) (styles Styles, err error) {
	query := db.Queries["stylesQuery"]

	var results Results
	err = db.GetDB().QueryRow(context.Background(), query, productId).Scan(&results)
	if err != nil {
		return
	}

	styles = Styles{ID: productId, Res: results}
	return
}
