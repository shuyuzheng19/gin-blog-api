CREATE TABLE categories
(
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    id         int PRIMARY KEY,
    name       VARCHAR                     NOT NULL UNIQUE
);

CREATE SEQUENCE categories_id_seq;

CREATE INDEX idx_categories_name ON categories (name);
CREATE INDEX idx_categories_deleted_at ON categories (deleted_at);

ALTER TABLE categories OWNER TO zsy;

COMMENT
ON COLUMN categories.id IS '分类ID';
COMMENT
ON COLUMN categories.name IS '分类名称';
