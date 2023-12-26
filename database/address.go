package database

import (
	"address/models"
	"strconv"
)

func InsertAddress(address models.Address) (models.Address, error) {
	if err := DB.QueryRow(`
        INSERT INTO addresses (name, longitude, latitude, active, created_at, updated_at, time_zone, complementary_informations, floor, lift, location_type, yard, door_code, loading_dock, side_loading)
        VALUES ($1, $2, $3, $4, NOW(), NOW(), $5, $6, $7, $8, $9, $10, $11, $12, $13)
        RETURNING id, name, longitude, latitude, active, created_at, updated_at,
            time_zone, complementary_informations, floor, lift, location_type, yard,
            door_code, loading_dock, side_loading`,
		address.Name, address.Longitude, address.Latitude, address.Active, address.TimeZone, address.ComplementaryInfo,
		address.Floor, address.Lift, address.LocationType, address.Yard, address.DoorCode,
		address.LoadingDock, address.SideLoading).Scan(
		&address.ID,
		&address.Name,
		&address.Longitude,
		&address.Latitude,
		&address.Active,
		&address.CreatedAt,
		&address.UpdatedAt,
		&address.TimeZone,
		&address.ComplementaryInfo,
		&address.Floor,
		&address.Lift,
		&address.LocationType,
		&address.Yard,
		&address.DoorCode,
		&address.LoadingDock,
		&address.SideLoading,
	); err != nil {
		return models.Address{}, err
	}
	return address, nil
}

func GetAddresses(isActive string, locationType string) ([]models.Address, error) {
	query := "SELECT * FROM addresses WHERE 1=1"
	args := []interface{}{}

	if isActive != "" {
		query += " AND active = $1"
		activeValue, err := strconv.ParseBool(isActive)
		if err != nil {
			return nil, err
		}
		args = append(args, activeValue)
	}

	if locationType != "" {
		query += " AND location_type = $2"
		args = append(args, locationType)
	}

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []models.Address
	for rows.Next() {
		var address models.Address
		err := rows.Scan(
			&address.ID,
			&address.Name,
			&address.Longitude,
			&address.Latitude,
			&address.Active,
			&address.CreatedAt,
			&address.UpdatedAt,
			&address.TimeZone,
			&address.ComplementaryInfo,
			&address.Floor,
			&address.Lift,
			&address.LocationType,
			&address.Yard,
			&address.DoorCode,
			&address.LoadingDock,
			&address.SideLoading,
		)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func GetAddressByID(addressID int) (models.Address, error) {
	var address models.Address
	err := DB.QueryRow("SELECT * FROM addresses WHERE id = $1", addressID).
		Scan(
			&address.ID,
			&address.Name,
			&address.Longitude,
			&address.Latitude,
			&address.Active,
			&address.CreatedAt,
			&address.UpdatedAt,
			&address.TimeZone,
			&address.ComplementaryInfo,
			&address.Floor,
			&address.Lift,
			&address.LocationType,
			&address.Yard,
			&address.DoorCode,
			&address.LoadingDock,
			&address.SideLoading,
		)
	if err != nil {
		return models.Address{}, err
	}
	return address, nil
}

func UpdateAddress(addressID int, updatedAddress models.Address) error {
	_, err := DB.Exec(`
        UPDATE addresses
        SET name=$1, longitude=$2, latitude=$3, active=$4, created_at=$5, updated_at=$6,
            time_zone=$7, complementary_informations=$8, floor=$9, lift=$10,
            location_type=$11, yard=$12, door_code=$13, loading_dock=$14, side_loading=$15
        WHERE id=$16
        RETURNING id, name, longitude, latitude, active, created_at, updated_at,
            time_zone, complementary_informations, floor, lift, location_type, yard,
            door_code, loading_dock, side_loading`,
		updatedAddress.Name, updatedAddress.Longitude, updatedAddress.Latitude,
		updatedAddress.Active, updatedAddress.CreatedAt, updatedAddress.UpdatedAt,
		updatedAddress.TimeZone, updatedAddress.ComplementaryInfo, updatedAddress.Floor,
		updatedAddress.Lift, updatedAddress.LocationType, updatedAddress.Yard,
		updatedAddress.DoorCode, updatedAddress.LoadingDock, updatedAddress.SideLoading,
		addressID,
	)
	return err
}

func DeleteAddress(addressID int) error {
	_, err := DB.Exec("DELETE FROM addresses WHERE id = $1", addressID)
	return err
}
