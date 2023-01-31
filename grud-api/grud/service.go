package grud

import (
	"errors"
	database "github.com/Lefree111/gorm-resta-api/api/entity"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func CreateApi(a *database.Data) (*database.Data, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return &database.Data{}, errors.New("api not added")
	}
	return a, nil
}

func GetApi(id int) (*database.Data, error) {
	var api database.Data
	res := db.First(&api, id)
	if res.RowsAffected == 0 {
		return nil, errors.New("api not found")
	}
	return &api, nil
}

func GetApis() ([]*database.Data, error) {
	var apis []*database.Data
	res := db.Find(&apis)
	if res.Error != nil {
		return nil, errors.New("apis not found")
	}
	return apis, nil
}

func UpdateApi(api *database.Data) (*database.Data, error) {
	var update_api database.Data
	res := db.Model(&update_api).Where(api.ID).Update(api)
	if res.RowsAffected == 0 {
		return &database.Data{}, errors.New("api not update")
	}
	return &update_api, nil
}

func DeleteApi(id int) error {
	var deleteApi database.Data
	res := db.Where(id).Delete(&deleteApi)
	if res.RowsAffected == 0 {
		return errors.New("api not deleted")
	}
	return nil
}
