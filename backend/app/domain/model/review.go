package model

import "time"

type Review struct {
	AuthorName      string
	ProfilePhotoUrl string
	Rating          int
	Time            time.Time
	Text            string
}
