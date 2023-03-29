package main

import (
	"github.com/enuesaa/teatime-app/backend/controller"
	"github.com/gin-gonic/gin"
)

func jsonMiddleware() gin.HandlerFunc {
	// https://stackoverflow.com/questions/41109065/golang-gin-gonic-content-type-not-setting-to-application-json-with-c-json
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(jsonMiddleware())

	base := router.Group("/api")
	{
		settingRoute := base.Group("/v1.Setting")
		settingCtl := controller.SettingController{}
		settingRoute.POST("/GetAppearance", settingCtl.Get)
		settingRoute.POST("/PutAppearance", settingCtl.Put)

		bookmarkRoute := base.Group("/v1.Bookmark")
		bookmarkCtl := controller.BookmarkController{}
		bookmarkRoute.POST("/ListBookmarks", bookmarkCtl.List)
		bookmarkRoute.POST("/GetBookmark", bookmarkCtl.Get)
		bookmarkRoute.POST("/AddBookmark", bookmarkCtl.Add)
		bookmarkRoute.POST("/UpdateBookmark", bookmarkCtl.Update)
		bookmarkRoute.POST("/DeleteBookmark", bookmarkCtl.Delete)

		feedRoute := base.Group("/v1.Feed")
		feedCtl := controller.FeedController{}
		feedRoute.POST("/AddFeed", feedCtl.Add)
		feedRoute.POST("/GetFeed", feedCtl.Get)
		feedRoute.POST("/ListItems", feedCtl.ListItems)
		feedRoute.POST("/GetAppearance", feedCtl.GetAppearance)
		feedRoute.POST("/UpdateAppearance", feedCtl.UpdateAppearance)
		feedRoute.POST("/Fetch", feedCtl.Fetch)
		feedRoute.POST("/DeleteFeed", feedCtl.Delete)

		boardRoute := base.Group("/v1.Board")
		boardCtl := controller.BoardController{}
		boardRoute.POST("/AddBoard", boardCtl.Add)
		boardRoute.POST("/Checkin", boardCtl.Checkin)
		boardRoute.POST("/ListTimeline", boardCtl.ListTimeline)
		boardRoute.POST("/ArchiveBoard", boardCtl.Archive)
		boardRoute.POST("/UnArchiveBoard", boardCtl.UnArchive)
	}

	return router
}
