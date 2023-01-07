CREATE TABLE "product_list" (
  "id" bigserial PRIMARY KEY,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT (now()),
  "product_name" varchar NOT NULL,
  "product_description" text NOT NULL,
  "product_market_price" int NOT NULL,
  "product_sale_price" int NOT NULL,
  "product_tags" jsonb NOT NULL,
  "pictures" jsonb NOT NULL,
  "colors" jsonb NOT NULL
);

CREATE TABLE "product_detail" (
  "id" bigserial PRIMARY KEY,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT (now()),
  "product_id" bigint NOT NULL,
  "color" char(7) NOT NULL,
  "size_os" int NOT NULL,
  "size_s" int NOT NULL,
  "size_m" int NOT NULL,
  "size_l" int NOT NULL,
  "size_xl" int NOT NULL
);

CREATE INDEX ON "product_list" ("product_name");

CREATE INDEX ON "product_list" ("product_tags");

CREATE INDEX ON "product_list" ("colors");

CREATE INDEX ON "product_list" ("product_tags", "colors");

CREATE INDEX ON "product_detail" ("product_id");

CREATE INDEX ON "product_detail" ("product_id", "color");

ALTER TABLE "product_detail" ADD FOREIGN KEY ("product_id") REFERENCES "product_list" ("id");
