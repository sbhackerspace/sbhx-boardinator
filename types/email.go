// 2014.02.14

package types

import (
	"fmt"
	"log"
	"time"

	"github.com/mattbaird/gochimp"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/sbhackerspace/sbhx-boardinator/helpers"
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
	Id      string   `json:"id"`
	To      []string `json:"to"`
	From    string   `json:"from"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`

	Status EmailStatus `json:"status"`

	CreatedAt  *time.Time `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at"`
}

func (e *Email) Failed(err error) {
	log.Printf("Error occurred with email: %v", err.Error())
	e.Status = FAILED
}

func (e *Email) SaveAndSend() {
	done := make(chan struct{})

	// Save
	go func() {
		if err := e.Save(); err != nil {
			e.Failed(err)
		}
		done <- struct{}{}
	}()

	// Send
	EmailQueue <- e

	// Wait for Save to finish before returning
	<-done
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
					e.Failed(err)
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
	now := helpers.Now()
	e.CreatedAt = &now
	e.ModifiedAt = &now
}

func (e *Email) Send() error {
	e.Status = SENDING

	msg := gochimp.Message{
		To: genRecipients(e.To),
		Subject: e.Subject,
		Text: e.Body,
	}
	async := false
	responses, err := mandrill.MessageSend(msg, async)
	if err != nil {
		e.Status = FAILED
		return err
	}
	e.Status = SUCCESS

	// Give us the good news
	for i := 0; i < len(responses); i++ {
		log.Printf("%s: %s\n", responses[i].Email, responses[i].Status)
	}

	log.Printf("Email sent: %+v\n", e)
	return nil
}

func genRecipients(emails []string) []gochimp.Recipient {
	recip := make([]gochimp.Recipient, len(emails))
	for i := 0; i < len(emails); i++ {
		recip[i] = gochimp.Recipient{Email: emails[i]}
	}
	return recip
}
