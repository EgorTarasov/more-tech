package model

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	Type        string `json:"type"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
