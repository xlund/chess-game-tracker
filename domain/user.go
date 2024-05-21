package domain

type UserClaims struct {
	Name      string `json:"name"`
	GivenName string `json:"givenName"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
}
