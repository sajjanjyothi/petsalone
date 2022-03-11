//integrationtest to cover integrations test
package integrationtest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sajjanjyothi/petsalone/pkg/api"
	"github.com/sajjanjyothi/petsalone/pkg/memorydb"
	"github.com/sajjanjyothi/petsalone/pkg/schema"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestMain(m *testing.M) {
	//var loginService auth.LoginService = auth.BasicLoginService()
	//var jwtService service.JWTService = service.JWTAuthService()
	//var loginController auth.LoginController = auth.LoginHandler(loginService, jwtService)

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

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	apiService := api.New(db)
	//Register all handlers
	router = api.RegisterHandlers(router, apiService)

	go func() {
		router.Run(":8080")
	}()
	m.Run()
}
func TestGetAll(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/api/pets")
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	var pets []schema.Pet
	err = json.Unmarshal(body, &pets)
	assert.NoError(t, err)

	assert.Equal(t, 3, len(pets))
}

func TestSpecifcPetType(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/api/pets/dog")
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	var pets []schema.Pet
	err = json.Unmarshal(body, &pets)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(pets))
}

func TestCreatePet(t *testing.T) {
	values := map[string]string{
		"missingSince": "2022-02-28T18:55:20.904841Z",
		"name":         "milo",
		"petType":      "dog",
	}
	jsonData, err := json.Marshal(&values)
	assert.NoError(t, err)

	resp, err := http.Post("http://localhost:8080/api/pets", "application/json",
		bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()

	resp, err = http.Get("http://localhost:8080/api/pets")
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	var pets []schema.Pet
	t.Log(string(body))
	err = json.Unmarshal(body, &pets)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(pets))
}
