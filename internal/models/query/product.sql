-- name: CreateProduct :one
INSERT INTO product_list (
  product_name, product_description, product_market_price, product_sale_price, product_tags, pictures, colors
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: CreateProductDetail :one
INSERT INTO product_detail (
  product_id, color, size_os, size_s, size_m, size_l, size_xl
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;