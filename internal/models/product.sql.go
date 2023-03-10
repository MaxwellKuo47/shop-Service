// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: product.sql

package models

import (
	"context"
	"encoding/json"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO product_list (
  product_name, product_description, product_market_price, product_sale_price, product_tags, pictures, colors
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, create_at, update_at, product_name, product_description, product_market_price, product_sale_price, product_tags, pictures, colors
`

type CreateProductParams struct {
	ProductName        string          `json:"product_name"`
	ProductDescription string          `json:"product_description"`
	ProductMarketPrice int32           `json:"product_market_price"`
	ProductSalePrice   int32           `json:"product_sale_price"`
	ProductTags        json.RawMessage `json:"product_tags"`
	Pictures           json.RawMessage `json:"pictures"`
	Colors             json.RawMessage `json:"colors"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (ProductList, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ProductName,
		arg.ProductDescription,
		arg.ProductMarketPrice,
		arg.ProductSalePrice,
		arg.ProductTags,
		arg.Pictures,
		arg.Colors,
	)
	var i ProductList
	err := row.Scan(
		&i.ID,
		&i.CreateAt,
		&i.UpdateAt,
		&i.ProductName,
		&i.ProductDescription,
		&i.ProductMarketPrice,
		&i.ProductSalePrice,
		&i.ProductTags,
		&i.Pictures,
		&i.Colors,
	)
	return i, err
}

const createProductDetail = `-- name: CreateProductDetail :one
INSERT INTO product_detail (
  product_id, color, size_os, size_s, size_m, size_l, size_xl
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, create_at, update_at, product_id, color, size_os, size_s, size_m, size_l, size_xl
`

type CreateProductDetailParams struct {
	ProductID int64  `json:"product_id"`
	Color     string `json:"color"`
	SizeOs    int32  `json:"size_os"`
	SizeS     int32  `json:"size_s"`
	SizeM     int32  `json:"size_m"`
	SizeL     int32  `json:"size_l"`
	SizeXl    int32  `json:"size_xl"`
}

func (q *Queries) CreateProductDetail(ctx context.Context, arg CreateProductDetailParams) (ProductDetail, error) {
	row := q.db.QueryRowContext(ctx, createProductDetail,
		arg.ProductID,
		arg.Color,
		arg.SizeOs,
		arg.SizeS,
		arg.SizeM,
		arg.SizeL,
		arg.SizeXl,
	)
	var i ProductDetail
	err := row.Scan(
		&i.ID,
		&i.CreateAt,
		&i.UpdateAt,
		&i.ProductID,
		&i.Color,
		&i.SizeOs,
		&i.SizeS,
		&i.SizeM,
		&i.SizeL,
		&i.SizeXl,
	)
	return i, err
}
