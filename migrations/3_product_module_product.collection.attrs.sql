-- +migrate Up
CREATE TABLE "products" (
  "product_id" varchar(255) PRIMARY KEY,
  "product_name" varchar(255) UNIQUE NOT NULL,
  "product_image" varchar(255) NOT NULL,
  "cate_id" varchar(255) NOT NULL,
  "product_des" varchar(255) NOT NULL,
  "collection_id" varchar(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);

CREATE TABLE "collections" (
  "collection_id" varchar(255) PRIMARY KEY,
  "collection_name" varchar(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);

CREATE TABLE "attributes" (
  "attr_id" varchar(255) PRIMARY KEY,
  "product_id" varchar(255) NOT NULL,
  "collection_id" varchar(255) NOT NULL,
  "attr_name" varchar(255) UNIQUE NOT NULL,
  "size" int NOT NULL,
  "price" decimal(15,2) NOT NULL,
  "promotion" numeric,
  "quantity" int NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);

-- +migrate Down
DROP TABLE products;
DROP TABLE collections;
DROP TABLE attributes;