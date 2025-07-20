package model

import "time"

type ChatCreate struct {
	UserNames []string
}

type MessageCreate struct {
	ChatId    int64
	From      string
	Text      string
	Timestamp time.Time
}
