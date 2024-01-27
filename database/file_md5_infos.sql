CREATE TABLE file_md5_infos
(
    md5           VARCHAR(255) NOT NULL PRIMARY KEY,
    url           TEXT UNIQUE,
    absolute_path TEXT
);


COMMENT
ON COLUMN file_md5_infos.md5 IS '文件md5';
COMMENT
ON COLUMN file_md5_infos.url IS '文件url';
COMMENT
ON COLUMN file_md5_infos.absolute_path IS '文件路径';
