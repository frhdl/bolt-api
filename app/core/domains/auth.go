package domains

import "github.com/dgrijalva/jwt-go"

//DefaultSessionID .
const DefaultSessionID = "bolt_api_session"

// JWTClaims .
type JWTClaims struct {
	LoggedUserID int  `json:"logged_user_id"`
	IsAdmin      bool `json:"is_admin"`
	jwt.StandardClaims
}
