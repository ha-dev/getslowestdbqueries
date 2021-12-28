package err

import (
	"errors"
)

var (
	ErrorNotFoundEnfPoint   = errors.New("sorry, endpoint is not found").Error()
	ErrorDatabaseConnection = errors.New("sorry, database connection have a problem")
	ErrorInsertSlowestQuery = errors.New("sorry, insert slowest query have a problem")
	ErrorExecuteQuery       = errors.New("sorry, run your query have a problem")
)
