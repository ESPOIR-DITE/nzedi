package entity

import (
	"time"
)

type User struct {
	Id          string
	AccountId   string
	DateOfBirth time.Time
	FirstName   string
	LastName    string
}
