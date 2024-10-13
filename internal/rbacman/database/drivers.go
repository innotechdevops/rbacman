package database

import (
	"github.com/innotechdevops/maria-driver/pkg/mariadriver"
	"github.com/innotechdevops/mgo-driver/pkg/mgodriver"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type Drivers interface {
	GetMariaDB() *sqlx.DB
	GetMongoDB() *mongo.Database
}

type drivers struct {
	MongoDB     *mongo.Database
	MariaDB     *sqlx.DB
	MariaDriver mariadriver.MariaDBDriver
	MongoDriver mgodriver.MongoDriver
}

func (d *drivers) GetMariaDB() *sqlx.DB {
	return d.MariaDB
}

func (d *drivers) GetMongoDB() *mongo.Database {
	return d.MongoDB
}

func NewDrivers(
	mongoDB mgodriver.MongoDriver,
	mariaDb mariadriver.MariaDBDriver,
) Drivers {
	return &drivers{
		MongoDB:     mongoDB.Connect(),
		MongoDriver: mongoDB,
		MariaDB:     mariaDb.Connect(),
		MariaDriver: mariaDb,
	}
}
