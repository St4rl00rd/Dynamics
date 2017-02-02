package movements

import (
	"time"

	"github.com/st4rl00rd/dynamics/types"
)

// Movement struct
type Movement struct {
	ID          int64            `db:"id" json:"id"`
	Quantity    float64          `db:"quantity" json:"quantity"`
	Ended       bool             `db:"ended" json:"ended"`
	StoreFromID int64            `db:"store_from_id" json:"store_from_id"`
	StoreToID   int64            `db:"store_to_id" json:"store_to_id"`
	UserFromID  int64            `db:"user_from_id" json:"user_from_id"`
	UserToID    int64            `db:"user_to_id" json:"user_to_id"`
	Details     types.DetailsMap `db:"details" json:"details"`
	SchemeID    int64            `db:"scheme_id" json:"scheme_id"`
	ProductID   int64            `db:"product_id" json:"product_id"`
	CreatedAT   time.Time        `db:"created_at" json:"created_at"`
	UpdatedAT   time.Time        `db:"updated_at" json:"updated_at"`
	Product     Product          `json:"product"`
}

// Product struct for Movement JSONresponse
type Product struct {
	ID           int64            `db:"id" json:"id"`
	SchemeID     int64            `db:"scheme_id" json:"scheme_id"`
	Name         string           `db:"name" json:"name"`
	Details      types.DetailsMap `db:"details" json:"details" json:"details"`
	CreatedAT    time.Time        `db:"created_at" json:"created_at"`
	UpdatedAT    time.Time        `db:"updated_at" json:"updated_at"`
	LastMovement LastMovement     `json:"last_movement"`
}

// LastMovement struct
type LastMovement struct {
	ID          int64            `db:"id" json:"id"`
	Quantity    float64          `db:"quantity" json:"quantity"`
	Ended       bool             `db:"ended" json:"ended"`
	StoreFromID int64            `db:"store_from_id" json:"store_from_id"`
	StoreToID   int64            `db:"store_to_id" json:"store_to_id"`
	UserFromID  int64            `db:"user_from_id" json:"user_from_id"`
	UserToID    int64            `db:"user_to_id" json:"user_to_id"`
	Details     types.DetailsMap `db:"details" json:"details" json:"details"`
	SchemeID    int64            `db:"scheme_id" json:"scheme_id"`
	ProductID   int64            `db:"product_id" json:"product_id"`
	CreatedAT   time.Time        `db:"created_at" json:"created_at"`
	UpdatedAT   time.Time        `db:"updated_at" json:"updated_at"`
}

// JSONResponse struct
type JSONResponse struct {
	Movements []Movement `json:"movements"`
}
