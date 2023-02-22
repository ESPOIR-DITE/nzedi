package entity

import "time"

type Account struct {
	Company  int
	Date     time.Time
	Email    string
	Id       int
	Password string
	Token    *string
	UserName *string
}
