package api

import "time"

type Discussion struct {
	Id           uint64
	CreationTime time.Time

	Users    []User
	Messages []Message
}
