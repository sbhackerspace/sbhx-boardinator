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
	once sync.Once
	db   *sql.DB

	postgresConnStr = "user=boardinator dbname=boardinator host=localhost sslmode=disable password=boardinator"
)

func init() {
	once.Do(initPostgres)
}

func initPostgres() {
	var err error
	db, err = sql.Open("postgres", postgresConnStr)
	if err != nil {
		log.Fatal(err)
	}

	// Create `tasks`
	_, err = db.Query(createTableTasks)
	if err != nil {
		log.Printf("Error creating tasks table: %v\n", err)
	}

	log.Println("Connected to Postgres (maybe)")
}

var createTableTasks = `CREATE TABLE tasks (
    Id          varchar(36) NOT NULL,
    Name        varchar(100) NOT NULL,
    Description varchar(4096),
    DueDate     timestamp with time zone,
    Assignee    varchar(100)
);`
