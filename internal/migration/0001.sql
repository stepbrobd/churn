-- +migrate Up
CREATE TABLE IF NOT EXISTS bank
(
    id VARCHAR(36) PRIMARY KEY,
    bank_name TEXT NOT NULL,
    max_account INTEGER,
    max_account_period INTEGER
);

CREATE TABLE IF NOT EXISTS product
(
    id VARCHAR(36) PRIMARY KEY,
    product_name TEXT NOT NULL,
    fee DECIMAL NOT NULL,
    issuing_bank VARCHAR(36) NOT NULL,
    FOREIGN KEY (issuing_bank) REFERENCES bank (id)
);

CREATE TABLE IF NOT EXISTS account
(
    id VARCHAR(36) PRIMARY KEY,
    account_alias TEXT NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    opened DATETIME NOT NULL,
    closed DATETIME,
    cl DECIMAL,
    FOREIGN KEY (product_id) REFERENCES product (id)
);


CREATE TABLE IF NOT EXISTS bonus
(
    id VARCHAR(36) PRIMARY KEY,
    bonus_type TEXT NOT NULL,
    spend DECIMAL NOT NULL,
    bonus_amount DECIMAL NOT NULL,
    unit TEXT NOT NULL,
    bonus_start DATETIME NOT NULL,
    bonus_end DATETIME NOT NULL,
    account_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account (id)
);


CREATE TABLE IF NOT EXISTS reward
(
    id VARCHAR(36) PRIMARY KEY,
    category TEXT NOT NULL,
    unit TEXT NOT NULL,
    reward DECIMAL NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product (id)
);

CREATE TABLE IF NOT EXISTS tx
(
    id VARCHAR(36) PRIMARY KEY,
    tx_timestamp DATETIME NOT NULL,
    amount DECIMAL NOT NULL,
    category TEXT NOT NULL,
    note TEXT,
    account_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account (id)
);

-- +migrate Down
DROP TABLE IF EXISTS tx;
DROP TABLE IF EXISTS reward;
DROP TABLE IF EXISTS bonus;
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS bank;
