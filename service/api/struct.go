package api

import "WASA/service/database"

// User struct represent a user in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See Fountain.FromDatabase (below) to understand why.
type User struct {
	ID         uint64
	username   string `json:"userName"`
	followers  []int  `json:"followers"`
	following  []int  `json:"following"`
	posts      []int  `json:"posts"`
	nFollowers uint64 `json:"nFollowers"`
	nFollowing uint64 `json:"nFollowing"`
	nPosts     uint64 `json:"nPhotos"`
}

type Photo struct {
	ID        uint64   `json:"id"`
	nLikes    uint64   `json:"nLikes"`
	nComments uint64   `json:"nComments"`
	date      string   `json:"date"`
	comments  []uint64 `json:"comments"`
	url       string   `json:"url"`
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
	u.username = user.Username
	u.followers = user.Followers
	u.following = user.Following
	u.posts = user.Posts
	u.nFollowers = user.Nfollowers
	u.nFollowing = user.Nfollowing
	u.nPosts = user.Nposts
}

// ToDatabase returns the user in a database-compatible representation
func (u *User) ToDatabase() database.User {
	return database.User{
		ID:         u.ID,
		Username:   u.username,
		Followers:  u.followers,
		Following:  u.following,
		Posts:      u.posts,
		Nfollowers: u.nFollowers,
		Nfollowing: u.nFollowing,
		Nposts:     u.nPosts,
	}
}
