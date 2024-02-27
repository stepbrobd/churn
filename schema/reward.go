package schema

type Reward struct {
	ID        string  `db:"id,key"`     // UUIDv7
	Category  string  `db:"category"`   // Reward category: "dining", "travel", "grocery", "fuel", "other"
	Unit      string  `db:"unit"`       // Reward unit: "point", "dollar"
	Reward    float32 `db:"reward"`     // Reward amount per dollar spent
	ProductID string  `db:"product_id"` // Foreign key to product.id
}
