package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/controllers"
	"github.com/kawojue/go-auth/middlewares"
)

var file *gin.RouterGroup

func FileRoutes(router *gin.RouterGroup) {
	file = router.Group("/file")
	{
		file.POST(
			"/avatar",
			middlewares.VerifyAuth(),
			middlewares.LimitRequestBodySize(5<<20),
			controllers.UploadAvatar,
		)

	}
}
