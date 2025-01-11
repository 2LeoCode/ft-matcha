package globals

import "fmt"

var AllowedMethods = []string{"GET", "POST", "OPTIONS"}
var AllowedOrigin = fmt.Sprintf("http://%s", ServerHost)

const MaxAge = 600
const AllowCredentials = true
