package application

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"warehouse/application"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/tests/factories"
)

func TestRackApplication_All(t *testing.T) {
	sizeOfInventories := 5
	racksList := factories.NewRackFactoryList(sizeOfInventories)
	repo := gorm.NewRackRepository()
	rackApplication := application.NewRackApplication(repo)
	racks, err := rackApplication.All()

	assert.Nil(t, err)
	var ids []string
	for _, rack := range racksList {
		ids = append(ids, rack.Id())
	}

	assert.Contains(t, ids, (*racks)[0].Id())
}

func TestRackApplication_Create(t *testing.T) {
	rack := factories.NewRackDomainFactory()
	repo := gorm.NewRackRepository()
	rackApplication := application.NewRackApplication(repo)
	newRack, err := rackApplication.Create(rack)

	assert.Nil(t, err)
	assert.NotNil(t, newRack.Id())
	assert.Equal(t, rack.Code, newRack.Code)
}

func TestRackApplication_Delete(t *testing.T) {
	rack := factories.NewRackFactory()
	repo := gorm.NewRackRepository()
	rackApplication := application.NewRackApplication(repo)
	err := rackApplication.Delete(rack)
	assert.Nil(t, err)

	findedRack, errFind := repo.Find(rack.Id())

	assert.NotNil(t, errFind)
	assert.IsType(t, errFind, errors.NewNotFoundError(""))
	assert.Nil(t, findedRack)
}

func TestRackApplication_Show(t *testing.T) {
	rack := factories.NewRackFactory()
	repo := gorm.NewRackRepository()
	rackApplication := application.NewRackApplication(repo)
	findedRack, err := rackApplication.Show(rack.Id())

	assert.Nil(t, err)
	assert.Equal(t, rack.Id(), findedRack.Id())
}

func TestRackApplication_Update(t *testing.T) {
	rack := factories.NewRackFactory()
	values := factories.NewRackObjectFactory()

	code := values["code"].(string)
	repo := gorm.NewRackRepository()
	rackApplication := application.NewRackApplication(repo)
	rack.Code = code
	updatedRack, errUpdate := rackApplication.Update(rack)

	assert.Nil(t, errUpdate)
	assert.Equal(t, updatedRack.Code, code)
}

func TestNewRackApplication(t *testing.T) {
	repo := gorm.NewRackRepository()
	rackApplication := application.NewRackApplication(repo)

	assert.NotNil(t, rackApplication)
}
