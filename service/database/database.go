/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	ID         uint64
	username   string
	nFollowers uint64
	nFollowing uint64
	nPosts     uint64
}

type Photo struct {
	ID        uint64
	nLikes    uint64
	nComments uint64
	url       string
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// Ping checks whether the database is available or not (in that case, an error will be returned)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table USERS exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Users (id INTEGER NOT NULL PRIMARY KEY,
			 username TEXT, nfollowers INTEGER NOT NULL, nfollowing INTEGER NOT NULL,
			 nposts INTEGER NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure <Users>: %w", err)
		}
	}

	// Check if table PHOTO exists. If not, the database is empty, and we need to create the structure
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Photos';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Photos (id INTEGER NOT NULL PRIMARY KEY,
			 nLikes INTEGER NOT NULL, nComments INTEGER NOT NULL, image BLOB NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure <Photos>: %w", err)
		}
	}

	// Check if table COMMENTS exists. If not, the database is empty, and we need to create the structure
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Comments';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Comments (user_id INTEGER NOT NULL, comment_id INTEGER NOT NULL,
			 photo_id INTEGER NOT NULL, comment TEXT NOT NULL,
			PRIMARY KEY (user_id,comment_id,photo_id), FOREIGN KEY (user_id) REFERENCES Users(id),
			FOREIGN KEY (photo_id) REFERENCES Photos(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure <Comments>: %w", err)
		}
	}

	// Check if table POSTS exists. If not, the database is empty, and we need to create the structure
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Posts';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Posts (user_id INTEGER NOT NULL, photo_id INTEGER NOT NULL,
			date DATE, PRIMARY KEY(user_id,photo_id), FOREIGN KEY (user_id) REFERENCES Users(id),
			FOREIGN KEY (photo_id) REFERENCES Photos(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure <Posts>: %w", err)
		}
	}

	// Check if table POSTS exists. If not, the database is empty, and we need to create the structure
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Likes';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Likes (user_id INTEGER NOT NULL, photo_id INTEGER NOT NULL,
			PRIMARY KEY(user_id,photo_id), FOREIGN KEY (user_id) REFERENCES Users(id),
			FOREIGN KEY (photo_id) REFERENCES Photos(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure <Likes>: %w", err)
		}
	}

	// Check if table FOLLOWS exists. If not, the database is empty, and we need to create the structure
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Follows';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Follows (follower_id INTEGER NOT NULL, followed_id INTEGER NOT NULL,
			PRIMARY KEY(follower_id,followed_id), FOREIGN KEY (follower_id) REFERENCES Users(id),
			FOREIGN KEY (followed_id) REFERENCES Users(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure <Follows>: %w", err)
		}
	}

	// Check if table BANS exists. If not, the database is empty, and we need to create the structure
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Bans';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Bans (banner_id INTEGER NOT NULL, banned_id INTEGER NOT NULL,
			PRIMARY KEY(banner_id,banned_id), FOREIGN KEY (banner_id) REFERENCES Users(id),
			FOREIGN KEY (banned_id) REFERENCES Users(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure <Bans>: %w", err)
		}
	}
	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
