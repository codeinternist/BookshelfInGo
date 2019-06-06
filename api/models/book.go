package models

import (
	"time"
)

type Status int

const (
	CheckedIn		Status = 1
	CheckedOut		Status = 0
)

type Book struct {
	Title			string
	Author			string
	Publisher		string
	PublishDate		time.Time
	Rating			int
	Status			Status
}