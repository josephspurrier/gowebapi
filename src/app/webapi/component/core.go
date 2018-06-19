package component

// NewCore returns the standard component dependencies.
func NewCore(l ILogger, d IDatabase, b IBind, resp IResponse, t IToken) Core {
	return Core{
		Log:      l,
		DB:       d,
		Bind:     b,
		Response: resp,
		Token:    t,
	}
}

// Core contains all the dependencies for the components.
type Core struct {
	Log      ILogger
	DB       IDatabase
	Bind     IBind
	Response IResponse
	Token    IToken
}
