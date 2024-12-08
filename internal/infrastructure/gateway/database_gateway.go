package gateway

import (
	"fmt"
	"original-card-game-backend/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConnectionError struct {
	Err error
}

func (e *DatabaseConnectionError) Error() string {
	return fmt.Sprintf("database connection error: %o", e.Err)
}

type DatabaseConfig struct {
	userName     string
	userPassword string
	databaseName string
	port         string
}

type DatabaseGateway struct {
	config DatabaseConfig
}

func (g *DatabaseGateway) createDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/main?charset=utf8mb4&parseTime=True&loc=Local",
		g.config.userName,
		g.config.userPassword,
		g.config.databaseName,
		g.config.port,
	)
}

func (g *DatabaseGateway) Connect() (*gorm.DB, error) {
	dsn := g.createDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, &DatabaseConnectionError{
			Err: err,
		}
	}

	return db, nil
}

func NewDatabaseGateway(
	config *configs.Config,
) (*DatabaseGateway, error) {
	return &DatabaseGateway{
		config: DatabaseConfig{
			userName:     config.MySQL.UserName,
			userPassword: config.MySQL.UserPassword,
			databaseName: config.MySQL.DatabaseName,
			port:         config.MySQL.Port,
		},
	}, nil
}
