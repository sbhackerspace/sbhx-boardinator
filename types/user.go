// 2014.02.14

package types

import (
	"time"
)

type User struct {
	Id            [16]byte `json:"id"`
	FirstName     string   `json:"first_name"`
	LastName      string   `json:"last_name"`
	Email         string   `json:"email"`
	IsBoardMember bool     `json:"is_board_member"`
	IsAdmin       bool     `json:"is_admin"`

	CreatedAt  *time.Time `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at"`
}
