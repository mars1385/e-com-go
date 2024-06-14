package types

type LoginUser struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Access  string
	Refresh string
}
