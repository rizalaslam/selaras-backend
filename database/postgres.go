// Package database implements postgres connection and queries.
package database

import (
	"log"

	"github.com/spf13/viper"

	"github.com/go-pg/pg"
)

// DBConn returns a postgres connection pool.
func DBConn() (*pg.DB, error) {
	viper.SetDefault("db_network", "tcp")
	viper.SetDefault("db_addr", "port=5432:5432")
	viper.SetDefault("db_user", "etvubvflgyijwa")
	viper.SetDefault("db_password", "8022939236364c1680930cd47b5ade7b263bc968b8357cd1b6f1c0538c41199e")
	viper.SetDefault("db_database", "devil1nej73nmr")

	db := pg.Connect(&pg.Options{
		Network:  viper.GetString("db_network"),
		Addr:     viper.GetString("db_addr"),
		User:     viper.GetString("db_user"),
		Password: viper.GetString("db_password"),
		Database: viper.GetString("db_database"),
	})

	if err := checkConn(db); err != nil {
		return nil, err
	}

	if viper.GetBool("db_debug") {
		db.AddQueryHook(&logSQL{})
	}

	return db, nil
}

type logSQL struct{}

func (l *logSQL) BeforeQuery(e *pg.QueryEvent) {}

func (l *logSQL) AfterQuery(e *pg.QueryEvent) {
	query, err := e.FormattedQuery()
	if err != nil {
		panic(err)
	}
	log.Println(query)
}

func checkConn(db *pg.DB) error {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}
