CREATE TABLE tags
(
    id         INT PRIMARY KEY,
    name       VARCHAR                  NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE SEQUENCE tags_id_seq;

CREATE INDEX idx_tags_name ON tags (name);
CREATE INDEX idx_tags_deleted_at ON tags (deleted_at);

COMMENT
ON COLUMN tags.id IS '标签ID';
COMMENT
ON COLUMN tags.name IS '标签名';
