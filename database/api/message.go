package api

import "time"

type Message struct {
	CreationTime time.Time
	Content      string

	Discussion *Discussion
	Author     *User
}
