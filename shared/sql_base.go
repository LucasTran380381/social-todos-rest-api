package shared

import "time"

type SQLBase struct {
	Id        int        `json:"id"`
	CreatedAt time.Time  `json:"create_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
