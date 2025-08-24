package service

import (
	"app/db"
	"app/model"
	"app/pkg/error"
	"time"
)

// ChaosService defines the interface for managing Chaoss.
type ChaosService interface {
	Attack(*model.Chaos) *error.Error
}

type ChaosServiceImpl struct {
	db *db.DataStore
}

func NewChaosService(db *db.DataStore) ChaosService {
	return &ChaosServiceImpl{db: db}
}

// Function Implementation

func (s ChaosServiceImpl) Attack(in *model.Chaos) *error.Error {
	stop := make(chan bool)

	// Start workers
	for i := 0; i < in.Thread; i++ {
		go func() {
			for {
				select {
				case <-stop:
					return
				default:
					// Busy work (tight loop keeps CPU busy)
					_ = 1 + 1
				}
			}
		}()
	}

	// Run for the given duration
	time.Sleep(time.Duration(in.Dose) * time.Second)

	// Stop workers
	for i := 0; i < in.Thread; i++ {
		stop <- true
	}
	close(stop)
	return nil
}
