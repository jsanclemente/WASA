package api

import "WASA/service/database"

// User struct represent a user in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See Fountain.FromDatabase (below) to understand why.
type User struct {
	ID         uint64
	Username   string `json:"userName"`
	Followers  []int  `json:"followers"`
	Following  []int  `json:"following"`
	Posts      []int  `json:"posts"`
	Nfollowers uint64 `json:"nFollowers"`
	Nfollowing uint64 `json:"nFollowing"`
	Nposts     uint64 `json:"nPhotos"`
}

type Photo struct {
	ID        uint64   `json:"id"`
	Image     []byte   `json:"image"`
	Ncomments uint64   `json:"nComments"`
	Date      string   `json:"date"`
	Time      string   `json:"time"`
	Comments  []uint64 `json:"comments"`
	Nlikes    uint64   `json:"nLikes"`
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.Fountain" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
func (u *User) FromDatabase(user database.User) {
	u.ID = user.ID
	u.Username = user.Username
	u.Followers = user.Followers
	u.Following = user.Following
	u.Posts = user.Posts
	u.Nfollowers = user.Nfollowers
	u.Nfollowing = user.Nfollowing
	u.Nposts = user.Nposts
}

// ToDatabase returns the user in a database-compatible representation
func (u *User) ToDatabase() database.User {
	return database.User{
		ID:         u.ID,
		Username:   u.Username,
		Followers:  u.Followers,
		Following:  u.Following,
		Posts:      u.Posts,
		Nfollowers: u.Nfollowers,
		Nfollowing: u.Nfollowing,
		Nposts:     u.Nposts,
	}
}

func (p *Photo) FromDatabase(photo database.Photo) {
	p.ID = photo.ID
	p.Image = photo.Image
	p.Ncomments = photo.Ncomments
	p.Date = photo.Date
	p.Time = photo.Time
	p.Comments = photo.Comments
	p.Nlikes = photo.Nlikes
}

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		ID:        p.ID,
		Image:     p.Image,
		Ncomments: p.Ncomments,
		Date:      p.Date,
		Time:      p.Time,
		Comments:  p.Comments,
		Nlikes:    p.Nlikes,
	}
}