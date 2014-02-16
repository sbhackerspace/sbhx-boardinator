// 2014.02.14

package types

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"time"
)

// TODO: Replace with Postgres DB
var emailDB = map[string]*Email{}

type Email struct {
	Id          string    `json:"id"`
	To			string    `json:"to"`
	From		string    `json:"from"`
	Subject     string    `json:"name"`
	Body		string    `json:"body"`

	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (e *Email) Save() error {
	if e == nil {
		return fmt.Errorf("Cannot save nil *Email to DB!\n")
	}
	// Populate fields
	id, err := uuid.NewV4() // TODO: Replace with channel read
	if err != nil {
		return fmt.Errorf("Unable to create new id: %v\n", err)
	}
	idStr := id.String()
	e.Id = idStr
	e.populateNew()

	// TODO: Replace with real DB
	emailDB[idStr] = e

	log.Printf("Email saved: %+v\n", e)
	return nil
}

func (e *Email) populateNew() {
	now := time.Now()
	e.CreatedAt = now
	e.ModifiedAt = now
}

func (e *Email) Send() error {
	//TODO
	log.Printf("Sent!")
	log.Printf("(*Email).Send: TODO\n")
	return nil
}
