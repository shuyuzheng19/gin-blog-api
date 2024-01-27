CREATE TABLE topics
(
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at  TIMESTAMP WITH TIME ZONE,
    id          int PRIMARY KEY,
    name        VARCHAR                  NOT NULL UNIQUE,
    description VARCHAR                  NOT NULL,
    cover_image VARCHAR                  NOT NULL,
    user_id     int REFERENCES users (id),
    CONSTRAINT fk_topics_user FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE SEQUENCE topics_id_seq;

CREATE INDEX idx_topics_name ON topics (name);
CREATE INDEX idx_topics_deleted_at ON topics (deleted_at);

COMMENT
ON COLUMN topics.id IS '专题ID';
COMMENT
ON COLUMN topics.name IS '专题名';
COMMENT
ON COLUMN topics.description IS '专题描述';
COMMENT
ON COLUMN topics.cover_image IS '专题封面';
COMMENT
ON COLUMN topics.user_id IS '创建专题的id';
