package modelos

type Device struct {
	ID        uint64   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Address   string   `json:"address,omitempty"`
	Latitude  float64  `json:"latitude,omitempty"`
	Longitude float64  `json:"longitude,omitempty"`
	Records   []Record `json:"record,omitempty"`
}
