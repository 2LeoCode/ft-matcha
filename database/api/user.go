package api

type User struct {
	Id           *uint64
	Mail         *string
	Password     *string
	FirstName    *string
	LastName     *string
	Age          *uint8
	Bio          *string
	Country      *string
	City         *string
	LastPosition *Position
	FameRating   *uint32
	Tags         []Tag
	Discussions  []Discussion
	Views        []User
	Kisses       []User
	Punches      []User
}
