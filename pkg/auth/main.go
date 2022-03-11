package auth

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/sajjanjyothi/petsalone/pkg/service"
)

type LoginService interface {
	LoginUser(email string, password string, db *sql.DB) bool
}

type LoginInformation struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func BasicLoginService() LoginService {
	return &LoginInformation{}
}

func (info *LoginInformation) LoginUser(username string, password string, db *sql.DB) bool {
	statement, err := db.Prepare("SELECT COUNT(*) FROM users WHERE username=? AND password_hash=?")
	if err != nil {
		fmt.Println("Prepare", err)
		return false
	}
	defer statement.Close()

	var output string
	err = statement.QueryRow(username, fmt.Sprintf("%x", md5.Sum([]byte(password)))).Scan(&output)
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		return false
	default:
		count, _ := strconv.Atoi(output)
		fmt.Printf("Counted %s rows\n", output)
		if count > 0 {
			return true
		}
		return false
	}
}

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context, db *sql.DB) string
}

type loginController struct {
	loginService LoginService
	jwtService   service.JWTService
}

func LoginHandler(loginService LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context, db *sql.DB) string {
	var credential LoginInformation
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	spew.Dump(credential)
	isUserAuthenticated := controller.loginService.LoginUser(credential.Username, credential.Password, db)
	spew.Dump("authed", isUserAuthenticated)
	if isUserAuthenticated {
		return controller.jwtService.GenerateToken(credential.Username, true)
	}
	return ""
}
