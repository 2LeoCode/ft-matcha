package api

type Tag struct {
	Id   uint64
	Name string

	Users []User
}
