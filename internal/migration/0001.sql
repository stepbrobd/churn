CREATE TABLE IF NOT EXISTS bank
(
    bank_alias TEXT UNIQUE PRIMARY KEY,
    bank_name TEXT NOT NULL,
    max_account INT,
    max_account_period INT,
    max_account_scope TEXT
);

CREATE TABLE IF NOT EXISTS product
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_alias TEXT NOT NULL,
    product_name TEXT NOT NULL,
    fee DECIMAL NOT NULL,
    issuing_bank TEXT NOT NULL,
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
    cl DECIMAL,
    FOREIGN KEY (product_id) REFERENCES product (
        id
    ) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS bonus
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    bonus_type TEXT NOT NULL,
    spend DECIMAL NOT NULL,
    bonus_amount DECIMAL NOT NULL,
    unit TEXT NOT NULL,
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
    category TEXT NOT NULL,
    unit TEXT NOT NULL,
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
    category TEXT NOT NULL,
    note TEXT,
    account_id INT NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account (
        id
    ) ON DELETE CASCADE ON UPDATE CASCADE
);
