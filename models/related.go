package models

import (
	"context"

	"github.com/SDC-Paprika/go-products/db"
)

type RelatedModel struct{}

func (r RelatedModel) Get(productId int) (related []int, err error) {
	query := "SELECT related_product_id FROM related WHERE current_product_id = $1"

	rows, err := db.GetDB().Query(context.Background(), query, productId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r int
		if err = rows.Scan(&r); err != nil {
			return nil, err
		}
		related = append(related, r)
	}
	return
}
