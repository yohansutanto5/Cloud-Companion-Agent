package service

import (
	"app/db"
	"app/model"
	"app/pkg/error"
)

// ConnectService defines the interface for managing Connects.
type ConnectService interface {
	Insert(Connect *model.Connect) *error.Error
	GetList() ([]model.Connect, *error.Error)
}

type ConnectServiceImpl struct {
	db *db.DataStore
}

func NewConnectService(db *db.DataStore) ConnectService {
	return &ConnectServiceImpl{db: db}
}

// Function Implementation

func (s ConnectServiceImpl) GetList() ([]model.Connect, *error.Error) {
	return s.db.GetListConnect()
}

func (s *ConnectServiceImpl) Insert(Connect *model.Connect) *error.Error {
	return s.db.InsertConnect(Connect)
}