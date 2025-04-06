package schema

type LoginRequestJSON struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseJSON struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
