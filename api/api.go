package api

import (
	"github.com/einkaaf/prss/router"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	engine := gin.Default()

	api := engine.Group("/api")
	{
		router.Zoomit_Router(api)
		router.Digiato_Router(api)
		router.Tasnim_Router(api)
		router.Tabnak_Router(api)
		router.YJC_Router(api)
		router.Zoomg_Router(api)
	}

	return engine
}
