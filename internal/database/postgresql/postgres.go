package database

import (
	"errors"
	"fmt"
	"sync"

	"waifunet/internal/config"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       int
	Email    string
	Username string
	Password string
}

type conPool struct {
	pool  *sqlx.DB
	mutex *sync.Mutex
}

// FIX ME: возможно удалить просто или инициализировать чисто в функции (в будующем посмотреть надо)
var (
	connStr  string
	driverDb string
)

// FIX ME: вынести в config
var (
	maxConnections int = 10
	numConnections int // текущее количество коннектов
)

func PostgreRoutine(cfg *config.Config) error {
	prepareConnectionDataBase(cfg)
	err := initPool()
	return err
}

func (p *conPool) getConnection() (*sqlx.DB, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if numConnections == maxConnections {
		err := errors.New("max connection limit")
		return nil, err
	}
	numConnections++
	return p.pool, nil
}

func (p *conPool) releaseConnection(conn *sqlx.DB) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	numConnections--
}

func prepareConnectionDataBase(cfg *config.Config) {
	driverDb = "postgres"
	connStr = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.DbName)
	fmt.Println(connStr)
}

func initPool() error {
	// create connPool of db
	pool, err := sqlx.Open(driverDb, connStr)
	if err != nil {
		return err
	}
	pool.SetMaxOpenConns(maxConnections)
	return nil
}
