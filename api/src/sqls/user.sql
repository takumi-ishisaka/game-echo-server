-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(128) NOT NULL PRIMARY KEY,
    name VARCHAR(128) NOT NULL

);
-- +migrate Down
DROP TABLE IF EXISTS users;
