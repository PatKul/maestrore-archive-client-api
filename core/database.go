package core

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDatabase struct {
	Db          *sql.DB
	Config      *Config
	Encryptor   *Encryptor
	IsConnected bool
}

func NewMySqlDatabase(config *Config, encryptor *Encryptor) *MySqlDatabase {
	return &MySqlDatabase{
		Config:    config,
		Encryptor: encryptor,
	}
}

func (d *MySqlDatabase) GetConnection() *sql.DB {
	return d.Db
}

func (d *MySqlDatabase) Connect() error {
	port := "3306"
	// password, error := d.Encryptor.Decrypt(d.Config.DatabasePassword)
	// if error != nil {
	// 	return fmt.Errorf("failed to decrypt database password: %s", ErrorInternalServerError)
	// }

	maxAttempt := 5
	retryInterval := time.Second
	maximumRetryInterval := 15 * time.Second

	attempt := 0

	for {
		databaseName := "maestrore_vision_gate"
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			d.Config.DatabaseUser,
			d.Config.DatabasePassword,
			d.Config.DatabaseHost,
			port,
			databaseName,
		)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			attempt++
			if attempt >= maxAttempt {
				return fmt.Errorf("failed to connect to database after %d attempts: %s", maxAttempt, err.Error())
			}
			time.Sleep(retryInterval)
			retryInterval *= 2
			if retryInterval > maximumRetryInterval {
				retryInterval = maximumRetryInterval
			}
			continue
		}

		d.Db = db
		d.IsConnected = true
		break
	}

	return nil
}

func (d *MySqlDatabase) Close() {
	if d.IsConnected {
		d.Db.Close()
	}
}
