package manager

import (
	"github.com/isayme/go-requestbin/app/conf"
	"github.com/isayme/go-requestbin/app/mongo"
)

// Manager ...
type Manager struct {
	Mongo *mongo.Mongo
}

var globalManager *Manager

// Init create global manager at startup
func Init(config *conf.Config) {
	manager := Manager{}

	mongo := mongo.NewMongo(&config.Mongo)
	manager.Mongo = mongo

	globalManager = &manager
}

// Get get global manager
func Get() *Manager {
	return globalManager
}
