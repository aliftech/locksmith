-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE wallets (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    ticker VARCHAR(30) NULL,
    mnemonic TEXT NULL,
    public_key VARCHAR(255) NULL,
    private_key TEXT NULL,
    address VARCHAR(255) NULL UNIQUE,
    derivation_index INT NOT NULL DEFAULT 0,
    passphrase_hash VARCHAR(64) NULL,
    created_by VARCHAR(255) NULL,
    created_at DATETIME NULL,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL
);