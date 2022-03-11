//api package to provide endpoint service
package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajjanjyothi/petsalone/pkg/schema"
	"github.com/sajjanjyothi/petsalone/pkg/service"
	"go.uber.org/zap"
)

type PetAPIService struct {
	db *sql.DB
}

func New(db *sql.DB) *PetAPIService {
	return &PetAPIService{
		db: db,
	}
}

// GetAllPets Returns all pets
// (GET /api/pets)
func (apiservice *PetAPIService) GetAllPets(c *gin.Context) {
	petStore := service.PetsService()
	pets, err := petStore.GetAll(apiservice.db)
	if err != nil {
		zap.L().Error("Cannot find the pets " + err.Error())
		c.JSON(http.StatusInternalServerError, handle_error(http.StatusNotFound, "cannot find the pets"))
	}
	c.JSON(http.StatusOK, &pets)
}

// AddPet Creates a new missing pet
// (POST /api/pets)
func (apiservice *PetAPIService) AddPet(c *gin.Context) {
	var pet schema.Pet

	if err := c.ShouldBindJSON(&pet); err != nil {
		zap.L().Error("Cannot create pet " + err.Error())
		c.JSON(http.StatusBadRequest, handle_error(http.StatusBadRequest, "Request invalid"))
	}
	petStore := service.PetsService()
	err := petStore.CreatePets(pet.Name, pet.PetType, pet.MissingSince, apiservice.db)
	if err != nil {
		zap.L().Error("Cannot create pet " + err.Error())
		c.JSON(http.StatusInternalServerError, handle_error(http.StatusInternalServerError, "cannot create pets"))
	}
	c.JSON(http.StatusOK, handle_error(http.StatusOK, "success"))
}

// FindPetByType Returns a specific pet type
// (GET /api/pets/{type})
func (apiservice *PetAPIService) FindPetByType(c *gin.Context, pType string) {
	petStore := service.PetsService()
	pets, err := petStore.GetPetsByType(pType, apiservice.db)
	if err != nil {
		zap.L().Error("Cannot find the pets " + err.Error())
		c.JSON(http.StatusInternalServerError, handle_error(http.StatusNotFound, "cannot find the pets"))
	}
	//No pets found
	if len(pets) == 0 {
		zap.L().Error("Cannot find the pets " + err.Error())
		c.JSON(http.StatusNotFound, handle_error(http.StatusNotFound, "cannot find the pets type"))
	}
	c.JSON(http.StatusOK, pets)
}

//handle_error handle http errors
func handle_error(code int32, message string) *schema.Error {
	errReturn := schema.Error{
		Code:    code,
		Message: message,
	}
	return &errReturn
}
