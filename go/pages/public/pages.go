package public

import (
	"matcha/go/pages"
	"matcha/go/pages/public/auth"
)

var Pages = []*pages.Page{
	auth.Page,
}
