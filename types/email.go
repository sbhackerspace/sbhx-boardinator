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
	emailDB    = map[string]*Email{}
	EmailQueue = make(chan *Email)
)

type EmailStatus string

const (
	QUEUED  EmailStatus = "queued"
	SENDING EmailStatus = "sending"
	SUCCESS EmailStatus = "success"
	FAILED  EmailStatus = "failed"
)

type Email struct {
	Id      string `json:"id"`
	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"name"`
	Body    string `json:"body"`

	Status EmailStatus `json:"status"`

	CreatedAt  *time.Time `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at"`
}

func handleEmailError(e *Email, err error) {
	log.Printf("Error occurred with email: %v", err.Error())
	e.Status = FAILED
}

func (e *Email) SaveAndSend() {
	// Save
	go func() {
		if err := e.Save(); err != nil {
			handleEmailError(e, err)
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
			email.Status = QUEUED
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
					handleEmailError(e, err)
					return
				}
				e.Status = SUCCESS
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
	e.CreatedAt = &now
	e.ModifiedAt = &now
}

//TODO
func (e *Email) Send() error {
	e.Status = SENDING
	log.Printf("Email sent: %+v\n", e)
	return nil
}
