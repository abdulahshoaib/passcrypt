package presenter

import (
	"encoding/json"

	"github.com/abdulahshoaib/pkg/entities"
)

type ValutServiceEntry struct {
	Service string `json:"service"`

	Email             string            `json:"email"`
	Username          string            `json:"Username "`
	Name              string            `json:"name"`
	MainPassword      string            `json:"main_password"`
	SecondaryPassword string            `json:"secondary_password"`
	Mobile            string            `json:"mobile"`
	TwoFAEnabled      bool              `json:"2fa_enabled"`
	TwoFAType         string            `json:"2fa_type"`
	RecoveryCodes     []string          `json:"recovery-codes"`
	AccessToken       string            `json:"access_token"`
	RefreshToken      string            `json:"refresh_token"`
	ApiKey            string            `json:"api_key"`
	ClientId          string            `json:"client_id"`
	ClientSecret      string            `json:"client_secret"`
	TokenType         string            `json:"token_type,omitempty"`
	Scopes            []string          `json:"scopes,omitempty"`
	TokenIssuedAt     string            `json:"token_issued_at,omitempty"`
	TokenExpiresAt    string            `json:"token_expires_at,omitempty"`
	SessionCookie     string            `json:"session_cookie,omitempty"`
	CustomHeaders     map[string]string `json:"custom_headers,omitempty"`
	SecurityQA        []SecurityQA      `json:"security_questions,omitempty"`
	Tags              []string          `json:"tags,omitempty"`
	CustomFields      map[string]string `json:"custom_fields,omitempty"`
}

type SecurityQA struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
