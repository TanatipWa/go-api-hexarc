package entity

import "github.com/TanatipWa/product-api/db"

type EnvConfig struct {
	Host  string
	Port  int
	Mongo db.MongoConfig
}
