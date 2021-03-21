package cmd

import (
	systemcontext "context"
	"fmt"
	"net/http"

	"github.com/getchipman/bolt-api/app/cache"
	"github.com/getchipman/bolt-api/app/common"
	"github.com/getchipman/bolt-api/app/context"
	"github.com/getchipman/bolt-api/app/core/domains"
	"github.com/getchipman/bolt-api/app/core/services/authsrv"
	"github.com/getchipman/bolt-api/app/core/services/userssrv"
	"github.com/getchipman/bolt-api/app/db"
	"github.com/getchipman/bolt-api/app/handlers"
	"github.com/getchipman/bolt-api/app/repositories/authrepo"
	"github.com/getchipman/bolt-api/app/repositories/usersrepo"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var apiCMD = &cobra.Command{
	Use:   "web",
	Short: "Start Web Application",
	RunE: func(cmd *cobra.Command, args []string) error {

		// --> New app Context.
		ctx := context.New().WithLogger()

		// --> New postgres database.
		dbConfig := db.NewPostgresConfig()
		db, err := db.NewPostgres(dbConfig)
		if err != nil {
			ctx.Logger.Errorf("Error initilize database instance - Error: %v", err.Error())
			panic(err)
		}

		// --> New redis cache.
		cacheConfig := cache.NewRedisConfig()
		cache, err := cache.NewRedis(cacheConfig)
		if err != nil {
			ctx.Logger.Errorf("Error initilize cache instance - Error: %v", err.Error())
			panic(err)
		}

		// --> Run migrations.
		err = db.RunFileQuery("migrations/" + schema)
		if err != nil {
			ctx.Logger.Errorf("Error to execute migrations - Error: %v", err.Error())
			panic(err)
		}

		// --> Cache Store Session.
		store, err := redis.NewStore(10, "tcp", cacheConfig.Addr, cacheConfig.Password, []byte(common.GetEnv("SECRET_TOKEN", "bolt_api_secret_abc_123")))
		if err != nil {
			ctx.Logger.Errorf("Error initilize cache store instance - Error: %v", err)
			panic(err)
		}

		// --> New Auth repo and service.
		authrepo := authrepo.New(db, cache)
		authsrv := authsrv.New(authrepo)

		// --> New User repo and service.
		usersrepo := usersrepo.New(db, cache)
		userssrv := userssrv.New(usersrepo)

		// --> Handler
		hdl := handlers.New(authsrv, userssrv)

		router := gin.New()

		router.Use(sessions.Sessions(domains.DefaultSessionID, store))
		router.Use(handlers.GlobalMiddleware())

		// Login and Registration
		router.POST("api/v1/user", handlers.HandlerAPI(false, hdl.Create))
		router.POST("api/v1/auth/login", handlers.HandlerAPI(false, hdl.LoginAuth))

		port := common.GetEnv("PORT", "9000")
		s := &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: router,
		}

		systemctx, _ := systemcontext.WithCancel(systemcontext.Background())
		done := make(chan struct{})
		go func() {
			<-systemctx.Done()
			ctx.Logger.Info("Shutting down API, waiting...")
			if err := s.Shutdown(systemcontext.Background()); err != nil {
				ctx.Logger.Error(err)
			}
			close(done)
		}()

		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			ctx.Logger.Error(err)
		}
		<-done

		return nil
	},
}

func init() {
	rootCmd.AddCommand(apiCMD)
}
