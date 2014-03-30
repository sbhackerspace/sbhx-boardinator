// Steve Phillips / elimisteve
// 2014.03.29

package types

import (
	"log"
	"os"
	"sync"

	"github.com/sbhackerspace/sbhx-gomandrill/messages"
)

var (
	drillOnce sync.Once
	sender    *messages.MandrillMessageSender

	mandrillKey = os.Getenv("MANDRILL_KEY")
)

func init() {
	drillOnce.Do(initMandrill)
}

func initMandrill() {
	var err error
	sender, err = messages.NewSender(mandrillKey)
	if err != nil {
		log.Fatalf("Mandrill error: %v\n", err)
	}
}
