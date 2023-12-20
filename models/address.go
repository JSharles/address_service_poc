package models

type Address struct {
	ID                int     `db:"id" json:"id"`
	Name              string  `db:"name" json:"name"`
	Longitude         float64 `db:"longitude" json:"longitude"`
	Latitude          float64 `db:"latitude" json:"latitude"`
	Active            bool    `db:"active" json:"active"`
	CreatedAt         string  `db:"created_at" json:"created_at"`
	UpdatedAt         string  `db:"updated_at" json:"updated_at"`
	TimeZone          string  `db:"time_zone" json:"time_zone"`
	ComplementaryInfo string  `db:"complementary_informations" json:"complementary_informations"`
	Floor             string  `db:"floor" json:"floor"`
	Lift              string  `db:"lift" json:"lift"`
	LocationType      string  `db:"location_type" json:"location_type"`
	Yard              string  `db:"yard" json:"yard"`
	DoorCode          string  `db:"door_code" json:"door_code"`
	LoadingDock       bool    `db:"loading_dock" json:"loading_dock"`
	SideLoading       bool    `db:"side_loading" json:"side_loading"`
}
