package domain

type UserClaims struct {
	ID        string `json:"user_id"`
	Name      string `json:"name"`
	GivenName string `json:"given_name"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
}
