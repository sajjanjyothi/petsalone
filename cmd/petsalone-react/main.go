package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sajjanjyothi/petsalone/pkg/api"
	"github.com/sajjanjyothi/petsalone/pkg/auth"
	"github.com/sajjanjyothi/petsalone/pkg/memorydb"
	"github.com/sajjanjyothi/petsalone/pkg/service"
	"go.uber.org/zap"
)

func main() {
	var loginService auth.LoginService = auth.BasicLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController auth.LoginController = auth.LoginHandler(loginService, jwtService)

	router := gin.Default()

	//Initialize the logger
	zapLogger, _ := zap.NewProduction()
	zap.ReplaceGlobals(zapLogger)
	defer func() {
		_ = zapLogger.Sync()
	}()
	db, err := memorydb.GetDB()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	router.Use(static.Serve("/", static.LocalFile("ui/react/build", true)))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	apiService := api.New(db)

	//TODO: SAJJAN: Disabling JWT token authentication
	// if gin.Mode() != gin.TestMode {
	// 	router.Use(middleware.AuthorizeJWT())
	// }

	//Register all handlers
	router = api.RegisterHandlers(router, apiService)

	api := router.Group("/api")
	{
		api.POST("users/login", func(ctx *gin.Context) {
			token := loginController.Login(ctx, db)
			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, nil)
			}
		})
	}
	router.Run(":8080")
}
