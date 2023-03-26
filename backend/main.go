// teatime-app backend
package main

import (
	"io"
	"os"	

	"github.com/enuesaa/teatime-app/backend/controller/setting"
	"github.com/enuesaa/teatime-app/backend/controller/bookmark"
	"github.com/gin-contrib/pprof"
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
		settingRoute.POST("/GetAppearance", setting.GetAppearance)

		bookmarkRoute := base.Group("/v1.Bookmark")
		bookmarkRoute.POST("/AddBookmark", bookmark.AddBookmark)
	}
	return router
}

func main() {
	f, _ := os.Create("tmp/gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	router := setupRouter()
	pprof.Register(router)
	router.Run(":80")
}
