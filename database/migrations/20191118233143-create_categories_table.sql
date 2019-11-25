
-- +migrate Up
CREATE TABLE IF NOT EXISTS categories (
    id      INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    name    VARCHAR(255)                NOT NULL,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE categories;
