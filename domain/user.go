package domain

type UserClaims struct {
	Name      string `json:"name"`
	GivenName string `json:"given_name"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
}
