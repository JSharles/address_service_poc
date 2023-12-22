package models

type LocationType string

const (
	Individual           LocationType = "individual"
	Company              LocationType = "company"
	RetailStore          LocationType = "retail_store"
	Event                LocationType = "event"
	Supermarket          LocationType = "supermarket"
	Warehouse            LocationType = "warehouse"
	DistributionPlatform LocationType = "distribution_platform"
)

type FloorType string

const (
	Basement    FloorType = "basement"
	Sidewalk    FloorType = "sidewalk"
	GroundFloor FloorType = "ground_floor"
	First       FloorType = "first"
	Second      FloorType = "second"
	Third       FloorType = "third"
	Fourth      FloorType = "fourth"
	Fifth       FloorType = "fifth"
	Sixth       FloorType = "sixth"
	Seventh     FloorType = "seventh"
)

type YardType string

const (
	None           YardType = "none"
	Inf_10m        YardType = "inf_10m"
	Between_10_30m YardType = "between_10_30m"
	Sup_30m        YardType = "sup_30m"
)

type Address struct {
	ID                int          `db:"id" json:"id"`
	Name              string       `db:"name" json:"name" validate:"required"`
	Longitude         float64      `db:"longitude" json:"longitude" validate:"required"`
	Latitude          float64      `db:"latitude" json:"latitude" validate:"required"`
	Active            bool         `db:"active" json:"active"`
	CreatedAt         string       `db:"created_at" json:"created_at"`
	UpdatedAt         string       `db:"updated_at" json:"updated_at"`
	TimeZone          string       `db:"time_zone" json:"time_zone"`
	ComplementaryInfo string       `db:"complementary_informations" json:"complementary_informations"`
	Floor             FloorType    `db:"floor" json:"floor"`
	Lift              string       `db:"lift" json:"lift"`
	LocationType      LocationType `db:"location_type" json:"location_type"`
	Yard              YardType     `db:"yard" json:"yard"`
	DoorCode          string       `db:"door_code" json:"door_code"`
	LoadingDock       bool         `db:"loading_dock" json:"loading_dock"`
	SideLoading       bool         `db:"side_loading" json:"side_loading"`
}
