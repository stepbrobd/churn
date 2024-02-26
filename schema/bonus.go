package schema

import "time"

type Bonus struct {
	ID      string    `db:"id,key"`  // UUIDv7
	Type    string    `db:"type"`    // Bonus category: "signup", "retention", "spend"
	Spend   int       `db:"spend"`   // Minimum spend to get the bonus
	Bonus   int       `db:"bonus"`   // Bonus amount
	Unit    string    `db:"unit"`    // Bonus unit: "point", "dollar"
	Start   time.Time `db:"start"`   // Bonus start date
	End     time.Time `db:"end"`     // Bonus end date
	Account string    `db:"account"` // Foreign key to Account.ID
}

func (b *Bonus) Schema() string {
	return `CREATE TABLE IF NOT EXISTS bonus (
	id VARCHAR(36) PRIMARY KEY,
	type TEXT NOT NULL,
	spend INTEGER NOT NULL,
	bonus INTEGER NOT NULL,
	unit TEXT NOT NULL,
	start DATETIME NOT NULL,
	end DATETIME NOT NULL,
	account VARCHAR(36) NOT NULL,
	FOREIGN KEY (account) REFERENCES account (id)
);`
}
