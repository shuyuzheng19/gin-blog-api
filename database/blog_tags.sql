

CREATE TABLE blogs_tags
(
    blog_id BIGINT NOT NULL,
    tag_id  BIGINT NOT NULL,
    PRIMARY KEY (blog_id, tag_id),
    CONSTRAINT fk_blogs_tags_blog FOREIGN KEY (blog_id) REFERENCES blogs (id),
    CONSTRAINT fk_blogs_tags_tag FOREIGN KEY (tag_id) REFERENCES tags (id)
);

COMMENT
ON COLUMN blogs_tags.blog_id IS '博客ID';
COMMENT
ON COLUMN blogs_tags.tag_id IS '标签ID';
