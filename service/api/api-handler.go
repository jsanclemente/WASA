package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/users", rt.wrap(rt.login))
	rt.router.PUT("/users/:userId/following/:followedId", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:userId/following/:followedId", rt.wrap(rt.unfollowUser))
	rt.router.PUT("/users/:userId/banned/:bannedId", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:userId/banned/:bannedId", rt.wrap(rt.unbanUser))
	rt.router.POST("/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/photos/:photoId", rt.wrap(rt.deletePhoto))
	rt.router.PUT("/photos/:photoId/likes/:userId", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/photos/:photoId/likes/:userId", rt.wrap(rt.unlikePhoto))
	rt.router.POST("/photos/:photoId/comments", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/photos/:photoId/comments/:commentId", rt.wrap(rt.uncommentPhoto))
	rt.router.PUT("/users/:userId/name", rt.wrap(rt.setMyUsername))
	rt.router.GET("/feed/:userId", rt.wrap(rt.getMyStream))
	rt.router.GET("/users/:userId/profile", rt.wrap(rt.getUserProfile))
	rt.router.GET("/photos/:photoId/comments", rt.wrap(rt.getComments))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
