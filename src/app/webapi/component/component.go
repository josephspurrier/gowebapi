package component

// New returns the standard component dependencies.
func New(l ILogger, d IDatabase, b IBind, resp IResponse) Core {
	return Core{
		Log:      l,
		DB:       d,
		Bind:     b,
		Response: resp,
	}
}

// Core contains all the dependencies.
type Core struct {
	Log      ILogger
	DB       IDatabase
	Bind     IBind
	Response IResponse
}
