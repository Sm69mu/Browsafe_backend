package models

import "time"

type Bookmark struct {
	ID       string    `json:"id"`
	URL      string    `json:"url"`
	CreateAt time.Time `json:"createAt"`
	UserID   string    `json:"userId"`
}
