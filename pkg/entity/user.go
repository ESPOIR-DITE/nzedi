package entity

import (
	"time"
)

type User struct {
	Id          int
	AccountId   int
	DateOfBirth time.Time
	FirstName   string
	LastName    string
}
