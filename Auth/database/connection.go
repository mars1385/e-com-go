package database

import (
	"AUTH/ent"
	"AUTH/ent/migrate"
	logger "AUTH/helper"
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	Database *ent.Client
)

func DbConnection() error {

	dbHost := viper.GetString("DB_HOST")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	dbSSl := viper.GetString("DB_SSL")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbSSl)

	client, err := ent.Open(dbHost, connectionString)

	if err != nil {

		logger.Errorf(err.Error())
		logger.Fatalf("Db connection error")
		return err
	}

	Database = client

	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	logger.Infof("Db connection success")

	return nil
}
