-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    firstname VARCHAR(255) NULL,
    lastname VARCHAR(255) NULL,
    email VARCHAR(255) NULL UNIQUE,
    password TEXT NULL,
    created_at DATETIME NULL,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL
)