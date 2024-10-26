package config

import (
	"context"
	"fmt"

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
	DB.AddQueryHook(dbLogger{})
	return DB
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	bytes, _ := q.FormattedQuery()
	fmt.Println("After query :" + string(bytes))
	return nil
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
