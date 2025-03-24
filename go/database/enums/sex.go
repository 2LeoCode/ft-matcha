package enums

import "matcha/go/utils"

type Sex int64

type sexes struct {
	utils.Enum[Sex]
	Male,
	Female Sex
}

var Sexes sexes = sexes{
	Male:   0,
	Female: 1,
}
