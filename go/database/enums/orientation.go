package enums

import "matcha/go/utils"

type Orientation int64

type orientations struct {
	utils.Enum[Orientation]
	Bisexual,
	Heterosexual,
	Homosexual Orientation
}

var Orientations orientations = orientations{
	Bisexual:     0,
	Heterosexual: 1,
	Homosexual:   2,
}
