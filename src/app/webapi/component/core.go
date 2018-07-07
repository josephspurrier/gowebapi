package component

// NewCore returns the standard component dependencies.
func NewCore(l ILogger, d IDatabase, q IQuery, b IBind, resp IResponse, t IToken) Core {
	return Core{
		Log:      l,
		DB:       d,
		Q:        q,
		Bind:     b,
		Response: resp,
		Token:    t,
	}
}

// Core contains all the dependencies for the components.
type Core struct {
	Log      ILogger
	DB       IDatabase
	Q        IQuery
	Bind     IBind
	Response IResponse
	Token    IToken
}
