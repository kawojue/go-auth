package routes

import "github.com/gin-gonic/gin"

var file *gin.RouterGroup

func FileRoutes(router *gin.RouterGroup) {
	file = router.Group("/file")
	{

	}
}
