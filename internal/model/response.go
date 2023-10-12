package model

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	Type string `json:"type"`
}