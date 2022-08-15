package config

import (
	"os"

	. "catalyst-token/models/admin-models"
	. "catalyst-token/models/invite-token-models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupConnection() *gorm.DB {
	// databaseURI := make(chan string, 1)

	// if os.Getenv("GO_ENV") != "production" {
	// 	databaseURI <- utils.GodotEnv("DEV_URI")
	// } else {
	// 	databaseURI <- os.Getenv("PROD_URI")
	// }

	db, err := gorm.Open(postgres.Open("postgres://postgres@localhost:5432/development"), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		Admin{},
		InviteToken{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
