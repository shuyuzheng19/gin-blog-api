CREATE TABLE blogs
(
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at  TIMESTAMP WITH TIME ZONE,
    id          BIGSERIAL PRIMARY KEY DEFAULT nextval('blogs_id_seq'::regclass),
    description VARCHAR                     NOT NULL,
    title       VARCHAR                     NOT NULL,
    cover_image TEXT                     NOT NULL,
    source_url  TEXT,
    content     TEXT,
    eye_count   BIGINT                DEFAULT 0,
    like_count  BIGINT                DEFAULT 0,
    category_id int REFERENCES categories (id),
    CONSTRAINT fk_blogs_category FOREIGN KEY (category_id) REFERENCES categories (id),
    user_id     int REFERENCES users (id),
    CONSTRAINT fk_blogs_user FOREIGN KEY (user_id) REFERENCES users (id),
    topic_id    int REFERENCES topics (id),
    CONSTRAINT fk_blogs_topic FOREIGN KEY (topic_id) REFERENCES topics (id)
);

CREATE SEQUENCE blogs_id_seq;

CREATE INDEX idx_blogs_deleted_at ON blogs (deleted_at);

COMMENT
ON COLUMN blogs.id IS '博客ID';
COMMENT
ON COLUMN blogs.description IS '博客描述';
COMMENT
ON COLUMN blogs.title IS '博客标题';
COMMENT
ON COLUMN blogs.cover_image IS '博客封面';
COMMENT
ON COLUMN blogs.source_url IS '博客原文链接';
COMMENT
ON COLUMN blogs.content IS '博客正文';
COMMENT
ON COLUMN blogs.eye_count IS '浏览量';
COMMENT
ON COLUMN blogs.like_count IS '浏览量';
COMMENT
ON COLUMN blogs.category_id IS '分类ID';
COMMENT
ON COLUMN blogs.user_id IS '用户ID';
COMMENT
ON COLUMN blogs.topic_id IS '专题ID';
