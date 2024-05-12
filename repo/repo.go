package repo

import (
	"context"
	"log"

	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	_ "github.com/go-pg/pgext"
)

type DB struct {
	*pg.DB
}

func NewDB() *DB {
	opt, err := pg.ParseURL("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Println("err ParseURL " + err.Error())
		panic(err)
	}
	db := pg.Connect(opt)
	log.Println("db")
	log.Println(db)

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		log.Println("err Ping " + err.Error())
		panic(err)
	}
	return &DB{db}
}

func InitSchema(db *DB) error {
	models := []interface{}{
		(*User)(nil),
		(*Account)(nil),
		(*Transaction)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			log.Println("err CreateTable " + err.Error())
			return err
		}
	}

	return nil
}
