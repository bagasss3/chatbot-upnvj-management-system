package database

import (
	"cbupnvj/config"

	"github.com/kamva/mgm/v3"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDatabase() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(config.DBDSN()))
	if err != nil {
		log.WithField("dbDSN", config.DBDSN()).Fatal("Failed to connect:", err)
	}
	log.Info("Success connect database")
}
