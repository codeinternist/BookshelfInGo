package models

import (
	"time"
)

type Status bool

const (
	CheckedIn		Status = true
	CheckedOut		Status = false
)

type Book struct {
	Title			string
	Author			string
	Publisher		string
	PublishDate		time.Time
	Rating			int
	Status			Status
}