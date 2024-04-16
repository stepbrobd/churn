CREATE TABLE IF NOT EXISTS bank
(
    bank_alias VARCHAR(64) UNIQUE PRIMARY KEY,
    bank_name VARCHAR(255) NOT NULL,
    max_account INT,
    max_account_period INT,
    max_account_scope VARCHAR(64)
);

CREATE TABLE IF NOT EXISTS product
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_alias VARCHAR(64) UNIQUE NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    fee DECIMAL NOT NULL,
    issuing_bank VARCHAR(64) NOT NULL,
    FOREIGN KEY (issuing_bank) REFERENCES bank (
        bank_alias
    ) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS account
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT NOT NULL,
    opened DATETIME,
    closed DATETIME,
    cl DECIMAL NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product (
        id
    ) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS bonus
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    bonus_type VARCHAR(64) NOT NULL,
    spend DECIMAL NOT NULL,
    bonus_amount DECIMAL NOT NULL,
    unit VARCHAR(64) NOT NULL,
    bonus_start DATETIME NOT NULL,
    bonus_end DATETIME NOT NULL,
    account_id INT NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account (
        id
    ) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS reward
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    category VARCHAR(64) NOT NULL,
    unit VARCHAR(64) NOT NULL,
    reward DECIMAL NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product (
        id
    ) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS tx
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    tx_timestamp DATETIME NOT NULL,
    amount DECIMAL NOT NULL,
    category VARCHAR(64) NOT NULL,
    note VARCHAR(255),
    account_id INT NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account (
        id
    ) ON DELETE CASCADE ON UPDATE CASCADE
);
