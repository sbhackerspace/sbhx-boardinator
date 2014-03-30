// Steve Phillips / elimisteve
// 2014.03.29

package types

import (
	"log"
	"os"
	"sync"

	"github.com/mattbaird/gochimp"
)

var (
	drillOnce sync.Once
	mandrill  *gochimp.MandrillAPI

	MANDRILL_KEY  = os.Getenv("MANDRILL_KEY")
	MANDRILL_USER = os.Getenv("MANDRILL_USER")
)

func init() {
	drillOnce.Do(initMandrill)
	log.Printf("MANDRILL_KEY  == %s\n", MANDRILL_KEY)
	log.Printf("MANDRILL_USER == %s\n", MANDRILL_USER)
	log.Printf("mandrill == %v\n", mandrill)
}

func initMandrill() {
	var err error
	mandrill, err = gochimp.NewMandrill(MANDRILL_KEY)
	if err != nil {
		log.Fatalf("Mandrill error: %v\n", err)
	}
}
