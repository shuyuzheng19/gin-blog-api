CREATE TABLE file_infos
(
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at    TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at    TIMESTAMP WITH TIME ZONE,
    id            int PRIMARY KEY ,
    old_name      VARCHAR                  NOT NULL,
    new_name      VARCHAR                  NOT NULL,
    user_id       int REFERENCES users (id),
    CONSTRAINT fk_file_infos_user FOREIGN KEY (user_id) REFERENCES users (id),
    suffix        char(4),
    size          BIGINT,
    absolute_path VARCHAR,
    md5           VARCHAR(255)                  NOT NULL,
    CONSTRAINT fk_file_infos_file_md5_info FOREIGN KEY (md5) REFERENCES file_md5_infos (md5),
    is_pub        BOOLEAN DEFAULT false
);

CREATE SEQUENCE file_infos_id_seq;

CREATE INDEX idx_file_infos_old_name ON file_infos (old_name);
CREATE INDEX idx_file_infos_new_name ON file_infos (new_name);
CREATE INDEX idx_file_infos_user_id ON file_infos (user_id);
CREATE INDEX idx_file_infos_suffix ON file_infos (suffix);
CREATE INDEX idx_file_infos_size ON file_infos (size);
CREATE INDEX idx_file_infos_absolute_path ON file_infos (absolute_path);
CREATE INDEX idx_file_infos_md5 ON file_infos (md5);
CREATE INDEX idx_file_infos_is_pub ON file_infos (is_pub);
CREATE INDEX idx_file_infos_deleted_at ON file_infos (deleted_at);
