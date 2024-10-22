-- +goose Up
CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title text,
    body text
);

-- +goose Down
DROP TABLE posts;