#import "./doc.typ": *

#show: doc.with(
  name: "Yifei Sun",
  id: "002245723",
  email: "ysun@ccs.neu.edu",
  institution: "Northeastern University",
  semester: "Spring 2024",
  course: "CS 5200",
  instructor: "Kathleen Durant",
  title: "Churn - Project Report",
  due: datetime(year: 2024, month: 4, day: 18),
)

= Group Information

Group name: SunY

Group member: Yifei Sun

= Technical Specification

- *Language*: Go (approved by professor)
- *Database*: MySQL or SQLite3 (depending on configuration, PostgreSQL "was" supported but abandoned)
- *Frontend*: Command Line Interface
- *Libraries*: mostly handles database connections, drivers, and commandline argument parsing, table display, etc.

= README File

Please see the `readme.md` file located in the root directory of the project for more information.

The README file covers the compilation and installation of the project, as well as the configuration options available to the users.

= Conceptual Design

The toll includes 6 tables: `bank`, `product`, `account`, `bonus`, `reward`, and `tx`.

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

= Logical Design

#image("../examples/churn.png", height: 450pt)

The image is located in the `examples` directory of the project.

= User Flow

Top level commands:

```shell
$ churn -h
Available Commands:
  account     Manage accounts (add, delete, edit)
  bank        Manage banks (add, delete, edit)
  bonus       Manage bonuses (add, delete, edit)
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  migration   Manage database migrations
  product     Manage product (add, delete, edit)
  reward      Manage rewards (add, delete, edit)
  stat        Show statistics
  tx          Manage transactions (add, delete, edit)
```

For `account`, `bonus`, `product`, `reward`, and `tx`, the subcommands are `add`, `delete`, and `edit`.
They are used to add, delete, and edit the corresponding tuples in the database with an interactive CLI.
The top level commands for each subcommands, they are meant to be used to interactively display (list) all the tuples in the database.

Example: `churn account` will list all the accounts in the database, and `churn account add` will add a new account to the database.

For `bank` and `product`, the subcommands are `add`, `delete`, `edit`, and `import`.
The `import` subcommand is used to import data from a json record (either local, or from a URL) to the database.

Example: `churn bank import https://churn.cards/bank.json` will import the bank data from the URL to the database.
`churn bank import ./static/bank.json` will import the bank data from the local file `./static/bank.json` to the database.

The `completion` command is used to generate the autocompletion script for the shells. The supported shells are `bash`, `zsh`, `fish`, and `powershell`.
The completion script can be sourced with `source <(churn completion bash)` for bash, and similar for other shells.


The `migration` command is used to manage the database migrations (currently only supports SQLite3).

The `stat` command is used to show statistics of the database, including the number of tuples in each table, the total amount of transactions, etc.

The `help` command is used to show help for any command.

```txt
   +-------+
+--+ Start |
|  +---+---+---+--------------+------------+
|      |       ^              ^            ^
|      v       | (quit)       |            |
|  +---+----+--+              |            |
|  | Action |                 | (start transaction, rollback if cannot commit)
|  +---+----+---+-------------+------------+----------------+
|      |        |             |            |                |
|      v        v             v            |                v
|  +---+-+      +--------+    +--------+   |         +------------------+
|  | Add |      | Update |    | Delete |   |         | Show (top level) |
|  +---+-+      +----+---+    +----+---+   |         +------------------+
|      |             |             |       |         v
|      v             v             v       | query db, application restrictions, rewards/
|  +---------------------------------------+ bonuses earned, tx history, etc.
|  | Account (instantiation of Product)    | -> add, delete, edit, list
|  | Bank (issuer of Product)              | -> add, delete, edit, list, import
|  | Bonus (associated with Account)       | -> add, delete, edit, list
|  | Product (card product)                | -> add, delete, edit, list, import
|  | Reward (associated with Product)      | -> add, delete, edit, list
|  | Transaction (associated with Account) | -> add, delete, edit, list
|  +---------------------------------------+
|
|
|
+-------------------+
| Completion Script | -> bash | zsh | fish | powershell
| Migration         | -> manage database migrations (SQLite3)
| Stat              | -> show statistics: account, bonus, reward, tx
| Help              | -> show help for any command
+-------------------+
```

Note: Go forces developers to handle errors explicitly,
thus, almost all errors are handled and displayed to the user.

= Lessons Learned

== Technical Expertise Gained

- DBMS for different databases (MySQL, SQLite3)
- `COALESCE` function in SQL

== Insights

- The project is a good practice for handling (multi-) database connections, transactions, and error handling
- Doing the whole project as a one person team is very exhausting, I've got to say I underestimated the amount of work needed to be done for a CLI app

== Alternative Designs

- Being 100% honest, if not required, I would not choose MySQL at all for this project, As a single user CLI app, SQLite3 is more than enough
to handle the data, and it would be possible to replicate the SQLite3 DB file as a user-registrable file, which is more user-friendly than MySQL (concurrecy featuers are not needed for this project)
- This project reaffirms that Go is very tidious to work with, especially when it comes to error handling

== What's Not Working

- In `cmd/stat.go`, I was planning to add some queries with complex joins and selection clauses, but I was soo exhausted that I just gave up on it (see the commented out code in the file (`statRewardCmd`))

= Future Work

- Pick up the commented out code in `cmd/stat.go` and add more complex queries to show more statistics
- Implement MySQL migration support instead of forcing users to import the schemas manually
- Add PostgreSQL support back to the project
- Add import functionality for all six tables
- Add a dump functionality to dump the database to csv or json file that's importable to other financial software
- Add Plaid API support to automatically import transactions from the user's bank account
