-- +migrate Up
CREATE TABLE "categories" (
  "cate_id" varchar(255) PRIMARY KEY,
  "cate_name" varchar(255) UNIQUE NOT NULL,
  "cate_image" varchar(255),
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);

-- +migrate Down
DROP TABLE categories;