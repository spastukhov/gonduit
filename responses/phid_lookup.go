package responses

import "github.com/uber/gonduit/entities"

// PHIDLookupResponse is the result of phid.lookup operations.
type PHIDLookupResponse map[string]*entities.PHIDResult
