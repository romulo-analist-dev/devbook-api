package modelos

import "time"

type Record struct {
	ID        uint64    `json:"id,omitempty"`
	Value     float64   `json:"value,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Device    *Device   `json:"device,omitempty"`
}
