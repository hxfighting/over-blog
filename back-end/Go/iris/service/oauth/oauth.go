package oauth

type OauthConfig struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	Callback string `json:"callback"`
}
