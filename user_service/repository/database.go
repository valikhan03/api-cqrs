package repository

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/stdlib"
)
func InitDB() (*sqlx.DB, error){
	err := godotenv.Load()
	if err != nil{
		return nil, err
	}
	return sqlx.Connect("postgres", os.Getenv("DB_CONN_STR"))
}