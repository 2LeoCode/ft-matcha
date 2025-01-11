package enums

type Sex int64

type sexes struct {
	enum[Sex]
	Male,
	Female Sex
}

var Sexes sexes = sexes{
	Male:   0,
	Female: 1,
}
