package db

import (
	"app/model"
	"app/pkg/error"
)


func (d *DataStore) GetListConnect() (Connects []model.Connect, err *error.Error) {
	e := d.Db.Find(&Connects).Error
	err = error.ParseMysqlError(e)
	return
}

func (d *DataStore) InsertConnect(Connect *model.Connect) (err *error.Error) {
	e := d.Db.Create(Connect).Error
	err = error.ParseMysqlError(e)
	return
}

func (d *DataStore) DeleteConnectByID(id int) (err *error.Error) {
	e := d.Db.Where("id = ?", id).Delete(&model.Connect{}).Error
	err = error.ParseMysqlError(e)
	return
}

func (d *DataStore) UpdateConnect(Connect *model.Connect) (err *error.Error) {
	e := d.Db.Save(&Connect).Error
	err = error.ParseMysqlError(e)
	return
}