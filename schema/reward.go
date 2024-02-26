package schema

type Reward struct {
	ID       string `db:"id,key"`   // UUIDv7
	Category string `db:"category"` // Reward category: "dining", "travel", "grocery", "fuel", "other"
	Unit     string `db:"unit"`     // Reward unit: "point", "dollar"
	Reward   int    `db:"reward"`   // Reward amount per dollar spent
	Product  string `db:"product"`  // Foreign key to Product.ID
}
