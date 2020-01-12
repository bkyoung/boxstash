package db

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"sync"
	"time"

	"boxstash/internal/boxstash/repository/shared/migrate/sqlite"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Connect to a database and verify with a ping.
func Connect(driver, datasource string) (*DB, error) {
	db, err := sql.Open(driver, datasource)
	if err != nil {
		logrus.StandardLogger().WithFields(logrus.Fields{
			"driver": driver,
			"datasource": datasource,
			"error": err,
		}).Error("ERROR connecting to database")
		return nil, err
	}
	logrus.StandardLogger().WithFields(logrus.Fields{
		"driver": driver,
		"datasource": datasource,
		"error": err,
	}).Debug("SUCCESS opening database")
	if err := pingDatabase(db); err != nil {
		logrus.StandardLogger().WithFields(logrus.Fields{
			"driver": driver,
			"datasource": datasource,
			"error": err,
		}).Error("ERROR pinging database")
		return nil, err
	}
	logrus.StandardLogger().WithFields(logrus.Fields{
		"driver": driver,
		"datasource": datasource,
		"error": err,
	}).Debug("SUCCESS pinging database")
	if err := setupDatabase(db, driver); err != nil {
		logrus.StandardLogger().WithFields(logrus.Fields{
			"driver": driver,
			"datasource": datasource,
			"error": err,
		}).Error("ERROR setting up to database schema")
		return nil, err
	}
	logrus.StandardLogger().WithFields(logrus.Fields{
		"driver": driver,
		"datasource": datasource,
		"error": err,
	}).Debug("SUCCESS migrating database schema")

	var engine Driver
	var locker Locker
	// TODO: Add support for MySQL and PostgreSQL
	switch driver {
	default:
		engine = Sqlite
		locker = &sync.RWMutex{}
	}

	return &DB{
		conn:   sqlx.NewDb(db, driver),
		lock:   locker,
		driver: engine,
	}, nil
}

// Ping the database to ensure a connection can be established
// before we proceed with the database setup and migration.  Try
// for up to 30 seconds.
func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		time.Sleep(time.Second)
	}
	return
}

// Setup the database by performing automated database migration
func setupDatabase(db *sql.DB, driver string) error {
	switch driver {
	default:
		return sqlite.Migrate(db)
	}
}

// Driver defines the database driver.
type Driver int

// Database driver enums.
const (
	Sqlite = iota + 1
	Mysql
	Postgres
)

// Scanner interface defines methods on an object that can be scanned for values
type Scanner interface {
	Scan(dest ...interface{}) error
}

// Locker interface defines methods on an db object that can lock/unlock
type Locker interface {
	// Read-write
	Lock()
	Unlock()

	// Read-only
	RLock()
	RUnlock()
}

// Binder interface defines methods for database field bindings
type Binder interface {
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
}

// Queryer interface defines methods for querying the database
type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Execer interface defines methods for executing read/write commands in the db
type Execer interface {
	Queryer
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// DB is a connection pool for the database
type DB struct {
	conn   *sqlx.DB
	lock   Locker
	driver Driver
}

// View executes a function as a read-only transaction. Any error that is
// returned from the inner function is returned by View
func (db *DB) View(fn func(Queryer, Binder) error) error {
	db.lock.RLock()
	err := fn(db.conn, db.conn)
	db.lock.RUnlock()
	return err
}

// Lock obtains a write lock on the database (sqlite only) and executes
// a function. Any error that is returned from that function is returned
// by Lock
func (db *DB) Lock(fn func(Execer, Binder) error) error {
	db.lock.Lock()
	err := fn(db.conn, db.conn)
	db.lock.Unlock()
	return err
}

// Update executes a function as a read-write transaction. If no error is
// returned from the function then the transaction is committed. If an
// error is returned then the transaction is rolled back. Any error
// that is returned from the function or the commit is returned by Update
func (db *DB) Update(fn func(Execer, Binder) error) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			err = tx.Rollback()
			debug.PrintStack()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx, db.conn)
	return err
}

// Driver returns the SQL driver
func (db *DB) Driver() Driver {
	return db.driver
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.conn.Close()
}
