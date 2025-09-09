-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE wallets (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    ticker VARCHAR(30) NOT NULL,
    mnemonic TEXT NOT NULL,
    public_key VARCHAR(255) NOT NULL,
    private_key TEXT NOT NULL,
    address VARCHAR(255) NOT NULL UNIQUE,
    derivation_index INT NOT NULL DEFAULT 0,
    passphrase_hash VARCHAR(64) NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL DEFAULT CURRENT_TIMESTAMP
);