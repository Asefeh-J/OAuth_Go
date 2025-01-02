package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Define API routes
	r.GET("/api/v1/admin", AuthHandler)
	r.GET("/api/v1/admin/callback", AuthCallbackHandler)
	r.GET("/api/v1/emails", FetchEmailsHandler)
	r.GET("/api/v1/icon", IconHandler)
	r.GET("/api/v1/drive/files", FetchDriveFilesHandler)
}

func StartServer() {
	r := gin.Default()
	RegisterRoutes(r)
	r.Run(":8888")
}
