package handlers

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/getchipman/bolt-api/app/common"
	"github.com/getchipman/bolt-api/app/context"
	"github.com/getchipman/bolt-api/app/core/domains"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

// RequestContextKey represent the name of the context.
const RequestContextKey = "request_context"

//HandlerAPI .
func HandlerAPI(mustBeAuthenticate bool, f func(ctx *context.Context, c *gin.Context) error) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		ctxVal := c.MustGet(RequestContextKey).(*context.Context)

		if mustBeAuthenticate && ctxVal.LoggedUserID == 0 {
			c.Status(401)
			return
		}

		if err := f(ctxVal, c); err != nil {
			ctxVal.Logger.Error(fmt.Errorf("%v: %s", err, debug.Stack()))
			http.Error(c.Writer, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}

// GlobalMiddleware is default middleware for all endpoints.
func GlobalMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		//start context
		ctx := context.New().WithLogger()
		ctx.HTTPPrefix = common.GetEnv("HTTPPREFIX", "")

		defer func() {
			if r := recover(); r != nil {
				http.Error(c.Writer, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		session := sessions.Default(c)
		token := session.Get("access_token")
		tokenByAPI := c.Request.Header["Authorization"]

		//Check api request
		if len(tokenByAPI) > 0 {
			claims := &domains.JWTClaims{}
			token := strings.Split(tokenByAPI[0], " ")
			ctx.Logger.Info("length: %v", len(token))
			if len(token) != 2 {
				c.Set(RequestContextKey, ctx)
				ctx.ResultError(0, "Token Format Invalid - Correct Format: Bearer %token").JSON(c, nil)
				c.Status(401)
				return
			}
			_, err := jwt.ParseWithClaims(token[1], claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(common.GetEnv("SECRET_TOKEN", "bolt_api_secret_abc_123")), nil
			})

			if err == nil && claims != nil {
				ctx.IsAdmin = claims.IsAdmin
				ctx.LoggedUserID = claims.LoggedUserID
			} else {
				ctx.Logger.WithField("Error", err).Errorf("Error to validate token")
			}
		} else if token != nil {
			//Check session request
			claims := &domains.JWTClaims{}
			_, err := jwt.ParseWithClaims(token.(string), claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(common.GetEnv("SECRET_TOKEN", "bolt_api_secret_abc_123")), nil
			})

			if err == nil {
				if claims.LoggedUserID != 0 {
					ctx.IsAdmin = claims.IsAdmin
					ctx.LoggedUserID = claims.LoggedUserID
				}
			} else {
				ctx.Logger.WithField("Error", err).Errorf("Error to validate token")
			}
		}

		c.Set(RequestContextKey, ctx)
		c.Next()
	})
}
