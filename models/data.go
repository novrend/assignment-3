package models

type Data struct {
	Water int `gorm:"varchar(100)" json:"water"`
	Wind  int `gorm:"varchar(100)" json:"wind"`
}
