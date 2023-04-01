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
	rt.router.DELETE("/users/:userId/banned/:bannedId", rt.wrap(rt.banUser))
	rt.router.POST("/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/photos/:photoId", rt.wrap(rt.deletePhoto))
	rt.router.PUT("/photos/:photoId/likes/:userId", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/photos/:photoId/likes/:userId", rt.wrap(rt.unlikePhoto))
	rt.router.PUT("/users/:userId/name", rt.wrap(rt.setMyUsername))
	rt.router.GET("/photos/:userId", rt.wrap(rt.getMyStream))
	rt.router.GET("/users/:userId}", rt.wrap(rt.getUserProfile))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
