package config

import (
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

var DB *pg.DB

func InitDatabase() *pg.DB {
	DB = pg.Connect(&pg.Options{
		Addr:     viper.GetString("database.host") + ":" + viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.name"),
	})
	if DB == nil {
		panic("Không thể kết nối đến cơ sở dữ liệu")
	}
	return DB
}

func Transaction(fn func(db *pg.Tx) error) error {
	// Begin transaction.
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	// Rollback on error.
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	// Commit on success.
	return tx.Commit()
}
