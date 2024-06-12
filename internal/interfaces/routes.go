package interfaces

import (
	"github.com/gin-gonic/gin"
	"gron/internal/service"
)

type Handler struct {
	gronService *service.GronService
}

func NewHandler(s *service.GronService) *Handler {
	return &Handler{
		gronService: s,
	}
}
func NewRouter(h *Handler) *gin.Engine {
	r := gin.Default()
	//r := engine.NewEngine(engine.WithLogger(false))
	// 使用gin中间件
	//r.Use(InfoLog())
	project := r.Group("xtimer")
	project.POST("/createTimer", h.CreateTimer)
	project.GET("/enableTimer", h.EnableTimer)
	project.POST("/callback", h.TestCallback)
	return r
}
