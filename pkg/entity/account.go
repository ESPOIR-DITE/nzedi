package entity

import "time"

type Account struct {
	Date     time.Time
	Email    string
	Id       int
	Password string
	Token    *string
	Username string
}
