package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sajjanjyothi/petsalone/pkg/auth"
	"github.com/sajjanjyothi/petsalone/pkg/memorydb"
	"github.com/sajjanjyothi/petsalone/pkg/service"
)

func main() {
	var loginService auth.LoginService = auth.BasicLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController auth.LoginController = auth.LoginHandler(loginService, jwtService)

	db, err := memorydb.GetDB()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("ui/angular/dist", true)))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

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
