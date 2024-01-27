CREATE TABLE users
(
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    id         int PRIMARY KEY,
    username   VARCHAR(16)              NOT NULL UNIQUE,
    password   VARCHAR                  NOT NULL,
    email      VARCHAR                  NOT NULL,
    avatar     VARCHAR DEFAULT '/logo.svg'::text,
    nick_name  VARCHAR(50)              NOT NULL,
    role_id    BIGINT REFERENCES roles (id),
    CONSTRAINT fk_public_users_role FOREIGN KEY (role_id) REFERENCES roles (id)
);

CREATE SEQUENCE users_id_seq;

CREATE INDEX idx_users_username ON users (username);
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_deleted_at ON users (deleted_at);

COMMENT
ON COLUMN users.id IS '用户ID';
COMMENT
ON COLUMN users.username IS '账号';
COMMENT
ON COLUMN users.password IS '密码';
COMMENT
ON COLUMN users.email IS '邮箱';
COMMENT
ON COLUMN users.avatar IS '头像';
COMMENT
ON COLUMN users.nick_name IS '用户名';
COMMENT
ON COLUMN users.role_id IS '角色ID';
