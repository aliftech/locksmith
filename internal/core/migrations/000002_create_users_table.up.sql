-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL
)