package err

import (
	"errors"
)

var (
	ErrorNotFoundEnfPoint = errors.New("sorry, endpoint is not found").Error()
)
