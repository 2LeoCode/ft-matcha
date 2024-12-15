package api

type Sex uint8

const (
	Male Sex = iota
	Female
)

type Orientation uint8

const (
	Bisexual Orientation = iota
	Heterosexual
	Homosexual
)

type User struct {
	Id          uint64
	Mail        string
	Password    string
	FirstName   string
	LastName    string
	Age         uint8
	Sex         Sex
	Orientation Orientation
	AgeGap
	Bio          string
	Country      string
	City         string
	LastPosition Position

	Tags        []Tag
	Discussions []Discussion
	Views       []User
	Kisses      []User
	Punches     []User
	Blocked     []User
	BlockedBy   []User
}

func (this *User) FameRating() float64 {
	return 0
}
