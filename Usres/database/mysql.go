package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type DatabaseConfig struct {
	DbUsername        string        `default:"root" split_words:"true"`
	DbPassword        string        `default:"" split_words:"true"`
	DbHost            string        `default:"localhost" split_words:"true"`
	DbPort            int           `default:"3306" split_words:"true"`
	DbName            string        `default:"main" split_words:"true"`
}

func NewMySQL() (*sql.DB, error) {
	db, err := sql.Open(
		"mysql", 
		"root:root@tcp(localhost:3306)/BalancherProject",)
	if err != nil {
				logrus.Error(err)
				return nil, err
	} else {
		logrus.Info(fmt.Sprintf(`{"label":"info", "message":"Database Connected ", "time":"%s"}`, time.Now()))
	}

	// make sure connection is available
	err = db.Ping()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	sql := "CREATE TABLE IF NOT EXISTS users(id int NOT NULL AUTO_INCREMENT, name varchar(255), last_name varchar(255), created_at TIMESTAMP DEFAULT(CURRENT_TIMESTAMP), updated_at TIMESTAMP DEFAULT(CURRENT_TIMESTAMP), PRIMARY  KEY (id));"
	
	stmt , err := db.Prepare(sql)

	if err != nil {
		logrus.Error(err)
	}

	result, err := stmt.Exec()
	if err != nil{
		logrus.Error(err)
		return nil, err
	}

	logrus.Info(fmt.Sprintf(`{"label":"info", "message":"%s", "time":"%s"}`, result, time.Now()))

	return db, nil
}

