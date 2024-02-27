#import "./doc.typ": *

#show: doc.with(
  name: "Yifei Sun",
  id: "002245723",
  email: "ysun@ccs.neu.edu",
  institution: "Northeastern University",
  semester: "Spring 2024",
  course: "CS 5200",
  instructor: "Kathleen Durant",
  title: "Churn - Project Proposal",
  due: datetime(year: 2024, month: 3, day: 1),
)

= Introduction

Credit card churing (or simply churing) is a term commonly used in multiple
well-known communities (e.g.
#link("https://www.reddit.com/r/churning")[r/churing] on Reddit,
#link("https://thepointsguy.com")[The Points Guy],
#link("https://www.doctorofcredit.com")[Doctor of Credit], etc.). Simply put, it
refers to the process of consumers applying for credit cards to take advantage
of sign-up bonuses (or other promotions) to maximize the amount of "cash back"
or "points" they can get. Many fears that churing is illegal or unethical.
However, it is not (well, if you follow the terms and conditions). It is simply
a way to "earn back" some of the money that consumers have spent.

Churning is fun, but complex. It requires a lot of planning and tracking. For
example, consumers need to keep track of the credit cards they have applied for,
the annual fees they need to pay, the minimum spending they need to meet to get
the sign-up bonuses, etc. Thus, to simplify the bookkeeping, this project aims
to develop a command-line tool to help consumers track their credit card churing
activities (there is no offical survey, but personally, I believe a good
percentage of churing enthusiasts are tech-savvy and would love to use a
command-line tool).

The motivation behind this project is my personal interest in credit card
churing and have been annoyed by the lack of a good tool to help me keep track
of my churing activities. I hope this tool can help me and other churing
enthusiasts replace to Google Sheets or other manual entry methods.

= Functionalities

The tool, named `churn`, at minimum, will provide the following functionalities:

- Users can add new credit/charge cards, with cards' name, account opening date,
  annual fee, credit limit (if applicable).
- Users can add new bonuses, with bonuses' type (e.g. sign-up bonus, referral
  bonus, retention offers, regular spending bonus, etc.), the bonus amount and
  unit, the minimum spending requirement, the bonus expiration date, etc.
- Users can add new transactions, with transactions' date, amount, category (e.g.
  dining, travel, grocery, etc.), and the card used.
- Users can derive rewards earned from transactions, either in cash back or
  points, based on the cards' reward rates and the transactions' categories.

Additionally, financial institutions may have certain restrictions on the number
of credit cards a consumer can apply for. For example, Chase has the infamous
5/24 rule, which means that if a consumer has opened 5 or more credit cards in
the past 24 months, they will not be approved for a new Chase credit card. Thus,
the tool will also provide the following functionalities:

- Users can assign custom rules to the tool, such as the 5/24 rule, to help them
  keep track of their eligibility for new credit cards.
- Users can check their eligibility for new credit cards for a certain institution
  based on the rules they have assigned.

Also, basic CRUD operations will be provided for all the data. For example,
users might get a credit limit increase, or a bonus might be extended, or an
acconut might be closed, etc.

= Implementation

The tool will be implemented in Go with MySQL as the database (as described
above, entities like cards, bonuses, transactions, etc. have a clear schema,
thus a relational database is a good fit). It's expected to be a command-line
tool, and it should be able to compile down to a single binary. Some libraries
that will be used include:
- MySQL driver for Go
- Cobra for command-line interface
- Viper for configuration management
- ...
With the correct setup, the program should be able to run on all major platforms
(e.g. Windows, macOS, Linux).

The project structure will be organized as follow:

```txt
- cmd/           # command-line interface related code
  - root.go
  - ...
- doc/           # documentation
  - ...
- internal       # core application logic
  - config/...   # config parsing
  - db/          # session management
    - mysql.go   # MySQL driver
    - sqlite3.go # SQLite3 driver
    - ...
  - migration/   # database migrations
    - 0001.sql   # initial schema
    - *.sql      # other migrations
  - ...
- schema/
  - account.go   # account entity with Go struct and SQL schema
  - bank.go      # bank entity with Go struct and SQL schema
  - bonus.go     # bonus entity with Go struct and SQL schema
  - ...          # other entities
- main.go        # entry point
- go.{mod,sum}   # module files
- readme.md
- license.txt
- ...
```

Although the project will use MySQL as the database, it's possible to extend the
tool to support other databases (e.g. PostgreSQL, SQLite, etc.) in the future,
given the clear separation of the database layer from the core application
logic. The database connections are managed by drivers and DSNs, thus, to extend
support to other databases, we only need to change the driver, DSN, and minor
change in SQL syntax (dialect) if necessary.

= Conceptual Design

The tool will have the following entities:

```txt
+---------+                        +---------+                +--------------------+
| Account | is a user account of > | Product | is issued by > | Bank               |
+---------+------------------------+---------+----------------+--------------------+
| id {PK} | 0..*            1..1   | id {PK} | 0..*    1..1   | id {PK}            |
| name    |                        | name    |                | name               |
| product |                        | issuer  |                | max_acconut        |
| opened  |                        | fee     |                | max_acconut_period |
| closed  |                        +----+----+                +--------------------+
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
| unit    |            |         | product  |
| acconut |            |         +----------+
| start   |            |
| end     |            | 0..*
+---------+            +-------------+
                       | tx          |
                       +-------------+
                       | id {PK}     |
                       | timestamp   |
                       | amount      |
                       | category    |
                       | description |
                       | acconut     |
                       +-------------+
```

In the implementation:
- Cards are representaed as "Account"s and they are instantiation of "Product"s (a
  bank's payment card product offering).
- Card products are issued by banks.
- Acconut related restrictions are enforced by the "Bank" entity.
- Usually each customer will have different bonuses for different cards and they
  usually variy from person to person, thus bonuses are associated with "Account"s.
- Rewards however, are associated with card product offerings, as they are usually
  fixed for each card product.
- Transactions are associated with "Account"s.

= Activity Diagram

```txt

```
