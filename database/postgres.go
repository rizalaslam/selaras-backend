// Package database implements postgres connection and queries.
package database

import (
	"log"

	"github.com/spf13/viper"

	"github.com/go-pg/pg"
)

// DBConn returns a postgres connection pool.
func DBConn() (*pg.DB, error) {
	viper.SetDefault("pg_url", "postgres://etvubvflgyijwa:8022939236364c1680930cd47b5ade7b263bc968b8357cd1b6f1c0538c41199e@ec2-54-173-77-184.compute-1.amazonaws.com:5432/devil1nej73nmr")

	opt, err := pg.ParseURL(viper.GetString("pg_url"))
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

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
