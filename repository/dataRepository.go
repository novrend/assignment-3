package repository

import (
	"assignment-3/database"
	"assignment-3/models"
)

func UpdateData(bodyJson models.Data) error {
	db := database.StartDB()
	defer db.Close()

	sqlStatement := `UPDATE data SET water = $1, wind = $2;`
	_, err := db.Exec(sqlStatement, bodyJson.Water, bodyJson.Wind)
	if err != nil {
		return err
	}
	return nil
}
