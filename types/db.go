// Steve Phillips / elimisteve
// 2014.03.22

package types

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	pgOnce sync.Once
	db     *sql.DB

	postgresConnStr = "user=boardinator dbname=boardinator host=localhost sslmode=disable password=boardinator"
)

func init() {
	pgOnce.Do(initPostgres)
}

func initPostgres() {
	var err error
	db, err = sql.Open("postgres", postgresConnStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("DB ping failed with error `%v`. Connection string: `%v`",
			err, postgresConnStr)
	}

	log.Println("Connected to Postgres")
}
