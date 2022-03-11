package service

import (
	"testing"
	"time"

	"github.com/sajjanjyothi/petsalone/pkg/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, err := memorydb.GetDB()
	assert.NoError(t, err)
	petStore := &petServices{}
	pets, err := petStore.GetAll(db)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(pets))

}

func TestGetByType(t *testing.T) {
	db, err := memorydb.GetDB()
	assert.NoError(t, err)
	petStore := &petServices{}
	pets, err := petStore.GetPetsByType("dog", db)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(pets))
}

func TestCreatePet(t *testing.T) {
	db, err := memorydb.GetDB()
	assert.NoError(t, err)
	petStore := &petServices{}
	err = petStore.CreatePets("milo", "dog", time.Now(), db)
	assert.NoError(t, err)
	pets, err := petStore.GetAll(db)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(pets))
}
