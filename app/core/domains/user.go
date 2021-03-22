package domains

// User represent a user in the application.
type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
