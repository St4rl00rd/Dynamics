package products

import (
	"time"

	"github.com/st4rl00rd/dynamics/types"
)

// Product struct
type Product struct {
	ID        int64            `db:"id"`
	SchemeID  int64            `db:"scheme_id"`
	Name      string           `db:"name"`
	Details   types.DetailsMap `db:"details" json:"details"`
	CreatedAT time.Time        `db:"created_at"`
	UpdatedAT time.Time        `db:"updated_at"`
}
