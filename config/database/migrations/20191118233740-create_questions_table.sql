
-- +migrate Up
CREATE TABLE IF NOT EXISTS questions (
    id            INTEGER UNSIGNED AUTO_INCREMENT NOT NULL,
    title         VARCHAR(255) NOT NULL,
    body          TEXT NOT NULL,
    category_id   INTEGER UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY question_category(category_id) REFERENCES categories(id)
);

-- +migrate Down
DROP TABLE questions;
