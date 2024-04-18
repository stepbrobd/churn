# Churn

[![Built with Nix](https://builtwithnix.org/badge.svg)](https://builtwithnix.org)

## Technical Specifications

- **Language**: Go
- **Database**: MySQL or SQLite3 (depending on configuration, PostgreSQL "was" supported but abandoned)

## Installation

Minimal requirements:

- [Go 1.21.6](https://go.dev/dl): you can use the installers listed in the "Featured downloads" section to download the installer for your platform for easier setup
- Libraries specified in `go.mod`: links to the libraries are available in the `go.mod` file

Get the source code:

```shell
git clone https://github.com/stepbrobd/churn.git && cd churn
```

Build the project:

```shell
go build
```

The binary will be available at `./churn`.

To temporarily add `churn` to your `PATH`, run:

```shell
export PATH=$PATH:$(pwd)
```

For better experience, add the shell completion script (pick one depending on you shell):

```shell
source <(churn completion bash)
# or
source <(churn completion fish)
# or
source <(churn completion zsh)
```

After this, you can run `churn` and press `TAB` to see the available commands.

For example:

```shell
$ churn <TAB>
account     -- Manage accounts (add, delete, edit)
bank        -- Manage banks (add, delete, edit)
bonus       -- Manage bonuses (add, delete, edit)
completion  -- Generate the autocompletion script for the specified shell
help        -- Help about any command
migration   -- Manage database migrations
product     -- Manage product (add, delete, edit)
reward      -- Manage rewards (add, delete, edit)
stat        -- Show statistics
tx          -- Manage transactions (add, delete, edit)
```

## Conceptual Design

```txt
+---------+                        +---------+                +--------------------+
| Account | is a user account of > | Product | is issued by > | Bank               |
+---------+------------------------+---------+----------------+--------------------+
| id {PK} | 0..*            1..1   | id {PK} | 0..*    1..1   | id {PK}            |
| name    |                        | name    |                | name               |
| product |                        | fee     |                | max_acconut        |
| opened  |                        +--+------+                | max_acconut_period |
| closed  |                           |                       +--------------------+
| CL      |                           |
+----+----+------------+ 1..1         |
     | 1..1            |              | 1..1
     |                 |              |
     | have bonuses v  |              | have rewards v
     |                 |              |
     | 0..*            |              | 0..*
+----+----+            |         +----+-----+
| Bonus   |            |         | Reward   |
+---------+            | make v  +----------+
| id {PK} |            |         | id {PK}  |
| type    |            |         | category |
| spend   |            |         | unit     |
| bonus   |            |         | reward   |
| unit    |            |         +----------+
| start   |            |
| end     |            |
+---------+            | 0..*
                       +-------------+
                       | tx          |
                       +-------------+
                       | id {PK}     |
                       | timestamp   |
                       | amount      |
                       | category    |
                       | description |
                       +-------------+
```

Banks have products, products instantiate as accounts (multiple same products can be issued to the same user),
accounts have bonuses (varies between users), and products have rewards (same for all users).
Transactions are made by accounts and are used to track the user's spending and reward categories.

## Logical Design

```mermaid
classDiagram
direction BT
class account {
   int product_id
   datetime opened
   datetime closed
   decimal(10) cl
   int id
}
class bank {
   varchar(255) bank_name
   int max_account
   int max_account_period
   varchar(64) max_account_scope
   varchar(64) bank_alias
}
class bonus {
   varchar(64) bonus_type
   decimal(10) spend
   decimal(10) bonus_amount
   varchar(64) unit
   datetime bonus_start
   datetime bonus_end
   int account_id
   int id
}
class product {
   varchar(64) product_alias
   varchar(255) product_name
   decimal(10) fee
   varchar(64) issuing_bank
   int id
}
class reward {
   varchar(64) category
   varchar(64) unit
   decimal(10) reward
   int product_id
   int id
}
class tx {
   datetime tx_timestamp
   decimal(10) amount
   varchar(64) category
   text note
   int account_id
   int id
}

account  -->  product : product_id:id
bonus  -->  account : account_id:id
product  -->  bank : issuing_bank:bank_alias
reward  -->  product : product_id:id
tx  -->  account : account_id:id
```

Exported from DataGrip's ERD (export to Mermaid).

## User Flow

0. Configuration:
    - Database type (MySQL or SQLite3)
    - Database connection string (DSN)
    - Migration (use the built in migration tool if SQLite3 is used, use the example dump for MySQL)
1. Add a bank:
   - Bulk import for banks are available
   - `churn bank import <a local json or a remote json>`
2. Add a product:
   - Bulk import for products are available
   - `churn product import <a local json or a remote json>`
3. Add an account
4. Add a bonus (optional)
5. Add a reward (optional)
6. Add a transaction (optional)
7. Display statistics (optional)
8. Migrations (optional, not fully implemented, future work)

## License

The contents inside this repository, excluding all submodules, are licensed under the [MIT License](license.txt).
Third-party file(s) and/or code(s) are subject to their original term(s) and/or license(s).
