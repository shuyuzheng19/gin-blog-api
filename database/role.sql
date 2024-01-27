CREATE TABLE roles
(
    id          int PRIMARY KEY,
    name        VARCHAR NOT NULL UNIQUE,
    description VARCHAR NOT NULL
);

CREATE INDEX idx_roles_name ON roles (name);

COMMENT
ON COLUMN roles.id IS '角色ID';
COMMENT
ON COLUMN roles.name IS '角色名';
COMMENT
ON COLUMN roles.description IS '角色描述';