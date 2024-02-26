package schema

import "time"

type Bonus struct {
	ID      string    // UUIDv7
	Type    string    // Bonus category: "signup", "retention", "spend"
	Spend   int       // Minimum spend to get the bonus
	Bonus   int       // Bonus amount
	Unit    string    // Bonus unit: "point", "dollar"
	Start   time.Time // Bonus start date
	End     time.Time // Bonus end date
	Account string    // Foreign key to Account.ID
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
