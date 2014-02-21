// 2014.02.14

package types

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"time"
)

// TODO: Replace with Postgres DB
var (
	emailDB = map[string]*Email{}
	EmailQueue = make(chan *Email)
)

type Email struct {
	Id          string    `json:"id"`
	To			string    `json:"to"`
	From		string    `json:"from"`
	Subject     string    `json:"name"`
	Body		string    `json:"body"`

	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type SentStatus string

const (
	QUEUED  SentStatus = "queued"
	SENDING SentStatus = "sending"
	SUCCESS SentStatus = "success"
	FAILED  SentStatus = "failed"
)

type EmailStatus struct {
	Email   *Email
	EmailId string
	Status  SentStatus
}

func (e *Email) SaveAndSend() {
	// Save
	go func() {
		if err := e.Save(); err != nil {
			//TODO: Handle error
		}
	}()
	// Send
	go func() {
		EmailQueue <- e
	}()
}

func StartEmailQueue() {
	var semaphore = make(chan bool, 10) // Can only send 10 at once

	go func() {
		for {
			email := <-EmailQueue
			// Spawn goroutine immediately
			go func(e *Email) {
				// Will only block if 10 emails are already being sent.
				semaphore <- true
				defer func() {
					// Drain value from channel to make room
					// for another email sender
					<-semaphore
				}()

				err := e.Send()
				if err != nil {
					// FAILED
					//TODO: Handle error
					return
				}
				// SUCCESS
			}(email)
		}
	}()
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
