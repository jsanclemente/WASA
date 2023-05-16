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

var ErrUserSubjectNotExists = errors.New("The user that starts the action doesn't exist")
var ErrUserPredicateNotExists = errors.New("The user that recieves the action doesn`t exists")
var ErrPhotoNotExits = errors.New("The photo doesn't exists")
var ErrCommentNotExists = errors.New("The comment doesn't exist")
var ErrUser1alreadyFollows2 = errors.New("The user A is already following user B")
var ErrNotFollowing = errors.New("The user A doesn't follow user B")
var ErrUserAIsBanned = errors.New("The user A is banned by B")
var ErrUserBIsBanned = errors.New("The user B is banned by A")
var ErrAlreadyBanned = errors.New("The user is already banned")
var ErrNotBanned = errors.New("The user you are trying to unban it's not banned")
var ErrNotHisPhoto = errors.New("The user is trying to delete a post he hasn't posted")
var ErrNotHisLike = errors.New("The user is trying to remove the like of a post he hasn't liked")
var ErrUsernameAlreadyRegistered = errors.New("This username is already used")
var ErrPhotoAlreadyLiked = errors.New("You can't like a photo twice")

type User struct {
	ID         uint64
	Username   string
	Followers  []int
	Following  []int
	Posts      []int
	Nfollowers uint64
	Nfollowing uint64
	Nposts     uint64
}

type Photo struct {
	ID        uint64 //	Id of the photo
	Username  string //	Username of the user that posts the photo
	Image     []byte
	Ncomments uint64
	Date      string
	Time      string
	Comments  []uint64 //	List of id's of all the comments for that photo
	Likes     []uint64 //	List of id's of users that liked the photo
	Nlikes    uint64
}

type Comment struct {
	IdComment uint64
	Username  string // User that comments the photo
	Comment   string
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// If the user already exists, the user is logged and his id is returned.
	// Otherwise, the user is created, and his id is returned
	Login(userName string) (uint64, error)

	// "follower" will follow the user "followed", if the operation succeed it will return the number of followed users
	FollowUser(follower uint64, followed uint64) (uint64, error)

	// "unfollower" will unfollow "unfollowed" user. Returns the number of followed users by the user after the operation.
	UnfollowUser(unfollower uint64, unfollowed uint64) (uint64, error)

	// "banner" will ban the "banned" user, it will return the number of the banned users after the operations
	BanUser(banner uint64, banned uint64) (uint64, error)

	// "unbanner" will unban "unbanned" user. It returns the number of banned users after the operation.
	UnbanUser(unbanner uint64, unbanned uint64) (uint64, error)

	// userName posts the image "url". Returns the id of the photo posted
	UploadPhoto(image []byte, idUser uint64) (uint64, error)

	// Deletes the photo idPhoto. Returns the id of the deleted photo
	DeletePhoto(idUser uint64, idPhoto uint64) (uint64, error)

	// The user "userId" likes the photo identified by "photoId". Returns the number of likes after the operation
	LikePhoto(userId uint64, photoId uint64) (uint64, error)

	// The user identified by "userId" unlikes the photo identified by "photoId". Returns the number of likes after the operation
	UnlikePhoto(userId uint64, photoId uint64) (uint64, error)

	// The user identified by "userId" comments the description "comment" to the photo "photoId". Returns the number of comments after the operation
	CommentPhoto(userId uint64, photoId uint64, comment string) (uint64, error)

	// The user removes the "commentId" on photo "photoId"
	UncommentPhoto(photoId uint64, commentId uint64) (uint64, error)

	// The username associated to user "userId" changes to "username". Returns the old username.
	SetMyUserName(userId uint64, username string) (string, error)

	// Given a username for one user, returns the stream for that user.
	GetMyStream(userId uint64) ([]Photo, error)

	// Given an id for one user, returns the profile for that user.
	GetUserProfile(userId uint64) (User, error)

	// Given an id for the user, checks if the user already exists
	UserExists(userId uint64) bool

	// If "follower" is following "followed" returns true, otherwise returns false
	IsFollowing(follower uint64, followed uint64) bool

	// If "photoId" returns true, otherwise returns false
	PhotoExists(photoId uint64) bool

	UserLiked(userId uint64, photoId uint64) bool

	GetComments(photoId uint64) ([]Comment, error)

	SearchUser(username string) ([]User, error)

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

	_, err := db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, fmt.Errorf("enabling foreign key support: %w", err)
	}

	// Check if table USERS exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Users (id INTEGER PRIMARY KEY,
			 username TEXT, nfollowers INTEGER NOT NULL, nfollowing INTEGER NOT NULL,
			 nposts INTEGER NOT NULL, nbans INTEGER NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure <Users>: %w", err)
		}
	}

	// Check if table PHOTO exists. If not, the database is empty, and we need to create the structure
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Photos';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Photos (id INTEGER PRIMARY KEY,
			 nLikes INTEGER NOT NULL, nComments INTEGER NOT NULL, imageData BLOB NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure <Photos>: %w", err)
		}
	}

	// Check if table COMMENTS exists. If not, the database is empty, and we need to create the structure
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Comments';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Comments (
						user_id INTEGER NOT NULL,
						comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
			 			photo_id INTEGER  NOT NULL,
						comment TEXT NOT NULL,
						FOREIGN KEY (user_id) REFERENCES Users(id),
						FOREIGN KEY (photo_id) REFERENCES Photos(id) ON DELETE CASCADE
					);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure <Comments>: %w", err)
		}
	}

	// Check if table POSTS exists. If not, the database is empty, and we need to create the structure
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Posts';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Posts (photo_id INTEGER NOT NULL, user_id INTEGER NOT NULL,
			date DATETIME, PRIMARY KEY(user_id,photo_id), FOREIGN KEY (photo_id) REFERENCES Photos(id) ON DELETE CASCADE, 
			FOREIGN KEY (user_id) REFERENCES Users(id));`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure <Posts>: %w", err)
		}
	}

	// Check if table LIKES exists. If not, the database is empty, and we need to create the structure
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Likes';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Likes (user_id INTEGER NOT NULL, photo_id INTEGER NOT NULL,
			PRIMARY KEY(user_id,photo_id), FOREIGN KEY (user_id) REFERENCES Users(id),
			FOREIGN KEY (photo_id) REFERENCES Photos(id) ON DELETE CASCADE);`
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
