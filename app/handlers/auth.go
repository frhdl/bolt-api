package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/frhdl/bolt-api/app/common"
	"github.com/frhdl/bolt-api/app/context"
	"github.com/frhdl/bolt-api/app/core/domains"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginAuth return a new token.
func (hdl *HTTPHandler) LoginAuth(ctx *context.Context, c *gin.Context) error {

	params := &domains.User{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		ctx.ResultError(1, "Invalid Request").JSON(c, nil)
		return nil
	}

	result, userID, userName, userEmail := hdl.authService.Login(ctx, params)
	if result.Error != nil || userID == 0 {
		result.JSON(c, nil)
		return nil
	}

	claims := &domains.JWTClaims{
		userID,
		false,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(common.GetEnv("SECRET_TOKEN", "bolt_api_secret_abc_123")))
	if err != nil {
		return err
	}

	session := sessions.Default(c)
	session.Set("access_token", t)
	session.Delete("state")
	err = session.Save()
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": t,
		"token_type":   "Bearer",
		"expires_in":   claims.ExpiresAt,
		"user_name":    userName,
		"user_email":   userEmail,
	})

	return nil
}
