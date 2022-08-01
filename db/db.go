package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

var db DB

func Init() {
	dbConfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	var err error
	db, err = ConnectDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB(connString string) (DB, error) {
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return DB{}, err
	}

	if err = db.Ping(context.Background()); err != nil {
		return DB{}, err
	}

	return DB{db}, nil
}

func GetDB() DB {
	return db
}

var Queries = map[string]string{
	// type Queries struct {
	// 	ProductsQuery string
	// 	DetailsQuery string
	// 	RelatedQuery string
	// 	StylesQuery string
	// }{
	"productsQuery": `
		SELECT
			*
		FROM
			products
		LIMIT $1 OFFSET $2`,

	"detailsQuery": `
		SELECT
			p.id,
			p.name,
			p.slogan,
			p.description,
			p.category,
			p.default_price,
			(
				SELECT
					json_agg(f)
				FROM (
					SELECT
						feature,
						value
					FROM
						features
					WHERE
						product_id = $1
				) AS f
			) AS features
		FROM
			products p
		WHERE
			id = $1`,

	"relatedQuery": `
		SELECT
			related_product_id
		FROM
			related
		WHERE
			current_product_id = $1`,

	"stylesQuery": `SELECT
		(
			SELECT
				json_agg(style_agg)
			FROM
				(
					SELECT
						styles.id AS style_id,
						styles.name,
						styles.original_price,
						styles.sale_price,
						styles.default AS "default?",
						(
							SELECT
								json_agg(photo_agg)
							FROM
								(
									SELECT
										thumbnail_url,
										url
									FROM
										photos
									WHERE
										photos.style_id = styles.id
								) AS photo_agg
						) AS photos,
						(
							SELECT
								json_object_agg(
									skus.id,
									json_build_object(
										'quantity', skus.quantity,
										'size', skus.size
									)
								) AS skus
							FROM
								skus
							WHERE
								skus.style_id = styles.id
						) AS skus
					FROM
						styles
					WHERE
						product_id = $1
				) AS style_agg
		) AS results`,
}
