package service

import (
	"database/sql"
	"time"

	"github.com/sajjanjyothi/petsalone/pkg/schema"
)

//jwt service
type PetService interface {
	GetAll(db *sql.DB) ([]schema.Pet, error)
	GetPetsByType(petType string, db *sql.DB) ([]schema.Pet, error)
	CreatePets(name, petType string, missing time.Time, db *sql.DB) error
}

type Pet struct {
	name         string
	petType      string
	missingSince time.Time
}

type petServices struct {
}

//auth-jwt
func PetsService() PetService {
	return &petServices{}
}

//GetAll Get all pets
func (service *petServices) GetAll(db *sql.DB) ([]schema.Pet, error) {
	var pets []schema.Pet
	var unsortedpets []schema.Pet

	row, err := db.Query("SELECT * FROM pets")
	if err != nil {
		return pets, err
	}
	defer row.Close()

	for row.Next() {
		var id int
		var name string
		var petType string
		var missingSince time.Time

		row.Scan(&id, &name, &petType, &missingSince)
		unsortedpets = append(unsortedpets, schema.Pet{Name: name, PetType: petType, MissingSince: missingSince})
	}

	// seems to be the best way to order by newest first
	for _, pet := range unsortedpets {
		if len(pets) == 0 {
			pets = append(pets, pet)
		} else if pets[0].MissingSince.Before(pet.MissingSince) {
			var temp = pets
			pets = []schema.Pet{pet}
			for _, tempitem := range temp {
				pets = append(pets, tempitem)
			}
		} else {
			pets = append(pets, pet)
		}
	}

	return pets, nil
}

//GetPetsByType Get pets with specific type
func (service *petServices) GetPetsByType(petType string, db *sql.DB) ([]schema.Pet, error) {
	var pets []schema.Pet
	var unsortedpets []schema.Pet

	row, err := db.Query("SELECT * FROM pets where pet_type=?", petType)
	if err != nil {
		return pets, err
	}
	defer row.Close()

	for row.Next() {
		var id int
		var name string
		var petType string
		var missingSince time.Time

		row.Scan(&id, &name, &petType, &missingSince)
		unsortedpets = append(unsortedpets, schema.Pet{Name: name, PetType: petType, MissingSince: missingSince})
	}

	// seems to be the best way to order by newest first
	for _, pet := range unsortedpets {
		if len(pets) == 0 {
			pets = append(pets, pet)
		} else if pets[0].MissingSince.Before(pet.MissingSince) {
			var temp = pets
			pets = []schema.Pet{pet}
			for _, tempitem := range temp {
				pets = append(pets, tempitem)
			}
		} else {
			pets = append(pets, pet)
		}
	}

	return pets, nil
}

//CreatePets create new pets
func (service *petServices) CreatePets(name, petType string, missing time.Time, db *sql.DB) error {

	statement, err := db.Prepare("INSERT INTO pets (name, pet_type, missing_since) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(name, petType, time.Now().AddDate(0, 0, -6))
	if err != nil {
		return err
	}
	return nil
}
