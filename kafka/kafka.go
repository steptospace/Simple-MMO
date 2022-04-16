package kafka

import "time"

type Character struct {
	ID       string
	CharName string
	History  string
}

type UserData struct {
	Username   string
	Characters []Character
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
