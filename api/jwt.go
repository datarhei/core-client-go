package api

// JWT is the JWT token and its expiry date
type JWT struct {
	Token  string `json:"token" jsonschema:"minLength=1"`
	Expire string `json:"expire" jsonschema:"format=date-time"`
}
