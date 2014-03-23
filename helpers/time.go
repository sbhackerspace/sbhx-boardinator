// Steve Phillips / elimisteve
// 2014.03.23

package helpers

import (
	"time"
)

var (
	LosAngeles, _ = time.LoadLocation("America/Los_Angeles")
)

func Now() time.Time {
	return time.Now().In(LosAngeles).Round(time.Second)
}
