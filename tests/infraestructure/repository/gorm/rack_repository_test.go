package gorm

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/tests/factories"
)

func TestRackRepository_All(t *testing.T) {
	sizeOfRacks := 5
	racks := factories.NewRackFactoryList(sizeOfRacks)
	rackRepo := gorm.NewRackRepository()
	allRacks, err := rackRepo.All()
	assert.Nil(t, err)

	var racksIds []string
	for _, rack := range *allRacks {
		racksIds = append(racksIds, rack.Id())
	}

	assert.Contains(t, racksIds, racks[0].Id())
	assert.Contains(t, racksIds, racks[1].Id())
	assert.Contains(t, racksIds, racks[2].Id())
}

func TestRackRepository_Create(t *testing.T) {
	rackRepo := gorm.NewRackRepository()
	rackData := factories.NewRackDomainFactory()

	rack, err := rackRepo.Create(rackData)
	assert.Nil(t, err)

	assert.Equal(t, rack.Name, rackData.Name)
	assert.Equal(t, rack.Code, rackData.Code)
	assert.NotNil(t, rack.Id())
}

func TestRackRepository_Delete(t *testing.T) {
	rack := factories.NewRackFactory()
	rackId := rack.Id()
	assert.NotNil(t, rackId)

	rackRepo := gorm.NewRackRepository()
	rackFounded, err := rackRepo.Find(rackId)
	assert.Nil(t, err)
	assert.NotNil(t, rackFounded)
	assert.Equal(t, rackId, rackFounded.Id())

	assert.Nil(t, rackRepo.Delete(rack))
	rackFounded, err = rackRepo.Find(rackId)
	assert.NotNil(t, err)
	assert.Nil(t, rackFounded)
}

func TestRackRepository_Find(t *testing.T) {
	rack := factories.NewRackFactory()
	rackId := rack.Id()
	assert.NotNil(t, rackId)

	rackRepo := gorm.NewRackRepository()
	rackFounded, err := rackRepo.Find(rackId)
	assert.Nil(t, err)
	assert.NotNil(t, rackFounded)
	assert.Equal(t, rackId, rackFounded.Id())
}

func TestRackRepository_Update(t *testing.T) {
	rack := factories.NewRackFactory()
	newName := gofakeit.LoremIpsumWord()
	rackRepo := gorm.NewRackRepository()
	rack.Name = newName
	rackUpdated, err := rackRepo.Update(rack)

	assert.Nil(t, err)
	assert.NotNil(t, rackUpdated)

	rackFounded, err := rackRepo.Find(rackUpdated.Id())
	assert.Nil(t, err)
	assert.NotNil(t, rackFounded)
	assert.Equal(t, newName, rackFounded.Name)
}

func TestNewRackRepository(t *testing.T) {
	repo := gorm.NewRackRepository()
	assert.NotNil(t, repo)
	assert.IsType(t, gorm.RackRepository{}, *repo)
}
