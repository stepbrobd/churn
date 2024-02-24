package schema

type Reward struct {
	ID       string // UUIDv7
	Category string // Reward category: "dining", "travel", "grocery", "fuel", "other"
	Unit     string // Reward unit: "point", "dollar"
	Reward   int    // Reward amount per dollar spent
}

func (r *Reward) Schema() string {
	return `CREATE TABLE IF NOT EXISTS reward (
	id VARCHAR(36) PRIMARY KEY,
	category TEXT NOT NULL,
	unit TEXT NOT NULL,
	reward INTEGER NOT NULL
);`
}
