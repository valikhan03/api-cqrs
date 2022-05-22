package repository

import(
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func InitDB() (*gorm.DB, error){
	db, err := gorm.Open(postgres.Open(""))
	if err != nil{
		return nil, err
	}

	return db, nil
}