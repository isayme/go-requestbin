package manager

import (
	"github.com/isayme/go-requestbin/app/conf"
	"github.com/isayme/go-requestbin/app/mongo"
)

// Manager ...
type Manager struct {
	Session *mongo.Session
}

var globalManager *Manager

// Init create global manager at startup
func Init(config *conf.Config) {
	manager := Manager{}

	session, err := mongo.NewSession(&config.Mongo)
	if err != nil {
		panic(err)
	}
	manager.Session = session

	globalManager = &manager
}

// Get get global manager
func Get() *Manager {
	return globalManager
}
