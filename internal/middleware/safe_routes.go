package middleware

// Routes these routes will be skipped by the middleware
var Routes = []string{
	"/rentio/api/v1/users/login",
	"/rentio/api/v1/users/register",
	"/rentio/api/v1/clients/login",
	"/rentio/api/v1/clients/register",
}
