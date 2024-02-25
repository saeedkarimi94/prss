package router

import (
	"github.com/einkaaf/prss/handler"
	"github.com/gin-gonic/gin"
)

func Zoomit_Router(r *gin.RouterGroup) {
	r.GET("/zoomit", handler.ZoomitHandler)
}

func Digiato_Router(r *gin.RouterGroup) {
	r.GET("/digiato", handler.DigiatoHandler)
}

func Tasnim_Router(r *gin.RouterGroup) {
	r.GET("/tasnim", handler.TasnimHandler)
}

func Tabnak_Router(r *gin.RouterGroup) {
	r.GET("/tabnak", handler.TabnakHandler)
}

func YJC_Router(r *gin.RouterGroup) {
	r.GET("/yjc", handler.YJCHandler)
}

func Zoomg_Router(r *gin.RouterGroup) {
	r.GET("/zoomg", handler.ZoomgHandler)
}

func Khodrobank_Router(r *gin.RouterGroup) {
	r.GET("/khodrobank", handler.KhodroBankHandler)
}
